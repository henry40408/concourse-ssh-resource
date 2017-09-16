package internal

import "github.com/spacemonkeygo/errors"

var (
	ArgumentError = errors.NewClass("ArgumentError")
	FileError     = errors.NewClass("FileError")
	SSHError      = errors.NewClass("SSHError")
	TimeoutError  = errors.NewClass("TimeoutError")
)
