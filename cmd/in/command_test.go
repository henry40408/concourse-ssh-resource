package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
)

type inResponse struct {
	Version  models.Version    `json:"version"`
	Metadata []models.Metadata `json:"metadata"`
}

func TestInCommand(t *testing.T) {
	var response inResponse

	in := bytes.NewBufferString(`{"source":{},"version":{}}`)
	out := bytes.NewBuffer([]byte{})
	err := inCommand(in, out)
	if !assert.NoError(t, err) {
		return
	}

	// test standard output
	err = json.NewDecoder(out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	assert.Empty(t, response.Version.Timestamp)
	assert.Empty(t, response.Metadata)
}
