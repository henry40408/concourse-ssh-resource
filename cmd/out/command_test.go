package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/afero"

	"github.com/icrowley/fake"
	"github.com/reconquest/hierr-go"
	"github.com/stretchr/testify/assert"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
)

func TestOutCommand(t *testing.T) {
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
	if !assert.NoError(t, err) {
		return
	}

	args := []string{"out", "/tmp"}
	in := bytes.NewBuffer(request)
	out := bytes.NewBuffer([]byte{})
	stdErr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	err = outCommand(fs, args, in, out, stdErr)
	if !assert.NoError(t, err) {
		return
	}

	// test standard output
	err = json.NewDecoder(out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	assert.NotEmpty(t, response.Version.Timestamp)
	assert.Equal(t, 0, len(response.Metadata))

	// test standard error
	stderrContent, err := ioutil.ReadAll(stdErr)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, fmt.Sprintf("STDOUT: %s\n", words), string(stderrContent))
}

func TestOutCommandWithInterpreter(t *testing.T) {
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
	if !assert.NoError(t, err) {
		return
	}

	args := []string{"out", "/tmp"}
	in := bytes.NewBuffer(request)
	out := bytes.NewBuffer([]byte{})
	stdErr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	err = outCommand(fs, args, in, out, stdErr)
	if !assert.NoError(t, err) {
		return
	}

	// test standard output
	err = json.NewDecoder(out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	assert.NotEmpty(t, response.Version.Timestamp)
	assert.Equal(t, 0, len(response.Metadata))

	// test standard error
	stderrContent, err := ioutil.ReadAll(stdErr)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, []byte(fmt.Sprintf("STDOUT: %s\n", words)), stderrContent)
}

func TestOutCommandWithMalformedJSON(t *testing.T) {
	args := []string{"out", "/tmp"}
	in := bytes.NewBufferString(`{`)
	out := bytes.NewBuffer([]byte{})
	stdErr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	err := outCommand(fs, args, in, out, stdErr)

	herr := err.(hierr.Error)
	assert.Equal(t, herr.GetMessage(), "unable to parse JSON from standard input")
}

func TestOutCommandWithBadConnectionInfo(t *testing.T) {
	request, err := json.Marshal(&outRequest{
		Params: models.Params{
			Interpreter: "/bin/sh",
			Script:      "uptime",
		},
		Source: models.Source{
			Host:     "localhost",
			User:     "root",
			Password: "",
		},
	})
	if !assert.NoError(t, err) {
		return
	}

	args := []string{"out", "/tmp"}
	in := bytes.NewBuffer(request)
	out := bytes.NewBuffer([]byte{})
	stdErr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	err = outCommand(fs, args, in, out, stdErr)
	herr := err.(hierr.Error)
	assert.Equal(t, herr.GetMessage(), "unable to run SSH command")
}

func TestOutCommandWithNoBaseDirectory(t *testing.T) {
	args := []string{"out"}

	in := bytes.NewBuffer([]byte{})
	out := bytes.NewBuffer([]byte{})
	stdErr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	err := outCommand(fs, args, in, out, stdErr)

	if !assert.Error(t, err) {
		return
	}

	herr := err.(hierr.Error)
	assert.Equal(t, herr.Error(), "need base directory, usage: out <base directory>")
}
