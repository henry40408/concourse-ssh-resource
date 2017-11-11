package mockio

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

func TestMockIO(t *testing.T) {
	randomString := fake.WordsN(3)
	reader := bytes.NewBufferString(randomString)
	io, err := NewMockIO(reader)

	inBytes, err := ioutil.ReadAll(io.In)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, []byte(randomString), inBytes)
}
