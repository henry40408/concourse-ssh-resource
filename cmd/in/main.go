package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	easyssh "github.com/appleboy/easyssh-proxy"
	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/spacemonkeygo/errors"
)

const (
	SSHTimeout = 10 * 60 // = 10 minutes
)

var (
	ArgumentError = errors.NewClass("ArgumentError")
	FileError     = errors.NewClass("FileError")
	SSHError      = errors.NewClass("SSHError")
	TimeoutError  = errors.NewClass("TimeoutError")
)

type InRequest struct {
	Request internal.Request
}

func Main(stdin, stdout *os.File, args []string) error {
	var request InRequest

	if len(args) < 2 {
		return ArgumentError.New("need at least one argument")
	}

	err := internal.NewRequestFromStdin(stdin, &request.Request)
	if err != nil {
		return err
	}

	source := request.Request.Source
	config := &easyssh.MakeConfig{
		Server:   source.Host,
		Port:     "22",
		User:     source.User,
		Password: source.Password,
		Key:      source.PrivateKey,
	}

	if source.Port != 0 {
		config.Port = strconv.Itoa(source.Port)
	}

	params := request.Request.Params
	stdoutChan, stderrChan, doneChan, errChan, err := config.Stream(params.Script, SSHTimeout)
	if err != nil {
		return SSHError.New("failed to run command: %s", err.Error())
	}

	baseDir := args[1]

	outFile, err := os.OpenFile(path.Join(baseDir, "stdout"), os.O_RDWR|os.O_CREATE, 0644)
	defer outFile.Close()
	if err != nil {
		return FileError.New("failed to create file for SSH stdout: %s", err.Error())
	}

	errFile, err := os.OpenFile(path.Join(baseDir, "stderr"), os.O_RDWR|os.O_CREATE, 0644)
	defer errFile.Close()
	if err != nil {
		return FileError.New("failed to create file for SSH stderr: %s", err.Error())
	}

	done := true

loop:
	for {
		select {
		case done = <-doneChan:
			break loop
		case outline := <-stdoutChan:
			outFile.WriteString(outline)
		case errline := <-stderrChan:
			errFile.WriteString(errline)
		case err = <-errChan:
		}
	}

	if err != nil {
		return SSHError.New("failed when running SSH command: %s", err.Error())
	}

	if !done {
		return TimeoutError.New("SSH command is timeout")
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
