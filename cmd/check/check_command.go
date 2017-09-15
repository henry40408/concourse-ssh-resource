package main

import (
	"time"

	"github.com/henry40408/ssh-shell-resource/internal/models"
)

func CheckCommand(request *models.CheckRequest) models.CheckResponse {
	versions := models.CheckResponse{}

	previousVersion := request.Version
	if !previousVersion.Timestamp.IsZero() {
		versions = append(versions, previousVersion)
	}

	versions = append(versions, models.Version{Timestamp: time.Now()})
	return versions
}
