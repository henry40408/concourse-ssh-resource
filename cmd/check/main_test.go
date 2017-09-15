package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spacemonkeygo/errors"
	"github.com/stretchr/testify/assert"
)

type MockStdio struct {
	In    *os.File
	Out   *os.File
	Error error
}

func TestMain(t *testing.T) {
	stdio, err := NewMockStdio()
	if err != nil {
		t.Error(err)
	}
	defer stdio.Cleanup()

	stdio.In.WriteString("{}")
	stdio.In.Seek(0, 0)

	err = Main(stdio.In, stdio.Out)
	if err != nil {
		t.Error(err)
	}
}

func TestMainNotValidJSON(t *testing.T) {
	stdio, err := NewMockStdio()
	if err != nil {
		t.Error(err)
	}
	defer stdio.Cleanup()

	err = Main(stdio.In, stdio.Out)
	assert.Equal(t, errors.GetMessage(err), "InvalidJSONError: stdin is not a valid JSON")
}

func NewMockStdio() (*MockStdio, error) {
	stdin, err := ioutil.TempFile(os.TempDir(), "stdin")
	if err != nil {
		return nil, err
	}

	stdout, err := ioutil.TempFile(os.TempDir(), "stdout")
	if err != nil {
		return nil, err
	}

	return &MockStdio{In: stdin, Out: stdout}, nil
}

func (m *MockStdio) Cleanup() {
	os.Remove(m.In.Name())
	os.Remove(m.Out.Name())
}
