package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ki365/ki365/server/abatement"
	e "github.com/ki365/ki365/server/endpoints"
	s "github.com/ki365/ki365/server/structure"
)

// Credit: https://www.asciiart.eu/text-to-ascii-art | figlet release 2.1
// Font: Slant by Glenn Chappell 3/93
// Configuration: Slant, PlusBox V2, Horizontal padding: 5, Vertical padding: 0
const header = `
+========================================+
|         __ __ _ _____ _____ ______     |
|        / //_/(_)__  // ___// ____/     |
|       / ,<  / / /_ </ __ \/___ \       |
|      / /| |/ /___/ / /_/ /___/ /       |
|     /_/ |_/_//____/\____/_____/        |
|                                        |
+========================================+
`

// SPAHandler serves a single page application.
func SPAHandler(staticPath string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Join internally call path.Clean to prevent directory traversal
		path := filepath.Join(staticPath, r.URL.Path)

		// check whether a file exists or is a directory at the given path
		fi, err := os.Stat(path)
		if os.IsNotExist(err) || fi.IsDir() {

			// set cache control header to prevent caching
			// this is to prevent the browser from caching the index.html
			// and serving old build of SPA App
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

			// file does not exist or path is a directory, serve index.html
			http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
			return
		}

		if err != nil {
			// if we got an error (that wasn't that the file doesn't exist) stating the
			// file, return a 500 internal server error and stop
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set cache control header to serve file for a year
		// static files in this case need to be cache busted
		// (usualy by appending a hash to the filename)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

		// otherwise, use http.FileServer to serve the static file
		http.FileServer(http.Dir(staticPath)).ServeHTTP(w, r)
	})
}

func ProjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projectID := chi.URLParam(r, "projectID")
		project, err := dbGetProject(projectID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "project", project)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func dbGetProject(id string) (*s.Project, error) {
	projects, err := s.ParseProjectsJSON(s.Dirs().Store.RepoConfig)
	if err != nil {
		return nil, err
	}

	for _, a := range projects.Projects {
		if a.Id == id {
			return &a, nil
		}
	}
	return nil, errors.New("project not found")
}

// confirm prompts the user with `s` and returns a bool indicating either a yes or no
// The user's input will be lowercased and trimmed.
// If the input begins with anything other than 'y', it returns false.
// It accepts an int `tries` which represents the number of attempts before it returns false
func confirm(s string, tries int) bool {
	r := bufio.NewReader(os.Stdin)

	for ; tries > 0; tries-- {
		fmt.Printf("%s [y (confirm) / n (cancel)]: ", s)

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (such as newline character)
		if len(res) < 2 {
			continue
		}

		v := strings.ToLower(strings.TrimSpace(res))[0]

		if v == 'y' {
			return true
		} else if v == 'n' {
			return false
		}
	}

	return false
}

func main() {

	// Flags and parameters
	address := flag.String("a", "localhost", "address of server")
	port := flag.String("p", "8100", "port to serve directory on")
	directory := flag.String("d", ".", "directory of files to serve")
	bypassConfirm := flag.Bool("y", false, "do not prompt for data folder creation")
	skipExamples := flag.Bool("e", false, "do not add example projects")

	// Executable flags
	setDocker := flag.Bool("docker", false, "set if running executable in docker container")
	setLocalBin := flag.Bool("local-bin", false, "set if using executables from ./bin directory for testing")

	flag.Parse()

	log.Println("Initializing Ki365...")

	// Always use local bin directory when executing in Docker environment
	if *setDocker {
		// Make sure to set use local bin as programs are not in PATH in docker image
		*setLocalBin = true
		// Temporarily set flag to use correct code path
		s.UseDockerKiCadCLI = true
	}

	// Set correct code paths based on bin prefix
	if *setLocalBin {
		s.GLTFPackExecutablePath = filepath.Join(s.ConditionalBinPrefix, s.GLTFPackExecutablePath)
		s.GitHTTPBackendExecutablePath = filepath.Join(s.ConditionalBinPrefix, s.GitHTTPBackendExecutablePath)
		// NOTE: tracespace is not statically compilable, meaning it cannot be used yet in the docker image
		s.TracespaceExecutablePath = filepath.Join(s.ConditionalBinPrefix, s.TracespaceExecutablePath)
	}

	s.CreateDB()

	// TODO: Check all dependencies exist
	// TODO: Check all dependencies exist if haven't periodically

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	log.Println("Checking data folder existence...")

	// TODO: check all folder existence
	_, err := os.Stat(filepath.Join(s.Dirs().DataDir, "store"))
	if err != nil {
		log.Println("Data directory does not exist, prompting...")
		var c bool
		if !*bypassConfirm {
			fmt.Println()
			c = confirm("Do you want to create a \"data\" folder in the current directory?", 3)
			fmt.Println()
		} else {
			log.Println("Bypass confirmation flag passed.")
		}
		if c || *bypassConfirm {
			log.Println("Creating data directory...")
			err := abatement.GenerateDataFolder(s.Dirs().DataDir)
			if err != nil {
				log.Fatal("Failed to create data directory.")
			}
			if !*skipExamples {
				log.Println("Adding example repositories...")
				abatement.GenerateExamples(filepath.Join("examples", "build"), s.Dirs().RepoDir, false)
				abatement.CopyManifest("./examples/manifest-examples.json", s.Dirs().Store.ExManifest)
			}
			log.Println("Creating data directory was successful!")
		} else {
			log.Println("Permission not given.")
			log.Fatal("Ki365 must have a data directory. Exiting...")
		}
	} else {
		log.Println("Data directory found!")
	}

	log.Println("Starting Ki365 API build...")

	// TODO: Wrap this in router.Route for api prefix
	apiPrefix := "/api"
	router.Get(apiPrefix+"/ping", e.EndpointGetPing)                        // responds with pong
	router.HandleFunc(apiPrefix+"/upload/project", e.EndpointUploadProject) // upload project endpoint
	router.HandleFunc(apiPrefix+"/git/project/*", e.EndpointGetRepository)
	router.Post(apiPrefix+"/toggle-example-projects", e.EndpointAddExampleProjects)      // lists all example projects
	router.Delete(apiPrefix+"/toggle-example-projects", e.EndpointRemoveExampleProjects) // unlists all example projects
	router.Get(apiPrefix+"/toggle-example-projects", e.EndpointCheckExampleProjects)     // check if example projects are listed
	router.Route(apiPrefix+"/projects", func(r chi.Router) {
		r.Get("/", e.EndpointGetProjects)             // get list of all projects in organization
		r.Post("/link", e.EndpointPostLinkConnection) // add project link to organization
		r.Route("/{projectID}", func(r chi.Router) {
			r.Use(ProjectCtx)
			r.Get("/", e.EndpointGetProject) // get files of specific project
			r.Get("/schematics", e.EndpointGetProjectSchematics)
			r.Get("/layouts", e.EndpointGetProjectLayouts)
			r.Get("/models", e.EndpointGetProjectModels)
			r.Delete("/", e.EndpointRemoveProject) // delete specific project
		})

	})
	router.Get("/", SPAHandler(*directory))
	router.NotFound(SPAHandler(*directory))

	srv := &http.Server{
		Handler:      router,
		Addr:         *address + ":" + *port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Ki365 Initialized.")

	fmt.Print(header)

	fmt.Println()
	fmt.Println("Ki365 API Server Backend - Engineering Collaboration")
	fmt.Println("\tCopyright © Ki365 Contributors")
	fmt.Println()

	log.Printf("Serving directory %s on HTTP address: %s:%s\n", *directory, *address, *port)

	log.Fatal(srv.ListenAndServe())
}
