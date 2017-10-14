package main

import (
	"encoding/json"
	"testing"

	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var response checkResponse

	io, err := mockio.NewMockIO([]byte(`{ "source": {}, "version": {} }`))
	captureError(t, err)

	err = checkCommand(io.In, io.Out)
	captureError(t, err)

	stdoutContent, err := io.ReadAll(mockio.OUT)
	captureError(t, err)

	err = json.Unmarshal(stdoutContent, &response)
	captureError(t, err)

	assert.Equal(t, 0, len(response))
}

func captureError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
