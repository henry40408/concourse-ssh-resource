package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/henry40408/ssh-shell-resource/internal"
)

type CheckRequest struct {
	Source  internal.Source  `json:"source"`
	Version internal.Version `json:"version"`
}

type CheckResponse []internal.Version

func Main(stdin io.Reader, stdout io.Writer) error {
	var request CheckRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return fmt.Errorf("unable to parse JSON from stdin: %s", err.Error())
	}

	response := make(CheckResponse, 0)
	if !request.Version.Timestamp.IsZero() {
		response = append(response, request.Version)
	}
	response = append(response, internal.Version{
		Timestamp: time.Now().Round(1 * time.Second),
	})

	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return fmt.Errorf("unable to dump JSON to stdout: %s", err.Error())
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
