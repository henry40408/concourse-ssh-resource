package ssh

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/appleboy/easyssh-proxy"
	hierr "github.com/reconquest/hierr-go"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
)

const defaultTimeout = 60 * 10 // = 10 minutes

// PerformSSHCommand runs command on remote machine via SSH.
// It puts script into file on remote machine, and runs it with interpreter.
func PerformSSHCommand(source *models.Source, params *models.Params, stdout, stderr io.Writer) error {
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

	interpreter := params.Interpreter
	if interpreter == "" {
		interpreter = "/bin/sh"
	}

	remoteScriptFileName, err := putScriptInLocalFile(config, params.Script)
	if err != nil {
		return err
	}

	command := fmt.Sprintf("%s %s", interpreter, remoteScriptFileName)
	stdoutChan, stderrChan, doneChan, errChan, err := config.Stream(command, defaultTimeout)
	if err != nil {
		return hierr.Errorf(err, "unable to run script on remote machine")
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
		return hierr.Errorf(err, "SSH command failed on remote machine")
	}

	if !done {
		return errors.New("SSH command timed out")
	}

	return nil
}

func putScriptInLocalFile(config *easyssh.MakeConfig, script string) (string, error) {
	localScriptFile, err := ioutil.TempFile(os.TempDir(), "script")
	defer localScriptFile.Close()
	if err != nil {
		return "", hierr.Errorf(err, "unable to create temporary file on local machine")
	}

	localScriptFile.WriteString(script)

	remoteScriptFileName := fmt.Sprintf("/tmp/script-%d", time.Now().Unix())

	err = config.Scp(localScriptFile.Name(), remoteScriptFileName)
	if err != nil {
		return "", hierr.Errorf(err, "unable to copy script to remote machine")
	}

	return remoteScriptFileName, nil
}
