package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/icrowley/fake"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/henry40408/ssh-shell-resource/pkg/mockio"
)

func TestMain(t *testing.T) {
	var response OutResponse

	words := fake.WordsN(3)
	request := OutRequest{
		Params: internal.Params{
			Interpreter: "/bin/sh",
			Script:      fmt.Sprintf(`echo "%s"`, words),
		},
		Source: internal.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
	}

	requestJSON, err := json.Marshal(&request)
	handleError(t, err)

	io, err := mockio.NewMockIO(requestJSON)
	handleError(t, err)
	defer io.Cleanup()

	err = Main(io.In, io.Out, io.Err)
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

func TestMainWithPython(t *testing.T) {
	var response OutResponse

	words := fake.WordsN(3)
	request := OutRequest{
		Params: internal.Params{
			Interpreter: "/usr/bin/python3",
			Script:      fmt.Sprintf(`print("%s")`, words),
		},
		Source: internal.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
	}

	requestJSON, err := json.Marshal(&request)
	handleError(t, err)

	io, err := mockio.NewMockIO(requestJSON)
	handleError(t, err)
	defer io.Cleanup()

	err = Main(io.In, io.Out, io.Err)
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
