package mockio

import (
	"io"
	"os"

	"github.com/spf13/afero"
)

// MockIO holds three file to imitate standard input, standard output,
// and standard error
type MockIO struct {
	In  afero.File
	Out afero.File
	Err afero.File
}

// NewMockIO returns new MockIO object. Content in `reader` would be read into
// standard input so caller can read from it like ordinary standard input
func NewMockIO(reader io.Reader) (*MockIO, error) {
	fs := afero.NewMemMapFs()

	in, err := fs.Create("stdin")
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 8)
	for {
		_, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		in.Write(buf)
	}
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
