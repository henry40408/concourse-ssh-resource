package main

import (
	"time"

	"github.com/henry40408/ssh-shell-resource/internal"
)

func CheckCommand(request *internal.CheckRequest) internal.CheckResponse {
	versions := internal.CheckResponse{}

	previousVersion := request.Version
	if !previousVersion.Timestamp.IsZero() {
		versions = append(versions, previousVersion)
	}

	versions = append(versions, internal.Version{Timestamp: time.Now()})
	return versions
}
