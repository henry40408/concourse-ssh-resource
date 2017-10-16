package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	hierr "github.com/reconquest/hierr-go"
)

func inCommand(stdin io.Reader, stdout io.Writer) error {
	var request inRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return hierr.Errorf(err, "unable to parse JSON from stdin")
	}

	metadataItems := make([]models.Metadata, 0)
	response := inResponse{
		Version:  request.Version,
		Metadata: metadataItems,
	}
	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return hierr.Errorf(err, "failed to dump JSON to stdout")
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
