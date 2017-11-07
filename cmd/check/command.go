package main

import (
	"fmt"
	"io"
)

func checkCommand(stdin io.Reader, stdout io.Writer) error {
	fmt.Fprintf(stdout, "[]")
	return nil
}
