package main

import (
	"fmt"
	"io"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
)

type inResponse struct {
	Version  models.Version
	Metadata []models.Metadata
}

func inCommand(stdin io.Reader, stdout io.Writer) error {
	fmt.Fprintf(stdout, `{"version":{},"metadata":[]}`)
	return nil
}
