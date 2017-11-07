package main

import (
	"fmt"
	"io"
)

func inCommand(stdin io.Reader, stdout io.Writer) error {
	fmt.Fprintf(stdout, `{"version":{},"metadata":[]}`)
	return nil
}
