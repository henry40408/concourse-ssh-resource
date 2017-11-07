package main

import (
	"encoding/json"
	"testing"

	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var response struct {
		Version  interface{}   `json:"version"`
		Metadata []interface{} `json:"metadata"`
	}

	io, err := mockio.NewMockIO([]byte(`{"source":{},"version":{}}`))
	assert.NoError(t, err)

	err = inCommand(io.In, io.Out)
	assert.NoError(t, err)

	// test stdout
	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	assert.NoError(t, err)

	assert.Empty(t, response.Version)
	assert.Empty(t, response.Metadata)
}
