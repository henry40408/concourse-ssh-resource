package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMainWithEmptyVersion(t *testing.T) {
	var response CheckResponse

	stdout, err := ioutil.TempFile(os.TempDir(), "stdout")
	handleError(t, err)
	defer stdout.Close()

	stdin := strings.NewReader(`{ "source": {}, "version": {} }`)

	err = Main(stdin, stdout)
	handleError(t, err)

	stdout.Seek(0, 0)
	stdoutContent, err := ioutil.ReadAll(stdout)
	handleError(t, err)

	fmt.Printf(string(stdoutContent))

	err = json.Unmarshal(stdoutContent, &response)
	handleError(t, err)

	assert.False(t, response[0].Timestamp.IsZero())
}

func TestMainWithVersion(t *testing.T) {
	var response CheckResponse

	previousVersion := time.Now().Add(-1 * time.Second)

	stdout, err := ioutil.TempFile(os.TempDir(), "stdout")
	handleError(t, err)
	defer stdout.Close()

	stdin := strings.NewReader(fmt.Sprintf(`{
		"source": {},
		"version": {
			"time": "%s"
		}
	}`, previousVersion.Format(time.RFC3339)))

	err = Main(stdin, stdout)
	handleError(t, err)

	stdout.Seek(0, 0)
	stdoutContent, err := ioutil.ReadAll(stdout)
	handleError(t, err)

	fmt.Printf(string(stdoutContent))

	err = json.Unmarshal(stdoutContent, &response)
	handleError(t, err)

	assert.Equal(t, 2, len(response))
	assert.Equal(t, previousVersion.Unix(), response[0].Timestamp.Unix())
	assert.True(t, response[1].Timestamp.After(previousVersion))
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
