package main

import (
	"fmt"
	"io"
	"os"
)

func inCommand(stdin io.Reader, stdout io.Writer) error {
	fmt.Fprintf(stdout, `{"version":{},"metadata":[]}`)
	return nil
}

func main() {
	err := inCommand(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
