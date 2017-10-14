package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/henry40408/concourse-ssh-resource/internal"
)

type inRequest struct {
	Source  internal.Source  `json:"source"`
	Version internal.Version `json:"version"`
	Params  internal.Params  `json:"params"`
}

type inResponse struct {
	Version  internal.Version    `json:"version"`
	Metadata []internal.Metadata `json:"metadata"`
}

func inCommand(stdin io.Reader, stdout io.Writer) error {
	var request inRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return fmt.Errorf("unable to parse JSON from stdin: %v", err)
	}

	metadataItems := make([]internal.Metadata, 0)
	response := inResponse{
		Version:  request.Version,
		Metadata: metadataItems,
	}
	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return fmt.Errorf("failed to dump JSON to stdout: %v", err)
	}

	return nil
}

func main() {
	err := inCommand(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
