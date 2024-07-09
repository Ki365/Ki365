package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
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

	repoPath, err := createAbsoluteProjectFolderPath(archivePath, unarchivePath)
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

	rel, err := createRelativeProjectFolderPath(archivePath, unarchivePath)
	if err != nil {
		return err
	}

	projects, err := parseProjectsJSON()
	if err != nil {
		return err
	}

	fmt.Printf("repoPath: %v", rel)
	list_sch := find(rel, ".kicad_sch")
	list_pcb := find(rel, ".kicad_pcb")

	// TODO: maybe want to move away from JSON file store for project repos, (fixed project index right now)
	proj := Project{
		"100",
		"https://dummyimage.com/470x300/0F6983/fff&text=Temp+Image",
		strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath)),
		strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath)),
		"No description available",
		"",
		strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath)),
		list_sch,
		list_pcb,
		[]string{},
	}

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

// TODO: iterate over all in slice argument
func processProjectFilePaths(projectPath string, filepathSlice []string) []string {
	// TODO: make a global variable for consistent changes
	str := "./repos/repos"
	var s []string
	s = append(s, filepath.Join(str, projectPath, filepathSlice[0]))
	return s
}

func find(root, ext string) []string {
	var a []string
	var b []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	for _, e := range a {
		b = append(b, strings.TrimPrefix(e, filepath.Join(root)))
	}
	return b
}
