package main

import (
	"encoding/json"
	"errors"
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
