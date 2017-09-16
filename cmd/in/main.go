package main

import (
	"fmt"
	"os"
	"path"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/spacemonkeygo/errors"
)

const (
	SSHTimeout = 10 * 60 // = 10 minutes
)

type InRequest struct {
	Request internal.Request
}

func Main(stdin, stdout *os.File, args []string) error {
	var request InRequest

	if len(args) < 2 {
		return internal.ArgumentError.New("need at least one argument")
	}

	err := internal.NewRequestFromStdin(stdin, &request.Request)
	if err != nil {
		return err
	}

	baseDir := args[1]

	outFile, err := os.OpenFile(path.Join(baseDir, "stdout"), os.O_RDWR|os.O_CREATE, 0644)
	defer outFile.Close()
	if err != nil {
		return internal.FileError.New("failed to create file for SSH stdout: %s", err.Error())
	}

	errFile, err := os.OpenFile(path.Join(baseDir, "stderr"), os.O_RDWR|os.O_CREATE, 0644)
	defer errFile.Close()
	if err != nil {
		return internal.FileError.New("failed to create file for SSH stderr: %s", err.Error())
	}

	err = internal.PerformSSHCommand(&request.Request, outFile, errFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Main(os.Stdin, os.Stdout, os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, errors.GetMessage(err))
		os.Exit(1)
	}
}
