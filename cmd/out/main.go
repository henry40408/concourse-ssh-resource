package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/henry40408/ssh-shell-resource/internal"
)

type OutRequest struct {
	Params internal.Params `json:"params"`
	Source internal.Source `json:"source"`
}

type OutResponse struct {
	Version  internal.Version    `json:"version"`
	Metadata []internal.Metadata `json:"metadata"`
}

type prefixWriter struct {
	prefix string
	writer io.Writer
}

func (pw *prefixWriter) Write(p []byte) (n int, err error) {
	return fmt.Fprintf(pw.writer, "%s: %s", pw.prefix, p)
}

func Main(stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	var request OutRequest

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

	err = internal.PerformSSHCommand(&request.Source, &request.Params, stdoutWriter, stderrWriter)
	if err != nil {
		return fmt.Errorf("failed to run SSH command: %v", err)
	}

	metadataItems := make([]internal.Metadata, 0)
	response := OutResponse{
		Version: internal.Version{
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
	err := Main(os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
