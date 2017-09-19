package mockio

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	IN  = 1
	OUT = 2
	ERR = 4
)

// MockIO holds three file to imitate stdin, stdout, and stderr
type MockIO struct {
	In  *os.File
	Out *os.File
	Err *os.File
}

// NewMockIO returns new MockIO object. `content` would be write into stdin
// so caller can read from it like ordinary stdin
func NewMockIO(content []byte) (*MockIO, error) {
	in, err := ioutil.TempFile(os.TempDir(), "stdin")
	if err != nil {
		return nil, err
	}

	// writes content into stdin
	_, err = in.Write(content)
	if err != nil {
		return nil, err
	}

	// resets cursor in stdin so caller can read it from the beginning
	in.Seek(0, 0)

	out, err := ioutil.TempFile(os.TempDir(), "stdout")
	if err != nil {
		return nil, err
	}

	stdErr, err := ioutil.TempFile(os.TempDir(), "stderr")
	if err != nil {
		return nil, err
	}

	return &MockIO{In: in, Out: out, Err: stdErr}, nil
}

// Cleanup removes all temporary files in MockIO
func (m *MockIO) Cleanup() {
	os.Remove(m.In.Name())
	os.Remove(m.Out.Name())
	os.Remove(m.Err.Name())
}

// ReadAll reads everything in imitated stderr file
func (m *MockIO) ReadAll(flag int) ([]byte, error) {
	switch flag {
	case IN:
		return readAll(m.In)
	case OUT:
		return readAll(m.Out)
	case ERR:
		return readAll(m.Err)
	default:
		return nil, fmt.Errorf("illegal flag: %d", flag)
	}

}

func readAll(file *os.File) ([]byte, error) {
	file.Seek(0, 0)

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
