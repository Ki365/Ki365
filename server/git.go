package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/ki365/ki365/kicad"
	"github.com/ki365/ki365/optimization"
)

type ErrorRepoNotClean struct {
	When time.Time
	What string
}

func (e *ErrorRepoNotClean) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func handleNewProject(archivePath string, unarchivePath string, id string, image string, prjName string, prjFolder string, desc string, prjLink string) error {

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

	// s, err := w.Status()
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(s)

	// fmt.Println(s.IsClean())
	// fmt.Println(!s.IsClean())
	// if !s.IsClean() {
	// 	return &ErrorRepoNotClean{
	// 		time.Now(),
	// 		"repository is not clean",
	// 	}
	// }

	rel, err := createRelativeProjectFolderPath(archivePath, unarchivePath)
	if err != nil {
		return err
	}

	projects, err := parseProjectsJSON(RepoConfig)
	if err != nil {
		return err
	}

	fmt.Printf("repoPath: %v", rel)
	list_sch := find(rel, ".kicad_sch")
	list_pcb := find(rel, ".kicad_pcb")

	list_glb := []string{}
	for _, p := range list_pcb {
		// TODO: each folder should have a unique generated cache dir
		s := filepath.Join(strings.TrimSuffix(filepath.Base(p), filepath.Ext(p)) + ".glb")
		list_glb = append(list_glb, s)
	}

	var projectName string

	if prjName != "" {
		projectName = prjName
	} else {
		projectName = strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath))
	}

	var projectFolder string

	if prjFolder != "" {
		projectFolder = prjFolder
	} else {
		strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath))
	}

	var description string

	if desc != "" {
		description = desc
	} else {
		description = "No description available"
	}

	var projectLink string

	if prjLink != "" {
		projectLink = prjLink
	} else {
		projectLink = strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath))
	}

	var imageLink string

	if image != "" {
		imageLink = image
	} else {
		imageLink = "https://dummyimage.com/470x300/0F6983/fff&text=Temp+Image"
	}

	// TODO: maybe want to move away from JSON file store for project repos, (fixed project index right now)
	proj := Project{
		id,
		imageLink,
		projectName,
		projectFolder,
		description,
		"",
		projectLink,
		list_sch,
		list_pcb,
		list_glb, // TODO: add a isgenerating flag to notify users of progress
	}

	projects.Projects = append(projects.Projects, proj)

	// TODO: move long processes like zip and generating files to asynchronous operations
	// TODO: frontend: expose multi model boards to the interface

	//  TODO: iterate over range of glb paths
	err = kicad.RequestGLTFModelFromKiCadCLI(
		"./"+filepath.Join(RepoDir, projectFolder, list_pcb[0]),
		"./"+filepath.Join(CacheGLBDir, projectFolder, filepath.Base(list_glb[0])))
	if err != nil {
		return err
	}

	optimization.OptimizeGLB("./" + filepath.Join(CacheGLBDir, projectFolder, filepath.Base(list_glb[0])))

	// Marshal new data
	fmt.Println("Marshalling JSON")
	b, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println("Writing to file")
	err = os.WriteFile(RepoConfig, b, 0755)
	if err != nil {
		return err
	}
	// encoder.Encode(b)
	fmt.Println(string(b))

	// new project is valid, return no error
	return nil
}

func handleRemoveProject(id string) error {
	// Load in projects from JSON store
	projects, err := parseProjectsJSON(RepoConfig)
	if err != nil {
		return err
	}

	// Search project repository list for project
	var proj int
	i := 0
	for ; i < len(projects.Projects); i++ {
		if projects.Projects[i].Id == id {
			proj = i
			break
		}
	}

	// catch no find
	if i == len(projects.Projects) {
		return errors.New("no project with such id found")
	}

	// Delete while keeping slice order in tack
	projects.Projects = append(projects.Projects[:proj], projects.Projects[proj+1:]...)

	// Marshal new data
	fmt.Println("Marshalling JSON")
	b, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}

	// Writing to file
	fmt.Println("Writing to file")
	err = os.WriteFile(RepoConfig, b, 0755)
	if err != nil {
		return err
	}
	// encoder.Encode(b)
	fmt.Println(string(b))

	// project removal was successful, return no error
	return nil
}

// TODO: iterate over all in slice argument
func processProjectFilePaths(basepath string, projectPath string, filepathSlice []string) []string {
	var s []string
	s = append(s, filepath.Join(basepath, projectPath, filepathSlice[0]))
	fmt.Println(s)
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
