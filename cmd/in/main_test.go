package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var response InResponse

	tempDir := os.TempDir()
	args := []string{"", tempDir}

	previousVersion := time.Now().Add(-1 * time.Second).Round(1 * time.Second)
	words := fake.WordsN(3)

	request := InRequest{
		Source: internal.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
		Version: internal.Version{
			Timestamp: previousVersion,
		},
		Params: internal.Params{
			Script: fmt.Sprintf(`
#!/bin/sh
echo "%s"
`, words),
		},
	}

	requestJSON, err := json.Marshal(&request)
	handleError(t, err)

	stdin := bytes.NewReader(requestJSON)

	stdout, err := ioutil.TempFile(tempDir, "stdout")
	handleError(t, err)
	defer stdout.Close()

	err = Main(stdin, stdout, args)
	handleError(t, err)

	// test stdout file
	stdoutFile, err := os.OpenFile(stdoutFilePath(tempDir), os.O_RDONLY, 0644)
	handleError(t, err)
	defer cleanUp(stdoutFile)
	stdoutContent, err := ioutil.ReadAll(stdoutFile)
	handleError(t, err)

	assert.Equal(t, words, string(stdoutContent))

	// test stderr file
	stderrFile, err := os.OpenFile(stderrFilepath(tempDir), os.O_RDONLY, 0644)
	handleError(t, err)
	defer cleanUp(stderrFile)
	stderrContent, err := ioutil.ReadAll(stderrFile)
	handleError(t, err)

	assert.Equal(t, "", string(stderrContent))

	// test stdout
	stdout.Seek(0, 0)
	responseJSON, err := ioutil.ReadAll(stdout)
	handleError(t, err)

	err = json.Unmarshal(responseJSON, &response)
	handleError(t, err)

	fmt.Printf(string(responseJSON))

	assert.Equal(t, previousVersion.Unix(), response.Version.Timestamp.Unix())
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func cleanUp(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}
