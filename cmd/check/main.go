package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/henry40408/ssh-shell-resource/internal"
)

type CheckRequest struct {
	Source  internal.Source  `json:"source"`
	Version internal.Version `json:"version"`
}

type CheckResponse []internal.Version

func Main(stdin io.Reader, stdout io.Writer) error {
	response := make(CheckResponse, 0)
	err := json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return fmt.Errorf("unable to dump JSON to stdout: %v", err)
	}

	return nil
}

func main() {
	err := Main(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
