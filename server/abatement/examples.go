package abatement

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-git/go-git/v5"
)

// Generates example projects with customizable source and destination folders.
// Clones source git projects into destination folder.
// If initial is enabled, generates bare and basic clones of repositories.
func GenerateExamples(source string, destination string, initial bool) {

	// Get list of all example projects (includes)
	entries, err := os.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: check if folder vs file
	// TODO: check if folder is submodule

	var wg sync.WaitGroup

	for _, e := range entries {
		// TODO: Add progress bar for large repos
		wg.Add(1)
		go func() {
			defer wg.Done()
			if initial {
				copyRepo(source, destination, e.Name(), true, true)
				copyRepo(source, destination, e.Name(), false, false)

			} else {
				t := filepath.Ext(e.Name()) == ".git"
				copyRepo(source, destination, e.Name(), t, false)
			}
		}()

	}
	wg.Wait()
}

// Copies a repo from source to destination parent folder with folder name.
// Bare controls whether clone should be a bare git repository.
// Append adds ".git" to the destination folder name.
func copyRepo(source, destination, name string, bare bool, append bool) {

	var destFullPath string
	if append {
		destFullPath = filepath.Join(destination, name+".git")
	} else {
		destFullPath = filepath.Join(destination, name)
	}

	_, err := git.PlainClone(destFullPath, bare, &git.CloneOptions{
		URL: filepath.Join(source, name),
	})

	var s string

	if bare {
		s = "bare/.git"
	} else {
		s = "non-bare/regular"
	}

	if err != nil {
		fmt.Printf("Error cloning %s as %s from %s\n", name, s, err.Error())
	} else {
		fmt.Println("Cloned " + filepath.Base(name) + " as " + s)
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
