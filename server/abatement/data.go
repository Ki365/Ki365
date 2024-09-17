package abatement

import (
	"os"
	"path/filepath"
)

func GenerateDataFolder(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	s := []string{
		"cache/repos/",
		"examples/",
		"releases/",
		"source/",
		"store/repos/",
		"upload/archives/",
	}

	for i := 0; i < len(s); i++ {
		p := filepath.Join(path, s[i])
		if err != nil {
			return err
		}
		err = os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
