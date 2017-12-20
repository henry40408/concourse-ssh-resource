//+build !test

package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()
	err := outCommand(fs, os.Args, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
