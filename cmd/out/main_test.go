package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/icrowley/fake"

	"github.com/henry40408/ssh-shell-resource/internal"
)

func TestMain(t *testing.T) {
	var response OutResponse

	tempDir := os.TempDir()

	words := fake.WordsN(3)
	request := OutRequest{
		Params: internal.Params{
			Script: fmt.Sprintf(`
#!/bin/sh
echo "%s"
`, words),
		},
		Source: internal.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
	}

	requestJSON, err := json.Marshal(&request)
	handleError(t, err)

	stdin := bytes.NewReader(requestJSON)

	stderr, err := ioutil.TempFile(tempDir, "stderr")
	handleError(t, err)
	defer stderr.Close()

	stdout, err := ioutil.TempFile(tempDir, "stdout")
	handleError(t, err)
	defer stdout.Close()

	err = Main(stdin, stdout, stderr)
	handleError(t, err)

	// test stdout
	stdout.Seek(0, 0)
	stdoutContent, err := ioutil.ReadAll(stdout)
	handleError(t, err)

	fmt.Printf(string(stdoutContent))

	err = json.Unmarshal(stdoutContent, &response)
	handleError(t, err)

	assert.False(t, response.Version.Timestamp.IsZero())
	assert.Equal(t, 0, len(response.Metadata))

	// test stderr
	stderr.Seek(0, 0)
	stderrContent, err := ioutil.ReadAll(stderr)
	handleError(t, err)

	assert.Equal(t, fmt.Sprintf("stdout: %s", words), string(stderrContent))
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
