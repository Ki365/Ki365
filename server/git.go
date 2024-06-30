package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
)

type ErrorRepoNotClean struct {
	When time.Time
	What string
}

func (e *ErrorRepoNotClean) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func handleNewProject(archivePath string, unarchivePath string, repoConfig string) error {

	repoPath, err := createProjectFolderPath(archivePath, unarchivePath)
	if err != nil {
		return err
	}

	fmt.Println(archivePath)
	fmt.Println(unarchivePath)
	fmt.Println(repoPath)

	// new project is a repository and has no uncommited files (clean)
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	fmt.Println(w)

	s, err := w.Status()
	if err != nil {
		return err
	}

	fmt.Println(s)

	fmt.Println(s.IsClean())
	fmt.Println(!s.IsClean())
	if !s.IsClean() {
		return &ErrorRepoNotClean{
			time.Now(),
			"repository is not clean",
		}
	}

	// repoConfig: ./repos/store/repos.json + unarchivepath: ./repos/repos"

	// TODO: change os.Executable() for repo directory
	expath, err := os.Getwd()
	if err != nil {
		return err
	}

	t1 := filepath.Join(expath, repoConfig) // absolute path
	t2 := filepath.Dir(t1)
	if err != nil {
		return err
	}

	err = os.MkdirAll(t2, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(repoConfig, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read in current json store
	// TODO: make this process able to just "add" project without reading in entire list
	// TODO: move this to a seperate function along with similar code in endpoints
	by, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var projects Projects

	json.Unmarshal(by, &projects)

	// TODO: maybe want to move away from JSON file store for project repos, (fixed project index right now)
	proj := Project{100, "https://dummyimage.com/470x300/0F6983/fff&text=Temp+Image", strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath)), "No description available", ""}

	projects.Projects = append(projects.Projects, proj)

	// Marshal new data
	fmt.Println("Marshalling JSON")
	b, err := json.Marshal(projects)
	if err != nil {
		return err
	}

	fmt.Println("Writing to file")
	err = os.WriteFile(repoConfig, b, 0755)
	if err != nil {
		return err
	}
	// encoder.Encode(b)
	fmt.Println(string(b))

	// new project is valid, return no error
	return nil
}
