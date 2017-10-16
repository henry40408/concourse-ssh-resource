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
	handleError(t, err)

	io, err := mockio.NewMockIO(requestJSON)
	handleError(t, err)

	err = inCommand(io.In, io.Out)
	handleError(t, err)

	// test stdout
	responseJSON, err := io.ReadAll(mockio.OUT)
	handleError(t, err)

	err = json.Unmarshal(responseJSON, &response)
	handleError(t, err)

	assert.Empty(t, response.Metadata)
	assert.True(t, (models.Version{}) == response.Version)
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
