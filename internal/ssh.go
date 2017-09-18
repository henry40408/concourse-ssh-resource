package internal

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	easyssh "github.com/appleboy/easyssh-proxy"
)

const DefaultTimeout = 60 * 10 // = 10 minutes

func PerformSSHCommand(source *Source, params *Params, stdout, stderr io.Writer) error {
	config := &easyssh.MakeConfig{
		Server:   source.Host,
		Port:     "22",
		User:     source.User,
		Password: source.Password,
		Key:      source.PrivateKey,
	}

	if source.Port != 0 {
		config.Port = strconv.Itoa(source.Port)
	}

	stdoutChan, stderrChan, doneChan, errChan, err := config.Stream(params.Script, DefaultTimeout)
	if err != nil {
		return fmt.Errorf("failed to run command: %s", err.Error())
	}

	done := true

loop:
	for {
		select {
		case done = <-doneChan:
			break loop
		case outline := <-stdoutChan:
			fmt.Fprintf(stdout, outline)
		case errline := <-stderrChan:
			fmt.Fprintf(stderr, errline)
		case err = <-errChan:
		}
	}

	if err != nil {
		return fmt.Errorf("failed when running SSH command: %s", err.Error())
	}

	if !done {
		return errors.New("SSH command times out")
	}

	return nil
}
