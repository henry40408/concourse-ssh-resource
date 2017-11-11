package main

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"

	"github.com/stretchr/testify/assert"
)

func TestCheckCommand(t *testing.T) {
	var response []models.Version

	io, err := mockio.NewMockIO(bytes.NewBuffer([]byte(`{"source":{},"version":{}}`)))
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	err = checkCommand(io.In, io.Out)
	if !assert.NoError(t, err) {
		return
	}

	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 0, len(response))
}

func TestCheckCommandWithVersion(t *testing.T) {
	var response checkResponse

	now := time.Now().Format(time.RFC3339)

	request := checkRequest{Version: models.Version{Timestamp: now}}
	requestJSON, err := json.Marshal(&request)
	if !assert.NoError(t, err) {
		return
	}

	io, err := mockio.NewMockIO(bytes.NewBuffer(requestJSON))
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	err = checkCommand(io.In, io.Out)
	if !assert.NoError(t, err) {
		return
	}

	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	if assert.Equal(t, 1, len(response)) {
		assert.Equal(t, now, response[0].Timestamp)
	}
}
