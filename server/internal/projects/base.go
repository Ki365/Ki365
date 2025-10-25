package projects

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ki365/ki365/server/internal/kicad"
	"github.com/ki365/ki365/server/internal/optimization"
	s "github.com/ki365/ki365/server/structure"

	bolt "go.etcd.io/bbolt"
)

// TODO: iterate over all in slice argument
func ProcessProjectFilePaths(basepath string, projectPath string, filepathSlice []string) []string {
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

func HandleNewProject(archivePath string, unarchivePath string, projectInst s.Project) error {

	// TODO: Check that repository in archivePath is a valid repository

	rel, err := createRelativeProjectFolderPath(archivePath, unarchivePath)
	if err != nil {
		return err
	}

	projects, err := s.ParseProjectsJSON(s.Dirs().Store.RepoConfig)
	if err != nil {
		return err
	}

	fmt.Printf("repoPath: %v\n", rel)
	list_sch := find(rel, ".kicad_sch")
	list_pcb := find(rel, ".kicad_pcb")

	list_glb := []string{}
	for _, p := range list_pcb {
		// TODO: each folder should have a unique generated cache dir
		s := filepath.Join(strings.TrimSuffix(filepath.Base(p), filepath.Ext(p)) + ".glb")
		list_glb = append(list_glb, s)
	}

	if projectInst.Id == "" {
		db := s.OpenDB()
		defer db.Close()

		err := db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("Increments"))
			if err != nil {
				return err
			}
			a, _ := b.NextSequence()
			fmt.Println(a)
			projectInst.Id = "P-" + strconv.Itoa(int(a))

			return nil
		})
		if err != nil {
			fmt.Println("Fatal database error:  " + err.Error())
			return err
		}
	}

	if projectInst.ProjectName == "" {
		projectInst.ProjectName = strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath))
	}

	if projectInst.ProjectFolder == "" {
		projectInst.ProjectFolder = strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath))
	}

	if projectInst.Description == "" {
		projectInst.Description = "No description available"
	}

	if projectInst.RepositoryLink == "" {
		projectInst.RepositoryLink = strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath))
	}

	if projectInst.Image == "" {
		projectInst.Image = "https://dummyimage.com/470x300/0F6983/fff&text=Temp+Image"
	}

	// TODO: maybe want to move away from JSON file store for project repos, (fixed project index right now)
	projectInst.Schematics = list_sch
	projectInst.Layouts = list_pcb
	projectInst.Models = list_glb

	projects.Projects = append(projects.Projects, projectInst)

	// TODO: move long processes like zip and generating files to asynchronous operations
	// TODO: frontend: expose multi model boards to the interface

	//  TODO: iterate over range of glb paths
	err = kicad.RequestGLTFModelFromKiCadCLI(
		"./"+filepath.Join(s.Dirs().RepoDir, projectInst.ProjectFolder, list_pcb[0]),
		"./"+filepath.Join(s.Dirs().Cache.GLBDir, projectInst.ProjectFolder, filepath.Base(list_glb[0])),
		s.UseDockerKiCadCLI)
	if err != nil {
		return err
	}

	optimization.OptimizeGLB("./"+filepath.Join(s.Dirs().Cache.GLBDir, projectInst.ProjectFolder, filepath.Base(list_glb[0])), s.GLTFPackExecutablePath)

	// Marshal new data
	fmt.Println("Marshalling JSON")
	b, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println("Writing to file")
	err = os.WriteFile(s.Dirs().Store.RepoConfig, b, 0755)
	if err != nil {
		return err
	}
	// encoder.Encode(b)
	fmt.Println(string(b))

	// new project is valid, return no error
	return nil
}

func HandleRemoveProject(id string) error {
	// Load in projects from JSON store
	projects, err := s.ParseProjectsJSON(s.Dirs().Store.RepoConfig)
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
	err = os.WriteFile(s.Dirs().Store.RepoConfig, b, 0755)
	if err != nil {
		return err
	}
	// encoder.Encode(b)
	fmt.Println(string(b))

	// project removal was successful, return no error
	return nil
}
