package endpoints

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cgi"
	"net/http/httputil"
	"os"
	"path"
	"path/filepath"

	"github.com/go-chi/render"

	p "github.com/ki365/ki365/server/internal/projects"
	s "github.com/ki365/ki365/server/structure"
)

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
	folderPath := filepath.Join(filepath.Join(s.DataDir, "upload/archives"))
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
	err = p.HandleNewProjectArchive(fp, s.RepoDir)
	if err != nil {
		fmt.Println(err)
	}

	// pass newly created project to handle project function
	err = p.HandleNewProject(fp, s.RepoDir, "", "", "", "", "", "")
	if err != nil {
		fmt.Println(err)
	}
}

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// var va = &ErrResponse{HTTPStatusCode: 401, StatusText: "Resource not found."}

func EndpointPostLinkConnection(w http.ResponseWriter, r *http.Request) {
	// Read project URL from request body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		render.Render(w, r, &ErrResponse{
			HTTPStatusCode: http.StatusInternalServerError,
			StatusText:     "Error reading request body.",
		})
		return
	}
	defer r.Body.Close() // Important: close the body after reading

	// Convert the byte slice to a string (or parse as JSON/XML etc.)
	bodyString := string(bodyBytes)
	// Print resulting string
	fmt.Printf("Received request body: %s\n", bodyString)

	// Quickly check if repository exists, otherwise error out and report
	err = p.CheckRemoteProject(bodyString)
	if err != nil {
		render.Render(w, r, &ErrResponse{
			HTTPStatusCode: http.StatusBadRequest,
			StatusText:     "Error checking remote project, check URL.",
		})
		return
	}

	// TODO: Add this code once queues are in place
	// Download repository to server
	err = p.HandleRemoteProject(bodyString)
	if err != nil {
		render.Render(w, r, &ErrResponse{
			HTTPStatusCode: http.StatusInternalServerError,
			StatusText:     "Error handling remote project, check remote.",
		})
		return
	}

	// Add repository to queue
	// TODO: refactor these functions
	err = p.HandleNewProject("./data/sources/"+path.Base(bodyString)+".git", "./data/sources/", "", "", "", "", "", "")
	// err = p.HandleNewProject("./data/store/linked/"+path.Base(bodyString), "./data/store/linked/", "", "", "", "", "", "")
	if err != nil {
		fmt.Println("ERROR handling new project!")
		fmt.Println(err)
		return
	}

	render.Render(w, r, &ErrResponse{HTTPStatusCode: http.StatusAccepted, StatusText: "Resource found and accepted for processing."})
	w.Write([]byte("Response successfull!"))
}

