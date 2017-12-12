package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/spf13/afero"

	"github.com/stretchr/testify/assert"
)

func TestOutCommandWithValuePlaceholder(t *testing.T) {
	args := []string{"out", "/tmp"}
	in := bytes.NewBufferString(`{
		"source": {
			"host": "localhost",
			"user": "root",
			"password": "toor"
		},
		"params": {
			"interpreter": "/bin/sh",
			"script": "echo <CODE_VERSION>",
			"placeholders": [{
				"name": "<CODE_VERSION>",
				"value": "test_ok"
			}]
		},
		"version": { "ref": "0.5.0" }
	}`)
	out := bytes.NewBuffer([]byte{})
	stdErr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	err := outCommand(fs, args, in, out, stdErr)
	if !assert.NoError(t, err) {
		return
	}

	stdErrContent, err := ioutil.ReadAll(stdErr)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, []byte("STDOUT: test_ok\n"), stdErrContent)
}

func TestOutCommandWithFilePlaceholder(t *testing.T) {
	args := []string{"out", "/tmp"}
	in := bytes.NewBufferString(`{
		"source": {
			"host": "localhost",
			"user": "root",
			"password": "toor"
		},
		"params": {
			"interpreter": "/bin/sh",
			"script": "echo <CODE_VERSION>",
			"placeholders": [{
				"name": "<CODE_VERSION>",
				"file": "somefile"
			}]
		},
		"version": { "ref": "0.5.0" }
	}`)
	out := bytes.NewBuffer([]byte{})
	stdErr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	fs.MkdirAll("tmp", 0755)
	afero.WriteFile(fs, "/tmp/somefile", []byte("first_line\nsecond_line\nthird_line"), 0644)

	err := outCommand(fs, args, in, out, stdErr)
	if !assert.NoError(t, err) {
		return
	}

	stdErrContent, err := ioutil.ReadAll(stdErr)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, []byte("STDOUT: first_line\n"), stdErrContent)
}
