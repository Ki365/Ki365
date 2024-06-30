package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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

func EndpointUploadProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading file...")

	// Fixed 167.7MB limit on file upload
	// TODO: have a better way to handle this
	r.ParseMultipartForm(10 << 2)

	file, header, err := r.FormFile("project-archive")

	if err != nil {
		fmt.Println("Error retrieving the uploaded file, canceling upload...")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded filename: %+v\n", header.Filename)
	fmt.Printf("Filesize: %+v\n", header.Size)
	fmt.Printf("MIME: header: %+v\n", header.Header)

	// TODO: Add environment variable to point to mount
	// Create folder structure if does not exist
	folderPath := filepath.Join("./repos/upload/archives")
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: remove original filename and obscure to increase security
	// Create empty file
	fp := filepath.Join(folderPath, header.Filename)

	dst, err := os.Create(fp)
	if err != nil {
		fmt.Println(err)
	}
	defer dst.Close()

	// Write contents of upload to file
	fileBytes, err := io.Copy(dst, file)
	if err != nil {
		fmt.Println(err)
	}

	// TODO: Add more information to JSON response for client
	// Return JSON response with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fileBytes)

	// Handle new imported file
	err = handleNewProjectArchive(fp, "./repos/repos")
	if err != nil {
		fmt.Println(err)
	}

	// pass newly created project to handle project function
	err = handleNewProject(fp, "./repos/repos", "./repos/store/repos.json")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	port := flag.String("p", "8100", "port to serve directory on")
	directory := flag.String("d", ".", "directory of files to serve")
	flag.Parse()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	apiPrefix := "/api"
	router.HandleFunc(apiPrefix+"/upload/project", EndpointUploadProject) // upload project endpoint
	router.Get(apiPrefix+"/ping", EndpointGetPing)                        // responds with pong
	router.Get(apiPrefix+"/projects", EndpointGetProjects)                // get list of all projects in federation
	router.Get("/", SPAHandler(*directory))
	router.NotFound(SPAHandler(*directory))

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + *port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(srv.ListenAndServe())
}
