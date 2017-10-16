package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var response outResponse

	words := fake.WordsN(3)
	request := outRequest{
		Params: models.Params{
			Interpreter: "/bin/sh",
			Script:      fmt.Sprintf(`echo "%s"`, words),
		},
		Source: models.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
	}

	requestJSON, err := json.Marshal(&request)
	handleError(t, err)

	io, err := mockio.NewMockIO(requestJSON)
	defer io.Cleanup()
	handleError(t, err)

	err = outCommand(io.In, io.Out, io.Err)
	handleError(t, err)

	// test stdout
	stdoutContent, err := io.ReadAll(mockio.OUT)
	handleError(t, err)

	err = json.Unmarshal(stdoutContent, &response)
	handleError(t, err)

	assert.False(t, response.Version.Timestamp.IsZero())
	assert.Equal(t, 0, len(response.Metadata))

	// test stderr
	stderrContent, err := io.ReadAll(mockio.ERR)
	handleError(t, err)

	assert.Equal(t, fmt.Sprintf("stdout: %s\n", words), string(stderrContent))
}

func TestMainWithInterpreter(t *testing.T) {
	var response outResponse

	words := fake.WordsN(3)
	request := outRequest{
		Params: models.Params{
			Interpreter: "/usr/bin/python3",
			Script:      fmt.Sprintf(`print("%s")`, words),
		},
		Source: models.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
	}

	requestJSON, err := json.Marshal(&request)
	handleError(t, err)

	io, err := mockio.NewMockIO(requestJSON)
	defer io.Cleanup()
	handleError(t, err)

	err = outCommand(io.In, io.Out, io.Err)
	handleError(t, err)

	// test stdout
	stdoutContent, err := io.ReadAll(mockio.OUT)
	handleError(t, err)
	fmt.Printf(string(stdoutContent))

	err = json.Unmarshal(stdoutContent, &response)
	handleError(t, err)

	assert.False(t, response.Version.Timestamp.IsZero())
	assert.Equal(t, 0, len(response.Metadata))

	// test stderr
	stderrContent, err := io.ReadAll(mockio.ERR)
	handleError(t, err)

	assert.Equal(t, fmt.Sprintf("stdout: %s\n", words), string(stderrContent))
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
