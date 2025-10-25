package structure

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Id               string   `json:"ID"`
	Image            string   `json:"Image"`
	ProjectName      string   `json:"ProjectName"`
	ProjectFolder    string   `json:"ProjectFolder"`
	Description      string   `json:"Description"`
	ShortDescription string   `json:"Short Description"` // TODO delete in favor of dynamically generated short description on react
	RepositoryLink   string   `json:"RepositoryLink"`
	Schematics       []string `json:"Schematics"`
	Layouts          []string `json:"Layouts"`
	Models           []string `json:"Models"`
}

type data struct {
	DataDir string
	RepoDir string
	Cache   cache
	Store   store
}

type cache struct {
	Dir    string
	GLBDir string
}

type store struct {
	Dir        string
	ExManifest string
	RepoConfig string
}

const dataDir = "./data"
const repoDir = "sources"
const storeDir = "store"
const cacheDir = "cache"
const gLBDir = "glb"

const repoConfig = "repos.json"
const repoManifest = "manifest-examples.json"

// Dirs() generates a nested structure of complete relative filepaths
// for an individual organization.
//
// The organization directories are setup as the following:
//
// - dataDir
//
//	1 - sources
//	2 - examples
//	3 - releases
//	4 - store
//	5 - upload
//	6 - cache
//	  	- glb
//
// TODO: Rename to Data
func Dirs() data {
	return data{
		DataDir: filepath.Join(dataDir),
		RepoDir: filepath.Join(dataDir, repoDir),
		Cache: cache{
			Dir:    filepath.Join(dataDir, cacheDir),
			GLBDir: filepath.Join(dataDir, cacheDir, gLBDir),
		},
		Store: store{
			Dir:        filepath.Join(dataDir, storeDir),
			ExManifest: filepath.Join(dataDir, storeDir, repoManifest),
			RepoConfig: filepath.Join(dataDir, storeDir, repoConfig),
		},
	}
}

// Conditional prefix changes executable locations for when local bin flag is set
// TODO: Refactor using interfaces
var ConditionalBinPrefix = "./bin/"
var GLTFPackExecutablePath = "gltfpack"
var GitHTTPBackendExecutablePath = "git-http-backend"
var TracespaceExecutablePath = "tracespace"

var UseDockerKiCadCLI = false

type PathStruct struct {
	RepoConfig string
}

func ParseProjectsJSON(conf string) (*Projects, error) {
	// repoConfig: ./repos/store/repos.json + unarchivepath: ./repos/repos"

	// TODO: change os.Executable() for repo directory
	expath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	t1 := filepath.Join(expath, conf) // absolute path
	t2 := filepath.Dir(t1)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(t2, 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(conf, os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read in current json store
	// TODO: make this process able to just "add" project without reading in entire list
	// TODO: move this to a seperate function along with similar code in endpoints
	by, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var projects Projects

	json.Unmarshal(by, &projects)

	return &projects, nil
}
