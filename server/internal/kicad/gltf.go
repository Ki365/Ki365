package kicad

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

// normalizeFilename returns only the filename without paths and any slashes (cross-platform)
func normalizeFilename(path string) string {
	// Replace all backslashes with forward slashes
	path = strings.ReplaceAll(path, "\\", "/")
	// Get only the filename
	name := filepath.Base(path)
	// Remove any forward slashes from the name (just in case)
	name = strings.ReplaceAll(name, "/", "")
	return name
}

// Input path must point to a file that exists.
// Output path does not require an existing file.
func RequestGLTFModelFromKiCadCLI(inputFile string, outputFilePath string) error {

	log.Println("Opening input file: " + inputFile)
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("error in opening")
		return err
	}

	homePath := "/home/kicad/"

	// TODO: move this functionality elsewhere to enable .env files
	pswd, wasSet := os.LookupEnv("KI365_KICAD_PASSWORD")
	if !wasSet {
		return fmt.Errorf("KI365_KICAD_PASSWORD was not set, cannot connect to KiCad instance")
	}

	config := &ssh.ClientConfig{
		User:            "kicad",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(pswd),
		},
	}

	// TODO: make localhost optional through command line argument
	log.Println("Dialing ssh client: localhost:2233")
	sshClient, err := ssh.Dial("tcp", ("localhost:2233"), config)
	if err != nil {
		fmt.Println("error in dialing")
		return err
	}
	defer sshClient.Close()

	scpClient, err := scp.NewClientBySSH(sshClient)
	if err != nil {
		fmt.Println("error in creating ssh client")
		return err
	}
	defer scpClient.Close()

	// Use normalized filename for copying
	err = scpClient.CopyFromFile(context.Background(), *f, filepath.Join(homePath, normalizeFilename(inputFile)), "0655")
	if err != nil {
		fmt.Println("error in copying file")
		return err
	}

	session, err := sshClient.NewSession()
	if err != nil {
		fmt.Println("error in ssh session")
		return err
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		fmt.Println("error in stdout pipe")
		return err
	}

	go func() {
		io.Copy(os.Stdout, stdout)
	}()
	err = session.Run("kicad-cli pcb export glb " + normalizeFilename(inputFile) + " " +
		"--subst-models --include-tracks --include-pads --include-inner-copper " +
		"--include-silkscreen --include-soldermask --include-zones -f")
	if err != nil {
		fmt.Println("error in kicad command")
		return err
	}

	err = os.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm)
	if err != nil {
		fmt.Println("error in creating directories")
		return err
	}

	of, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("error in opening new file")
		return err
	}

	// TODO: Check for a suitable binary in the path
	// NOTE: Unless KiCad CLI supports standalone binary executable, docker will be the only supported platform for this service
	log.Println("Copying resulting GLTF...")
	err = scpClient.CopyFromRemote(context.Background(), of, filepath.Join(homePath, normalizeFilename(outputFilePath)))
	if err != nil {
		fmt.Println("error in copying from remote")
		return err
	}

	// Successfully generated GLB file from KiCad CLI return no errors
	return nil
}
