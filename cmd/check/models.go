package main

import "github.com/henry40408/concourse-ssh-resource/internal/models"

type checkRequest struct {
	Source  models.Source  `json:"source"`
	Version models.Version `json:"version"`
}

type checkResponse []models.Version
