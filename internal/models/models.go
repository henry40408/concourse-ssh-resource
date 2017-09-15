package models

import "time"

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type CheckResponse []Version

type Source struct {
	Host       string  `json:"host"`
	Port       float64 `json:"port"`
	User       string  `json:"user"`
	Password   string  `json:"password"`
	PrivateKey string  `json:"private_key"`
}

type Version struct {
	Timestamp time.Time `json:"time"`
}