func EndpointGetProjects(w http.ResponseWriter, r *http.Request) {
	// TODO: create and pass enum to these functions regarding config paths
	file, err := os.Open(s.RepoConfig)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, "no projects", http.StatusNoContent)
			return
		}
		http.Error(w, "no projects", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	var projects s.Projects

	json.Unmarshal(b, &projects)

	if len(projects.Projects) < 1 {
		http.Error(w, "no projects", http.StatusNoContent)
		return
	}

	b, err = json.Marshal(projects)
	if err != nil {
		// TODO: is there a better error code?
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	// TODO: maybe want to use json.Indent to make json more readable?
	// TODO: add error checking to make sure data is valid before sending

	w.Write(b)
}

// Gets latest project zip of repository
func EndpointGetProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// TODO: Call function to update zip file of project
	project, ok := ctx.Value("project").(*s.Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	// TODO: Serve zip file
	// http.ServeFile(w, r, )
	w.Write([]byte(fmt.Sprintf("title:%s", project.Id)))
}

func EndpointGetProjectSchematics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	project, ok := ctx.Value("project").(*s.Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	// TODO: remove this in favor of calling during git push receive operation
	s := p.ProcessProjectFilePaths(s.RepoDir, project.ProjectFolder, project.Schematics)

	// TODO: obscure filename to increase security
	http.ServeFile(w, r, s[0])
	// TODO: maybe include metadata for parsing client side
	// w.Write([]byte(fmt.Sprintf("title:%s", project.Image)))
}

func EndpointGetProjectLayouts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	project, ok := ctx.Value("project").(*s.Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	// TODO: remove this in favor of calling during git push receive operation
	s := p.ProcessProjectFilePaths(s.RepoDir, project.ProjectFolder, project.Layouts)

	// TODO: obscure filename to increase security
	http.ServeFile(w, r, s[0])
	// TODO: maybe include metadata for parsing client side
	// w.Write([]byte(fmt.Sprintf("title:%s", project.Image)))
}

func EndpointGetProjectModels(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	project, ok := ctx.Value("project").(*s.Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	// TODO: check if any project.Models else return safe error
	// TODO: check for all files present

	s := p.ProcessProjectFilePaths(s.CacheGLBDir, project.ProjectFolder, project.Models)
	http.ServeFile(w, r, s[0])
	// w.Write([]byte(fmt.Sprintf("title:%s", project.Description)))
}

func EndpointGetRepository(w http.ResponseWriter, r *http.Request) {
	rp, err := filepath.Abs(s.RepoDir)
	if err != nil {
		fmt.Println(err)
	}

	bp, err := filepath.Abs(s.GitHTTPBackendExecutablePath)
	if err != nil {
		fmt.Println(err)
	}

	env := []string{
		fmt.Sprintf("GIT_PROJECT_ROOT=%s", rp),
		"GIT_HTTP_EXPORT_ALL=",
		"REMOTE_USER=githttp",
	}

	// debug git request
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	var stdErr bytes.Buffer
	handler := &cgi.Handler{
		Path:   bp,
		Root:   "/repos", // TODO: Does this do anything?
		Env:    env,
		Stderr: &stdErr,
	}
	http.StripPrefix("/api/git/project", handler).ServeHTTP(w, r)

	if stdErr.Len() > 0 {
		log.Println("[backend]", stdErr.String())
	}
}

// TODO: change from "add" to "enable"
func EndpointAddExampleProjects(w http.ResponseWriter, r *http.Request) {

	projects, err := s.ParseProjectsJSON(s.RepoConfigDemo)
	if err != nil {
		fmt.Println(err)
	}

	var prj *s.Project

	for _, a := range projects.Projects {
		// if a.RepositoryLink == "ext-con-breakout-board.git" {
		prj = &a
		// }

		// Handle new imported file
		// TODO: iterate over all example files
		// TODO: use ".git" folders and remove other folders
		fp := filepath.Join(s.DataDir, "examples", prj.ProjectFolder)

		// err = handleNewProjectArchive(fp, "./repos/repos")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		log.Print("Starting project: ")
		fmt.Println(prj)
		// fmt.Println(&prj.Id)

		// TODO: verify project integrity

		// pass newly created project to handle project function
		err = p.HandleNewProject(fp, s.RepoDir, prj.Id, prj.Image, prj.ProjectName, prj.ProjectFolder, prj.Description, prj.RepositoryLink)
		if err != nil {
			fmt.Println(err)
		}

	}

	// TODO: add response JSON
}

func EndpointRemoveExampleProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := s.ParseProjectsJSON(s.RepoConfigDemo)
	if err != nil {
		fmt.Println(err)
	}

	var prj *s.Project

	for _, a := range projects.Projects {
		// if a.RepositoryLink == "ext-con-breakout-board.git" {
		prj = &a

		p.HandleRemoveProject(prj.Id)
	}

	// TODO: add response JSON
}

func EndpointRemoveProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	project, ok := ctx.Value("project").(*s.Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	p.HandleRemoveProject(project.Id)

	// TODO: add response JSON

}

// TODO: This should be cached
func EndpointCheckExampleProjects(w http.ResponseWriter, r *http.Request) {
	examples, err := s.ParseProjectsJSON(s.RepoConfigDemo)
	if err != nil {
		fmt.Println(err)
	}

	projects, err := s.ParseProjectsJSON(s.RepoConfig)
	if err != nil {
		fmt.Println(err)
	}

	for _, a := range projects.Projects {
		for _, b := range examples.Projects {
			if a.Id == b.Id {
				fmt.Println(a)
				fmt.Println(b)
				w.Write([]byte("true"))
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("false"))
}
