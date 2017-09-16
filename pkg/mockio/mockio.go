package mockio

import (
	"io/ioutil"
	"os"
)

type MockIO struct {
	In  *os.File
	Out *os.File
	Err *os.File
}

func NewMockIO(content []byte) (*MockIO, error) {
	in, err := ioutil.TempFile(os.TempDir(), "stdin")
	if err != nil {
		return nil, err
	}

	_, err = in.Write(content)
	if err != nil {
		return nil, err
	}
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

func (m *MockIO) Cleanup() {
	os.Remove(m.In.Name())
	os.Remove(m.Out.Name())
	os.Remove(m.Err.Name())
}
