package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"
)

type inResponse struct {
	Version  models.Version    `json:"version"`
	Metadata []models.Metadata `json:"metadata"`
}

func TestInCommand(t *testing.T) {
	var response inResponse

	reader := bytes.NewBuffer([]byte(`{"source":{},"version":{}}`))
	io, err := mockio.NewMockIO(reader)
	if !assert.NoError(t, err) {
		return
	}

	err = inCommand(io.In, io.Out)
	if !assert.NoError(t, err) {
		return
	}

	// test standard output
	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	assert.Empty(t, response.Version.Timestamp)
	assert.Empty(t, response.Metadata)
}
