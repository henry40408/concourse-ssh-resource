package internal

import (
	"fmt"
	"io"
)

type PrefixWriter struct {
	Writer io.Writer
	Prefix string
}

func (pw *PrefixWriter) Write(p []byte) (n int, err error) {
	return fmt.Fprintf(pw.Writer, "%s: %s", pw.Prefix, p)
}
