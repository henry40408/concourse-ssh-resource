package main

import (
	"encoding/json"
	"io"
	"time"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	"github.com/henry40408/concourse-ssh-resource/internal/ssh"
	hierr "github.com/reconquest/hierr-go"
)

func outCommand(stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	var request outRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return hierr.Errorf(err, "unable to parse JSON from stdin")
	}

	stdoutWriter := &prefixWriter{
		prefix: "stdout",
		writer: stderr,
	}

	stderrWriter := &prefixWriter{
		prefix: "stderr",
		writer: stderr,
	}

	err = ssh.PerformSSHCommand(&request.Source, &request.Params, stdoutWriter, stderrWriter)
	if err != nil {
		return hierr.Errorf(err, "failed to run SSH command")
	}

	metadataItems := make([]models.Metadata, 0)
	response := outResponse{
		Version: models.Version{
			Timestamp: time.Now().Round(1 * time.Second),
		},
		Metadata: metadataItems,
	}

	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return hierr.Errorf(err, "failed to dump JSON to stdout")
	}

	return nil
}
