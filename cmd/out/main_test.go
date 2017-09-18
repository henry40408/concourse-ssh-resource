package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/henry40408/ssh-shell-resource/pkg/mockio"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

const scriptTemplate = `
#!/bin/sh
echo "%s"
`

func TestMain(t *testing.T) {
	var response internal.Request

	words := fake.WordsN(3)
	request := internal.Request{
		Source: internal.Source{
			Host:     "localhost",
			Port:     22,
			User:     "root",
			Password: "toor",
		},
		Params: internal.Params{
			Script: fmt.Sprintf(scriptTemplate, words),
		},
	}

	requestJSON, err := json.Marshal(&request)
	if err != nil {
		t.Error(err)
	}

	mockio, err := mockio.NewMockIO(requestJSON)
	if err != nil {
		t.Error(err)
	}

	err = Main(mockio.In, mockio.Out, mockio.Err)
	if err != nil {
		t.Error(err)
	}

	mockio.Err.Seek(0, 0)

	reader := bufio.NewReader(mockio.Err)
	errContent, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}

	expected := fmt.Sprintf("stdout: %s", words)
	assert.Equal(t, expected, string(errContent))

	mockio.Out.Seek(0, 0)
	responseJSON, err := ioutil.ReadAll(mockio.Out)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(responseJSON, &response)
	if err != nil {
		t.Error(err)
	}

	assert.False(t, response.Version.Timestamp.IsZero())
}
