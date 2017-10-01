package mockio

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/afero"
)

const (
	// IN represents imitated stdin
	IN = 1 << iota
	// OUT represents imitated stdout
	OUT = 1 << iota
	// ERR represents imitated stderr
	ERR = 1 << iota
)

// MockIO holds three file to imitate stdin, stdout, and stderr
type MockIO struct {
	In  afero.File
	Out afero.File
	Err afero.File
}

// NewMockIO returns new MockIO object. `content` would be write into stdin
// so caller can read from it like ordinary stdin
func NewMockIO(content []byte) (*MockIO, error) {
	fs := afero.NewMemMapFs()

	in, err := fs.Create("stdin")
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

	out, err := fs.Create("stdout")
	if err != nil {
		return nil, err
	}

	stdErr, err := fs.Create("stderr")
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

func readAll(file afero.File) ([]byte, error) {
	file.Seek(0, 0)

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
