package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/reconquest/hierr-go"
	"github.com/spf13/afero"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	"github.com/henry40408/concourse-ssh-resource/internal/ssh"
)

type outRequest struct {
	Params models.Params `json:"params"`
	Source models.Source `json:"source"`
}

type outResponse struct {
	Version  models.Version    `json:"version"`
	Metadata []models.Metadata `json:"metadata"`
}

func outCommand(fs afero.Fs, args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	var request outRequest

	if len(args) < 2 {
		return fmt.Errorf("need base directory, usage: %s <base directory>", args[0])
	}

	baseDir := args[1]

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return hierr.Errorf(err, "unable to parse JSON from standard input")
	}

	outWriter := &prefixWriter{prefix: "STDOUT", writer: stderr}
	errWriter := &prefixWriter{prefix: "STDERR", writer: stderr}
	err = ssh.PerformSSHCommand(fs, &request.Source, &request.Params, outWriter, errWriter, baseDir)
	if err != nil {
		return hierr.Errorf(err, "unable to run SSH command")
	}

	response := outResponse{
		Version: models.Version{
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Metadata: make([]models.Metadata, 0),
	}

	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return hierr.Errorf(err, "unable to dump JSON to standard output")
	}

	return nil
}
