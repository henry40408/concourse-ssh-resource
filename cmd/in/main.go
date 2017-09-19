package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/henry40408/ssh-shell-resource/internal"
)

type InRequest struct {
	Source  internal.Source  `json:"source"`
	Version internal.Version `json:"version"`
	Params  internal.Params  `json:"params"`
}

type InResponse struct {
	Version  internal.Version    `json:"version"`
	Metadata []internal.Metadata `json:"metadata"`
}

func Main(stdin io.Reader, stdout io.Writer) error {
	var request InRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return fmt.Errorf("unable to parse JSON from stdin: %s", err.Error())
	}

	metadataItems := make([]internal.Metadata, 0)
	response := InResponse{
		Version:  request.Version,
		Metadata: metadataItems,
	}
	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return fmt.Errorf("failed to dump JSON to stdout: %s", err.Error())
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
