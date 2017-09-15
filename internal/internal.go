package internal

import (
	"encoding/json"
	"os"
	"time"

	"github.com/spacemonkeygo/errors"
)

var (
	InvalidJSONError = errors.NewClass("InvalidJSONError")
	OutputError      = errors.NewClass("OutputError")
)

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

func NewRequestFromStdin(stdin *os.File, request *Request) error {
	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return InvalidJSONError.New("stdin is not a valid JSON")
	}
	return nil
}

func RespondToStdout(stdout *os.File, response interface{}) error {
	err := json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return OutputError.New("unable to output JSON")
	}
	return nil
}
