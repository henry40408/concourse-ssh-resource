package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/henry40408/concourse-ssh-resource/pkg"
)

type checkRequest struct {
	Source  pkg.Source  `json:"source"`
	Version pkg.Version `json:"version"`
}

type checkResponse []pkg.Version

func checkCommand(stdin io.Reader, stdout io.Writer) error {
	response := make(checkResponse, 0)
	err := json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return fmt.Errorf("unable to dump JSON to stdout: %v", err)
	}

	return nil
}

func main() {
	err := checkCommand(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
