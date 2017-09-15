package main

import (
	"fmt"
	"os"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/spacemonkeygo/errors"
)

func Main(stdin, stdout *os.File) error {
	var request internal.CheckRequest

	err := internal.NewRequestFromStdin(stdin, &request)
	if err != nil {
		return err
	}

	response := CheckCommand(&request)

	err = internal.RespondToStdout(stdout, &response)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Main(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, errors.GetMessage(err))
		os.Exit(1)
	}
}
