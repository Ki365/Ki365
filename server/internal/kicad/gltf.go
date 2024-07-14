package kicad

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

// Input path must point to a file that exists.
// Output path does not require an existing file.
func RequestGLTFModelFromKiCadCLI(inputFile string, outputFilePath string) error {

	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("error in opening")
		return err
	}

	homePath := "/home/kicad/"

	config := &ssh.ClientConfig{
		User:            "kicad",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(""), // TODO: Include from env configuration
		},
	}

	sshClient, err := ssh.Dial("tcp", ("172.18.0.2:22"), config)
	if err != nil {
		fmt.Println("error in dialing")
		return err
	}
	fmt.Println(sshClient)
	defer sshClient.Close()

	scpClient, err := scp.NewClientBySSH(sshClient)
	if err != nil {
		fmt.Println("error in creating ssh client")
		return err
	}
	defer scpClient.Close()

	err = scpClient.CopyFromFile(context.Background(), *f, filepath.Join(homePath, filepath.Base(inputFile)), "0655")
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

	var b bytes.Buffer
	session.Stdout = &b

	fmt.Println(f.Name())

	err = session.Run("kicad-cli pcb export glb " + filepath.Base(inputFile) + " " +
		"--subst-models --include-tracks --include-pads --include-inner-copper " +
		"--include-silkscreen --include-soldermask --include-zones -f")
	if err != nil {
		fmt.Println("error in kicad command")
		return err
	}
	fmt.Println(b.String())

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

	// NOTE: Unless KiCad CLI supports standalone binary executable, docker will be the only supported platform for this service
	err = scpClient.CopyFromRemote(context.Background(), of, filepath.Join(homePath, filepath.Base(outputFilePath)))
	if err != nil {
		fmt.Println("error in copying from remote")
		return err
	}

	// Successfully generated GLB file from KiCad CLI return no errors
	return nil
}
