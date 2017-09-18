package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

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

func Main(stdin io.Reader, stdout io.Writer, args []string) error {
	var request InRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return fmt.Errorf("unable to parse JSON from stdin: %s", err.Error())
	}

	if len(args) < 2 {
		return fmt.Errorf("need at least one argument")
	}

	baseDir := args[1]

	stdoutFile, err := os.OpenFile(stdoutFilePath(baseDir), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file for SSH stdout: %s", err.Error())
	}
	defer stdoutFile.Close()

	stderrFile, err := os.OpenFile(stderrFilepath(baseDir), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file for SSH stderr: %s", err.Error())
	}
	defer stderrFile.Close()

	err = internal.PerformSSHCommand(&request.Source, &request.Params, stdoutFile, stderrFile)
	if err != nil {
		return fmt.Errorf("failed to run SSH command: %s", err.Error())
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
	err := Main(os.Stdin, os.Stdout, os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func stdoutFilePath(baseDir string) string {
	return path.Join(baseDir, "stdout")
}

func stderrFilepath(baseDir string) string {
	return path.Join(baseDir, "stderr")
}
