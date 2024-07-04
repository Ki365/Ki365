package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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
