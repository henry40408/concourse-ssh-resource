package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
	hierr "github.com/reconquest/hierr-go"
)

func checkCommand(stdin io.Reader, stdout io.Writer) error {
	var response checkResponse

	version := models.Version{}
	response = append(response, version)

	err := json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return hierr.Errorf(err, "unable to dump JSON to stdout")
	}

	return nil
}

func main() {
	err := checkCommand(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
