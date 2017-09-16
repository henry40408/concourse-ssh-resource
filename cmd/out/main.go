package main

import (
	"fmt"
	"os"

	"github.com/henry40408/ssh-shell-resource/internal"
	"github.com/spacemonkeygo/errors"
)

func Main(stdin, stdout, stderr *os.File) error {
	var request internal.Request

	err := internal.NewRequestFromStdin(stdin, &request)
	if err != nil {
		return err
	}

	outWriter := internal.PrefixWriter{
		Writer: stderr,
		Prefix: "stdout",
	}

	errWriter := internal.PrefixWriter{
		Writer: stderr,
		Prefix: "stderr",
	}

	err = internal.PerformSSHCommand(&request, &outWriter, &errWriter)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Main(os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, errors.GetMessage(err))
		os.Exit(1)
	}
}
