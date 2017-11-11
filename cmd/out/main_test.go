package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var response outResponse

	words := fake.WordsN(3)

	request, err := json.Marshal(&outRequest{
		Params: models.Params{
			Interpreter: "/bin/sh",
			Script:      fmt.Sprintf(`echo "%s"`, words),
		},
		Source: models.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
	})
	assert.NoError(t, err)

	io, err := mockio.NewMockIO(bytes.NewBuffer(request))
	defer io.Cleanup()
	assert.NoError(t, err)

	err = outCommand(io.In, io.Out, io.Err)
	assert.NoError(t, err)

	// test standard output
	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	assert.NoError(t, err)

	assert.NotEmpty(t, response.Version.Timestamp)
	assert.Equal(t, 0, len(response.Metadata))

	// test standard error
	io.Err.Seek(0, 0)
	stderrContent, err := ioutil.ReadAll(io.Err)
	assert.NoError(t, err)

	assert.Equal(t, fmt.Sprintf("stdout: %s\n", words), string(stderrContent))
}

func TestMainWithInterpreter(t *testing.T) {
	var response outResponse

	words := fake.WordsN(3)
	request, err := json.Marshal(&outRequest{
		Params: models.Params{
			Interpreter: "/usr/bin/python3",
			Script:      fmt.Sprintf(`print("%s")`, words),
		},
		Source: models.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
	})
	assert.NoError(t, err)

	io, err := mockio.NewMockIO(bytes.NewBuffer(request))
	defer io.Cleanup()
	assert.NoError(t, err)

	err = outCommand(io.In, io.Out, io.Err)
	assert.NoError(t, err)

	// test standard output
	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	assert.NoError(t, err)

	assert.NotEmpty(t, response.Version.Timestamp)
	assert.Equal(t, 0, len(response.Metadata))

	// test standard error
	io.Err.Seek(0, 0)
	stderrContent, err := ioutil.ReadAll(io.Err)
	assert.NoError(t, err)

	assert.Equal(t, fmt.Sprintf("stdout: %s\n", words), string(stderrContent))
}
