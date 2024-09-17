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
	for _, e := range entries {
		var cmd *exec.Cmd
		if bare {
			fmt.Print(e.Name() + " .git: ")
			cmd = exec.Command("git", "clone", "--bare", filepath.Join(source, e.Name()), filepath.Join(destination, e.Name()+".git"))
		} else {
			fmt.Print(e.Name() + " regu: ")
			cmd = exec.Command("git", "clone", filepath.Join(source, e.Name()), filepath.Join(destination, e.Name()))
		}
		_, err := cmd.Output()

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
