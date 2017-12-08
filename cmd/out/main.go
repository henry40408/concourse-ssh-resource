//+build !test

package main

import (
	"fmt"
	"os"
	"log"
	"errors"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Accessing first argument", errors.New("usage: %s <base directory>\n"))
		os.Exit(1)
	}
	var baseDir string = os.Args[1]

	err := outCommand(os.Stdin, os.Stdout, os.Stderr, baseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
