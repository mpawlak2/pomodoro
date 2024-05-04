package cli

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

var FailedToGetInputError = errors.New("failed to get input")

func GetInputFromTextEditor() (string, error) {
	filename := "POMODORO_MESSAGE.txt"
	editor := "vim"

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", FailedToGetInputError
	}

	// read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", FailedToGetInputError
	}

	return string(content), nil
}
