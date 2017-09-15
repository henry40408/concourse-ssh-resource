package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/henry40408/ssh-shell-resource/internal/models"
	"github.com/spacemonkeygo/errors"
)

var (
	CheckError       = errors.NewClass("CheckError")
	InvalidJSONError = errors.NewClass("InvalidJSONError")
	OutputError      = errors.NewClass("OutputError")
)

func Main(stdin, stdout *os.File) error {
	var request models.CheckRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return InvalidJSONError.New("stdin is not a valid JSON")
	}

	response := CheckCommand(&request)

	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return OutputError.New("unable to output JSON")
	}

	return nil
}

func main() {
	err := Main(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, errors.GetMessage(err))
		os.Exit(1)
	}
}
