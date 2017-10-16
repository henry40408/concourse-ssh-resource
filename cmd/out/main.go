package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	"github.com/henry40408/concourse-ssh-resource/internal/ssh"
)

func outCommand(stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	var request outRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return fmt.Errorf("unable to parse JSON from stdin: %v", err)
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
		return fmt.Errorf("failed to run SSH command: %v", err)
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
		return fmt.Errorf("failed to dump JSON to stdout: %v", err)
	}

	return nil
}

func main() {
	err := outCommand(os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
