package optimization

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func OptimizeGLB(inputFile string) error {
	abs, err := filepath.Abs(inputFile)
	if err != nil {
		fmt.Println("error in resolving input name")
	}

	// EX: ./repos/cache/glb/{project}/example.glb -> ./repos/cache/glb/{project}/example-old.glb
	newAbs := filepath.Join(filepath.Dir(abs), strings.TrimSuffix(filepath.Base(abs), filepath.Ext(abs))+"-opt"+filepath.Ext(abs))
	oldAbs := filepath.Join(filepath.Dir(abs), strings.TrimSuffix(filepath.Base(abs), filepath.Ext(abs))+"-old"+filepath.Ext(abs))

	args := []string{"-cc", "-tc", "-mi", "-i", abs, "-o", newAbs} //, "-si 1"}
	// args = append(args, )+newAbs
	// args = append(args, newAbs)+
	// args = append(args, "-o"+abs)
	fmt.Print("\n\n\n")
	fmt.Println(args)
	fmt.Print("\n\n\n")
	// fmt.Println("-o " + abs)
	output, err := exec.Command("./bin/gltfpack", args...).Output()
	if err != nil {
		fmt.Print("\n\n\n")
		fmt.Println(string(output))
		fmt.Print("\n\n\n")
		fmt.Println("error in calling gltfpack")
		fmt.Println(err)
		return err
	}

	err = os.Rename(abs, oldAbs)
	if err != nil {
		fmt.Println("error in renaming unoptimized file")
		// panic("rename")
		return err
	}

	err = os.Rename(newAbs, abs)
	if err != nil {
		fmt.Println("error in renaming optimized file")
		// panic("rename")
		return err
	}

	// panic("stop")

	return nil
}
