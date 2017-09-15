package internal

import "time"

type Request struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
	Params  Params  `json:"params"`
}

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

type Params struct {
	Script string `json:"script"`
}
