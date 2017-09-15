package main

import "github.com/henry40408/ssh-shell-resource/internal"

type CheckRequest struct {
	internal.Request
}

type CheckResponse []internal.Version
