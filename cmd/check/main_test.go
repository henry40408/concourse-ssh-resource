package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
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

	assert.Equal(t, 0, len(response))
}

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
