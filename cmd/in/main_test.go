package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/henry40408/ssh-shell-resource/pkg/mockio"
)

func TestMain(t *testing.T) {
	var response inResponse

	request := inRequest{
		Source: internal.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
		Version: internal.Version{},
		Params:  internal.Params{},
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
	assert.True(t, (internal.Version{}) == response.Version)
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
