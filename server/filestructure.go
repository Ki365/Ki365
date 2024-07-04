package main

import (
	"archive/tar"
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func handleNewProjectArchive(archivePath string, unarchivePath string) error {
	// TODO: move seperate file types as dedicated functions with defer keyword on file

	// unzip tar or zip file to repos/

	// check if file exists
	_, err := os.Stat(archivePath)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}

	fmt.Println(archivePath)
	fmt.Println(unarchivePath)
	// EXAMPLE: repos\upload\archives\ki-102.zip + ./repos/repos = C:\Users\{USER}\Code\Ki365\Ki365\repos\repos\ki-102
	t3, err := createAbsoluteProjectFolderPath(archivePath, unarchivePath)
	if err != nil {
		return err
	}

	fmt.Printf("Unarchiving project to: %s \n", t3)

	err = os.MkdirAll(t3, 0755)
	if err != nil {
		return err
	}

	// TODO: check functionality of untarring
	// TODO: add automated checks to these functions
	// check and try to untar archive
	af, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	// defer file close for odd cases
	defer af.Close()

	tr := tar.NewReader(af)
	_, err = tr.Next()

	if err == nil {
		// is tar file open new tar reader
		tr = tar.NewReader(af)

		// TODO: maybe add support for gzip compression
		for {
			hdr, err := tr.Next()

			// error checking
			switch {
			case err == io.EOF:
				// end of file, exit
				return nil
			case err != nil:
				return err
			case hdr == nil:
				continue
			}
			// create file/directory location path
			target := filepath.Join(path.Dir(archivePath), hdr.Name)
			switch hdr.Typeflag {
			// directory, create if does not exist
			case tar.TypeDir:
				if _, err := os.Stat(target); err != nil {
					if err := os.MkdirAll(target, 0755); err != nil {
						return err
					}
				}
			// file, create and copy data
			case tar.TypeReg:
				// open/create file with same permissions as source
				f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
				if err != nil {
					return err
				}
				if _, err := io.Copy(f, tr); err != nil {
					return err
				}
				f.Close()
			}
		}
	}

	err = af.Close()
	if err != nil {
		return err
	}

	// file was not a tar file try zip
	zr, err := zip.OpenReader(archivePath)
	if err == nil {
		// defer zr.Close()
		for _, f := range zr.File {
			fmt.Printf("Unzipping %s\n", f.Name)

			rc, err := f.Open()
			if err != nil {
				return err
			}
			// defer rc.Close()

			target := filepath.Join(t3, f.Name) // full path of future file

			if f.FileInfo().IsDir() {
				err = os.MkdirAll(target, 0755)
				if err != nil {
					return err
				}
				continue
			}

			f, err := os.Create(target)
			if err != nil {
				return err
			}
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	} else {
		// file was not any of the above file types, return err
		// TODO: create err which is more suitable
		return err

	}

	// unarchive successful, return no error
	return nil
}

func createAbsoluteProjectFolderPath(archivePath string, unarchivePath string) (projectPath string, err error) {

	// TODO: change os.Executable() for repo directory
	expath, err := os.Getwd()
	if err != nil {
		return "!", err
	}
	// create project folder path
	t1 := filepath.Join(expath, unarchivePath)                                      // absolute path
	t2 := strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath)) // archive filename without suffix
	return filepath.Join(t1, t2), nil                                               // project folder
}

func createRelativeProjectFolderPath(archivePath string, unarchivePath string) (projectPath string, err error) {

	// TODO: change os.Executable() for repo directory
	// create project folder path
	t2 := strings.TrimSuffix(filepath.Base(archivePath), filepath.Ext(archivePath)) // archive filename without suffix
	return filepath.Join(unarchivePath, t2), nil                                    // project folder
}
