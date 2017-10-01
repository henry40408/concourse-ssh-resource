package main

import (
	"fmt"
	"io"
)

type prefixWriter struct {
	prefix string
	writer io.Writer
}

func (pw *prefixWriter) Write(p []byte) (n int, err error) {
	return fmt.Fprintf(pw.writer, "%s: %s", pw.prefix, p)
}
