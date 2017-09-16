package internal

import (
	"fmt"
	"os"
	"strconv"

	easyssh "github.com/appleboy/easyssh-proxy"
)

func PerformSSHCommand(request *Request, stdout, stderr *os.File) error {
	source := request.Source
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

	params := request.Params
	stdoutChan, stderrChan, doneChan, errChan, err := config.Stream(params.Script, SSHTimeout)
	if err != nil {
		return SSHError.New("failed to run command: %s", err.Error())
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
		return SSHError.New("failed when running SSH command: %s", err.Error())
	}

	if !done {
		return TimeoutError.New("SSH command times out")
	}

	return nil
}
