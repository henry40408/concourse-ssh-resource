package main

import "github.com/henry40408/concourse-ssh-resource/internal/models"

type inRequest struct {
	Source  models.Source  `json:"source"`
	Version models.Version `json:"version"`
	Params  models.Params  `json:"params"`
}

type inResponse struct {
	Version  models.Version    `json:"version"`
	Metadata []models.Metadata `json:"metadata"`
}
