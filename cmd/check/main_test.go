package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/spacemonkeygo/errors"
	"github.com/stretchr/testify/assert"
)

type MockStdio struct {
	In    *os.File
	Out   *os.File
	Error error
}

func TestCheckCommandReturnDifferentResponse(t *testing.T) {
	request := CheckRequest{}

	response := CheckCommand(&request)
	assert.Equal(t, 1, len(response))

	time.Sleep(1 * time.Millisecond)

	anotherResponse := CheckCommand(&request)
	assert.Equal(t, 1, len(anotherResponse))

	responseTime := response[0].Timestamp.UnixNano()
	anotherResponseTime := anotherResponse[0].Timestamp.UnixNano()
	assert.NotEqual(t, responseTime, anotherResponseTime)
}

func TestCheckCommandReturnPreviousVersion(t *testing.T) {
	version := internal.Version{Timestamp: time.Now()}
	request := CheckRequest{
		Request: internal.Request{Version: version},
	}

	time.Sleep(1 * time.Millisecond)

	response := CheckCommand(&request)
	assert.Equal(t, 2, len(response))

	requestTime := request.Version.Timestamp.UnixNano()
	responseTime := response[0].Timestamp.UnixNano()
	assert.Equal(t, requestTime, responseTime)
}

func TestCheckCommandResponseTimeIsGreaterThanRequestTime(t *testing.T) {
	version := internal.Version{Timestamp: time.Now()}
	request := CheckRequest{
		Request: internal.Request{Version: version},
	}

	time.Sleep(1 * time.Millisecond)

	response := CheckCommand(&request)

	requestTime := request.Version.Timestamp.UnixNano()
	responseTime := response[1].Timestamp.UnixNano()
	assert.True(t, responseTime > requestTime)
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
