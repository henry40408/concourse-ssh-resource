package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/icrowley/fake"
)

func TestMain(t *testing.T) {
	var response InResponse

	words := fake.WordsN(3)
	request := InRequest{
		Source: internal.Source{
			Host:     "localhost",
			User:     "root",
			Password: "toor",
		},
		Version: internal.Version{},
		Params: internal.Params{
			Script: fmt.Sprintf(`
#!/bin/sh
echo "%s"
`, words),
		},
	}

	requestJSON, err := json.Marshal(&request)
	handleError(t, err)

	stdin := bytes.NewReader(requestJSON)

	stdout, err := ioutil.TempFile(os.TempDir(), "stdout")
	handleError(t, err)
	defer stdout.Close()

	err = Main(stdin, stdout)
	handleError(t, err)

	// test stdout
	stdout.Seek(0, 0)
	responseJSON, err := ioutil.ReadAll(stdout)
	handleError(t, err)

	err = json.Unmarshal(responseJSON, &response)
	handleError(t, err)

	fmt.Printf(string(responseJSON))

	assert.Empty(t, response.Metadata)
	assert.True(t, (internal.Version{}) == response.Version)
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func cleanUp(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}
