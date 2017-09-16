package main

import (
	"fmt"
	"os"
	"time"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/spacemonkeygo/errors"
)

type CheckRequest struct {
	internal.Request
}

type CheckResponse []internal.Version

func CheckCommand(request *CheckRequest) CheckResponse {
	versions := CheckResponse{}

	previousVersion := request.Request.Version
	if !previousVersion.Timestamp.IsZero() {
		versions = append(versions, previousVersion)
	}

	versions = append(versions, internal.Version{Timestamp: time.Now()})
	return versions
}

func Main(stdin, stdout *os.File) error {
	var request CheckRequest

	err := internal.NewRequestFromStdin(stdin, &request.Request)
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
