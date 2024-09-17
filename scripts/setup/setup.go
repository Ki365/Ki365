package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
)

// var s = []string{
// 	"KI365_KICAD_PASSWORD",
// 	"KI365_TRACESPACE_PASSWORD",
// }

func main() {
	// Parameters to change

	err := copyFile("./.env.template", "./.env")
	if err != nil {
		fmt.Println(err)
	}
	err = randomizePasswords()
	if err != nil {
		fmt.Println(err)
	}

}

func copyFile(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dest, os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func randomizePasswords() error {
	// TODO: Change this to change file automatically
	// f, err := os.Open(s)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	fmt.Println("Generating random passwords:")
	for i := 0; i < 3; i++ {
		s, err := generatePassword(20, true, true, true)
		if err != nil {
			return err
		}
		fmt.Println(s)
	}
	return nil
}

func generatePassword(length int, useLetters bool, useSpecial bool, useNum bool) (string, error) {
	charset := ""
	var p []byte

	if useLetters {
		charset += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useSpecial {
		charset += "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	}
	if useNum {
		charset += "0123456789"
	}
	if charset == "" {
		return "", errors.New("must set one argument as true")
	}

	for i := 0; i < length; i++ {
		rand := rand.Intn(len(charset))
		p = append(p, charset[rand])
	}
	return string(p), nil
}
