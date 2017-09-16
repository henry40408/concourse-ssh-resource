package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/icrowley/fake"
	"github.com/spacemonkeygo/errors"
	"github.com/stretchr/testify/assert"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/henry40408/ssh-shell-resource/pkg/mockio"
)

type mockmockio struct {
	in  *os.File
	out *os.File
}

const scriptTemplate = `
#!/bin/sh
echo -e '%s'
`

func TestMain(t *testing.T) {
	words := fake.WordsN(3)
	request := internal.Request{
		Source: internal.Source{
			Host:     "localhost",
			Port:     22,
			User:     "root",
			Password: "toor",
		},
		Params: internal.Params{
			Script: fmt.Sprintf(scriptTemplate, words),
		},
	}

	args := []string{"", os.TempDir()}

	content, err := json.Marshal(&request)
	if err != nil {
		t.Error(err)
	}

	mockio, err := mockio.NewMockIO(content)
	defer mockio.Cleanup()
	if err != nil {
		t.Error(err)
	}

	err = Main(mockio.In, mockio.Out, args)
	if err != nil {
		t.Error(err)
	}

	outFile, err := stdoutFile()
	defer cleanupFile(outFile)
	if err != nil {
		t.Error(err)
	}

	reader := bufio.NewReader(outFile)
	outContent, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, words, string(outContent))

	errFile, err := stderrFile()
	defer cleanupFile(errFile)
	if err != nil {
		t.Error(err)
	}

	reader = bufio.NewReader(errFile)
	errContent, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "", string(errContent))
}

func TestMainWithNoArguments(t *testing.T) {
	args := []string{""}
	err := Main(nil, nil, args)
	assert.Equal(t, "ArgumentError: need at least one argument", errors.GetMessage(err))
}

func TestMainWithMalformedJSON(t *testing.T) {
	args := []string{"", os.TempDir()}

	content := []byte("{")

	mockio, err := mockio.NewMockIO(content)
	defer mockio.Cleanup()
	if err != nil {
		t.Error(err)
	}

	err = Main(mockio.In, mockio.Out, args)
	assert.Equal(t, "InvalidJSONError: stdin is not a valid JSON", errors.GetMessage(err))
}

func TestMainFailToRunSSHCommand(t *testing.T) {
	words := fake.WordsN(3)
	request := internal.Request{
		Source: internal.Source{
			Host:     "localhost",
			Port:     22,
			User:     "root",
			Password: "root",
		},
		Params: internal.Params{
			Script: fmt.Sprintf(scriptTemplate, words),
		},
	}

	args := []string{"", os.TempDir()}

	content, err := json.Marshal(&request)
	if err != nil {
		t.Error(err)
	}

	mockio, err := mockio.NewMockIO(content)
	defer mockio.Cleanup()
	if err != nil {
		t.Error(err)
	}

	err = Main(mockio.In, mockio.Out, args)
	assert.Equal(t, internal.SSHError, errors.GetClass(err))
}

func stdoutFile() (*os.File, error) {
	return os.Open(path.Join(os.TempDir(), "stdout"))
}

func stderrFile() (*os.File, error) {
	return os.Open(path.Join(os.TempDir(), "stderr"))
}

func cleanupFile(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}
