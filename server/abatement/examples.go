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
func GenerateExamples(source string, destination string) {

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

			t := filepath.Ext(e.Name()) == ".git"

			_, err = git.PlainClone(filepath.Join(destination, e.Name()), t, &git.CloneOptions{
				URL: filepath.Join(source, e.Name()),
			})

			var s string

			if t {
				s = ".git"
			} else {
				s = "regular"
			}

			if err != nil {
				fmt.Printf("Error cloning %s as %s from %s\n", filepath.Base(e.Name()), s, err)
			} else {
				fmt.Println("Cloned " + filepath.Base(e.Name()) + " as " + s)
			}
		}()

	}
	wg.Wait()
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
