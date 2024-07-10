package main

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
	"path/filepath"
)

func EndpointGetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func EndpointGetProjects(w http.ResponseWriter, r *http.Request) {
	// TODO: create and pass enum to these functions regarding config paths
	file, err := os.Open("./repos/store/repos.json")
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

	var projects Projects

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
	project, ok := ctx.Value("project").(*Project)
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
	project, ok := ctx.Value("project").(*Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	// TODO: remove this in favor of calling during git push receive operation
	s := processProjectFilePaths(project.ProjectFolder, project.Schematics)

	// TODO: obscure filename to increase security
	http.ServeFile(w, r, s[0])
	// TODO: maybe include metadata for parsing client side
	// w.Write([]byte(fmt.Sprintf("title:%s", project.Image)))
}

func EndpointGetProjectLayouts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	project, ok := ctx.Value("project").(*Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	// TODO: remove this in favor of calling during git push receive operation
	s := processProjectFilePaths(project.ProjectFolder, project.Layouts)

	// TODO: obscure filename to increase security
	http.ServeFile(w, r, s[0])
	// TODO: maybe include metadata for parsing client side
	// w.Write([]byte(fmt.Sprintf("title:%s", project.Image)))
}

func EndpointGetProjectModels(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	project, ok := ctx.Value("project").(*Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	w.Write([]byte(fmt.Sprintf("title:%s", project.Description)))
}

var (
	// WARNING: This binary only correctly compiles and runs on x86-64 linux.
	// TODO: enable correct binary compilation based on running OS
	git_http_backend_bin = "./bin/git-http-backend"
	git_repos            = "./repos/repos"
)

func EndpointGetRepository(w http.ResponseWriter, r *http.Request) {
	rp, err := filepath.Abs(git_repos)
	if err != nil {
		fmt.Println(err)
	}

	bp, err := filepath.Abs(git_http_backend_bin)
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
		Root:   "/repos",
		Env:    env,
		Stderr: &stdErr,
	}
	http.StripPrefix("/api/git/project", handler).ServeHTTP(w, r)

	if stdErr.Len() > 0 {
		log.Println("[backend]", stdErr.String())
	}
}

func EndpointAddExampleProjects(w http.ResponseWriter, r *http.Request) {

	projects, err := parseProjectsJSON("./repos/store/repos-demo.json")
	if err != nil {
		fmt.Println(err)
	}

	var prj *Project

	for _, a := range projects.Projects {
		// if a.RepositoryLink == "ext-con-breakout-board.git" {
		prj = &a
		// }

		// Handle new imported file
		// TODO: iterate over all example files
		fp := filepath.Join("./repos/examples/" + prj.ProjectFolder + ".zip")

		// err = handleNewProjectArchive(fp, "./repos/repos")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		fmt.Println(prj)
		fmt.Println(prj.Id)
		fmt.Println(&prj.Id)

		// TODO: verify project integrity

		// pass newly created project to handle project function
		err = handleNewProject(fp, "./repos/repos", prj.Id, prj.Image, prj.ProjectName, prj.ProjectFolder, prj.Description, prj.RepositoryLink)
		if err != nil {
			fmt.Println(err)
		}

	}

	// TODO: add response JSON
}

func EndpointRemoveExampleProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := parseProjectsJSON("./repos/store/repos-demo.json")
	if err != nil {
		fmt.Println(err)
	}

	var prj *Project

	for _, a := range projects.Projects {
		// if a.RepositoryLink == "ext-con-breakout-board.git" {
		prj = &a

		handleRemoveProject(prj.Id)
	}

	// TODO: add response JSON
}

func EndpointRemoveProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	project, ok := ctx.Value("project").(*Project)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	handleRemoveProject(project.Id)

	// TODO: add response JSON

}

// TODO: This should be cached
func EndpointCheckExampleProjects(w http.ResponseWriter, r *http.Request) {
	examples, err := parseProjectsJSON("./repos/store/repos-demo.json")
	if err != nil {
		fmt.Println(err)
	}

	projects, err := parseProjectsJSON("./repos/store/repos.json")
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
