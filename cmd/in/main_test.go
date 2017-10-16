package main

import (
	"encoding/json"
	"testing"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var response inResponse

	request := inRequest{
		Source: models.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
		Version: models.Version{},
		Params:  models.Params{},
	}

	requestJSON, err := json.Marshal(&request)
	assert.NoError(t, err)

	io, err := mockio.NewMockIO(requestJSON)
	assert.NoError(t, err)

	err = inCommand(io.In, io.Out)
	assert.NoError(t, err)

	// test stdout
	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	assert.NoError(t, err)

	assert.Empty(t, response.Metadata)
	assert.Equal(t, models.Version{}, response.Version)
}
