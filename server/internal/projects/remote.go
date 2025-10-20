package projects

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

func CheckRemoteProject(remote string) error {
	repo, err := git.Init(memory.NewStorage(), nil)

	// Create remotes in repo
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{remote},
	})
	if err != nil {
		fmt.Println("Failed to add remote to local repo.")
		return err
	}

	// Fetch repos and determine if it was a sucess
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
	})
	if err != nil {
		fmt.Println("Failed to fetch remote repository, probably an incorrect URL.")
		return err
	}

	return nil
}

func HandleRemoteProject(remote string) error {
	// Initial repo on disk
	// repo, err := git.PlainClone("./data/store/linked/"+path.Base(remote), false, &git.CloneOptions{
	repo, err := git.PlainClone("./data/sources/"+path.Base(remote), false, &git.CloneOptions{
		URL:               remote,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})
	if err != nil {
		fmt.Println("Failed to initialize git repo, check storage provider.")
		fmt.Println(err.Error())
		return nil
	}

	// List remotes for debugging logs
	list, err := repo.Remotes()
	if err != nil {
		fmt.Println("Failed to get remotes on local repo.")
		return nil
	}

	for _, r := range list {
		fmt.Println(r)
	}

	// List branches for debugging logs
	l, err := repo.Branches()
	if err != nil {
		fmt.Println("Failed to get branches on remote repo.")
		return nil
	}

	l.ForEach(func(r *plumbing.Reference) error {
		fmt.Println(r.Name())
		return nil
	})

	// Sucess, return no error
	return nil
}
