package main

import (
	"encoding/json"
	"io"

	"github.com/reconquest/hierr-go"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
)

type checkRequest struct {
	Source  models.Source  `json:"source"`
	Version models.Version `json:"version"`
}

type checkResponse []models.Version

func checkCommand(stdin io.Reader, stdout io.Writer) error {
	var request checkRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return hierr.Errorf(err, "unable to parse JSON from standard input")
	}

	response := checkResponse{}
	if (models.Version{}) != request.Version {
		response = append(response, request.Version)
	}

	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return hierr.Errorf(err, "unable to dump JSON to standard output")
	}

	return nil
}
