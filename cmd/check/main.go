package main

import (
	"fmt"
	"io"
	"os"
)

func checkCommand(stdin io.Reader, stdout io.Writer) error {
	fmt.Fprintf(stdout, "[]")
	return nil
}

func main() {
	err := checkCommand(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
