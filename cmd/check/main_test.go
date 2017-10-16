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
	assert.NoError(t, err)

	err = checkCommand(io.In, io.Out)
	assert.NoError(t, err)

	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(response))
	assert.Zero(t, response[0].Timestamp)
}
