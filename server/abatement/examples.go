package abatement

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Generates example projects with customizable source and destination folders.
// Clones source git projects into destination folder.
func GenerateExamples(source string, destination string, bare bool) {
	// TODO: make sure to check git submodules
	// TODO: use go git "github.com/go-git/go-git/v5"

	// Get list of all example projects (includes)
	entries, err := os.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: check if folder vs file
	// TODO: check if folder is submodule

	// Iterate all entries to build .git folders

	var git_cmd string

	// TODO: Make global
	_, err = os.Stat("./bin/git")
	if err != nil {
		git_cmd = "git"
	} else {
		git_cmd = "./bin/git"
	}

	var cmd *exec.Cmd

	// NOTE: May need to call to disable safe directory (required for "shared" drives like WSL):
	// git config --global --add safe.directory '*'

	// TODO: maybe convert to go-git:
	// _, err = git.PlainClone(filepath.Join(destination, e.Name()+".git"), true, &git.CloneOptions{
	// 	URL: filepath.Join(source, e.Name()),
	// })

	for _, e := range entries {
		if bare {
			fmt.Print(e.Name() + " .git: ")
			cmd = exec.Command(git_cmd, "clone", "--bare", filepath.Join(source, e.Name()), filepath.Join(destination, e.Name()+".git"))
		} else {
			fmt.Print(e.Name() + " regu: ")
			cmd = exec.Command(git_cmd, "clone", filepath.Join(source, e.Name()), filepath.Join(destination, e.Name()))
		}
		s, err := cmd.CombinedOutput()

		fmt.Println(string(s[:]))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println()
		}
	}
}

func CopyManifest(source string, destination string) {
	from, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}
