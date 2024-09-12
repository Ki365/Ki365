package abatement

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Generates example projects with customizable source and destination folders
// Clones source git projects into destination folder
func GenerateExamples(source string, destination string) {
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
		fmt.Print(e.Name() + " .git: ")
		cmd := exec.Command("git", "clone", "--bare", source+e.Name(), destination+e.Name()+".git")
		_, err := cmd.Output()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println()
		}

		// TODO: remove this block once caching data is done

		fmt.Print(e.Name() + " regu: ")
		cmd = exec.Command("git", "clone", source+e.Name(), destination+e.Name())
		_, err = cmd.Output()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println()
		}
	}
}
