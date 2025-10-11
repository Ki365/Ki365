package kicad

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

// RequestGLTFModelFromKiCadCLI
// Input path must point to a file that exists.
// Output path does not require an existing file.
func RequestGLTFModelFromKiCadCLI(inputFile string, outputFilePath string, isDocker bool) error {
	if isDocker {
		return getGLTFModelFromDocker(inputFile, outputFilePath)
	} else {
		return getGLTFModelFromLocal(inputFile, outputFilePath)
	}
}

func getGLTFModelFromDocker(inputFile string, outputFilePath string) error {

	log.Println("Opening input file: " + inputFile)
	openedInputFile, err := os.Open(inputFile)
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

	// TODO: This should be modified to have some type of hashing algorithm to avoid similarly named projects
	// TODO: Need to include entire project to ensure local footprints are present for the exporter
	err = scpClient.CopyFromFile(context.Background(), *openedInputFile, filepath.Join(homePath, filepath.Base(inputFile)), "0655")
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
	err = session.Run("kicad-cli pcb export glb " + filepath.Base(inputFile) + " " +
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

	log.Println("Copying resulting GLTF...")
	err = scpClient.CopyFromRemote(context.Background(), of, filepath.Join(homePath, filepath.Base(outputFilePath)))
	if err != nil {
		fmt.Println("error in copying from remote")
		return err
	}

	// Successfully generated GLB file from KiCad CLI return no errors
	return nil
}

func getGLTFModelFromLocal(inputFile string, outputFilePath string) error {
	err := os.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm)
	if err != nil {
		fmt.Println("error in creating directories")
		return err
	}

	// outputFile, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	// if err != nil {
	// 	fmt.Println("error in opening new file")
	// 	return err
	// }

	fmt.Println("inputFile: " + inputFile)
	fmt.Println("inputFileBase: " + filepath.Base(inputFile))
	fmt.Println("outputFilePath: " + outputFilePath)
	fmt.Println("outputFilePathBase: " + filepath.Base(outputFilePath))

	cmdStringArgs := []string{"pcb", "export", "glb", inputFile,
		"--subst-models", "--include-tracks", "--include-pads", "--include-inner-copper",
		"--include-silkscreen", "--include-soldermask", "--include-zones", "-o", outputFilePath, "-f"}

	output, err := exec.Command("kicad-cli", cmdStringArgs...).Output()
	// This exit status is because of the missing components
	if err != nil && err.Error() != "exit status 2" {
		fmt.Print("\n\n\n")
		fmt.Println(cmdStringArgs)
		fmt.Print("\n\n\n")
		fmt.Println(string(output))
		fmt.Print("\n\n\n")
		fmt.Println("error in calling kicad-cli")
		fmt.Println(err)
		return err
	}

	// var out bytes.Buffer
	// var stderr bytes.Buffer
	// cmd.Stdout = &out
	// cmd.Stderr = &stderr
	// err = cmd.Run()
	// if err != nil {
	// 	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	// 	return err
	// }
	// fmt.Println("Result: " + out.String())

	// return errors.New("function unimplemented")
	return nil
}
