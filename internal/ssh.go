package internal

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	easyssh "github.com/appleboy/easyssh-proxy"
)

const defaultTimeout = 60 * 10 // = 10 minutes

// PerformSSHCommand runs command on target machine via SSH.
// It redirects content from stdout and stderr of comamnd on target machine and
// returns error if anything goes wrong
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

	stdoutChan, stderrChan, doneChan, errChan, err := config.Stream(params.Script, defaultTimeout)
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
			fmt.Fprintf(stdout, "%s\n", outline)
		case errline := <-stderrChan:
			fmt.Fprintf(stderr, "%s\n", errline)
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
