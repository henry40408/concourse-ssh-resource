package main

import "github.com/henry40408/concourse-ssh-resource/internal/models"

type outRequest struct {
	Params models.Params `json:"params"`
	Source models.Source `json:"source"`
}

type outResponse struct {
	Version  models.Version    `json:"version"`
	Metadata []models.Metadata `json:"metadata"`
}
