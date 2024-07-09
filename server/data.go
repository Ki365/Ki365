package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Id               string   `json:"ID"`
	Image            string   `json:"Image"`
	ProjectName      string   `json:"ProjectName"`
	ProjectFolder    string   `json:"ProjectFolder"`
	Description      string   `json:"Description"`
	ShortDescription string   `json:"Short Description"` // TODO delete in favor of dynamically generated short description on react
	RepositoryLink   string   `json:"RepositoryLink"`
	Schematics       []string `json:"Schematics`
	Layouts          []string `json:"Layouts`
	Models           []string `json:"Models`
}

var RepoConfig = "./repos/store/repos.json"

type PathStruct struct {
	RepoConfig string
}

func parseProjectsJSON() (*Projects, error) {
	// repoConfig: ./repos/store/repos.json + unarchivepath: ./repos/repos"

	// TODO: change os.Executable() for repo directory
	expath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	t1 := filepath.Join(expath, RepoConfig) // absolute path
	t2 := filepath.Dir(t1)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(t2, 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(RepoConfig, os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read in current json store
	// TODO: make this process able to just "add" project without reading in entire list
	// TODO: move this to a seperate function along with similar code in endpoints
	by, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var projects Projects

	json.Unmarshal(by, &projects)

	return &projects, nil
}
