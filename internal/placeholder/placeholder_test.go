package placeholder

import (
	"bytes"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/henry40408/concourse-ssh-resource/internal/models"
)

const (
	baseDir = "/tmp"
)

func TestReplacePlaceholdersWithValue(t *testing.T) {
	stderr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	params := &models.Params{
		Script: `echo "<placeholder>"`,
		Placeholders: []models.Placeholder{
			models.Placeholder{Name: "<placeholder>", Value: "foobar"},
		},
	}

	script, err := ReplacePlaceholders(stderr, fs, baseDir, params)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, `echo "foobar"`, script)
}

func TestReplacePlaceholdersWithFile(t *testing.T) {
	stderr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	fs.MkdirAll(baseDir, 0755)
	afero.WriteFile(fs, "/tmp/somefile", []byte("foobar"), 0644)

	params := &models.Params{
		Script: `echo "<placeholder>"`,
		Placeholders: []models.Placeholder{
			models.Placeholder{Name: "<placeholder>", File: "somefile"},
		},
	}

	script, err := ReplacePlaceholders(stderr, fs, baseDir, params)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, `echo "foobar"`, script)
}

func TestReplacePlaceholdersWithEmptyFile(t *testing.T) {
	stderr := bytes.NewBuffer([]byte{})

	fs := afero.NewMemMapFs()
	fs.MkdirAll(baseDir, 0755)
	afero.WriteFile(fs, "/tmp/somefile", []byte(""), 0644)

	params := &models.Params{
		Script: `echo "<placeholder>"`,
		Placeholders: []models.Placeholder{
			models.Placeholder{Name: "<placeholder>", File: "somefile"},
		},
	}

	_, err := ReplacePlaceholders(stderr, fs, baseDir, params)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "File for placeholder '<placeholder>' seems to be empty")
	}
}

func TestReplacePlaceholdersWithNonusedPlaceholder(t *testing.T) {
	stderr := bytes.NewBuffer([]byte{})
	fs := afero.NewMemMapFs()
	params := &models.Params{
		Script: `echo "<placeholder>"`,
		Placeholders: []models.Placeholder{
			models.Placeholder{Name: "<place>", Value: "foobar"},
		},
	}

	_, err := ReplacePlaceholders(stderr, fs, baseDir, params)
	if !assert.NoError(t, err) {
		return
	}

	assert.Contains(t, stderr.String(), "Placeholder '<place>' is not found in script")
}

func TestReplacePlaceholdersWithPlaceholderWhichHasNoValueNorFile(t *testing.T) {
	fs := afero.NewMemMapFs()

	params := &models.Params{
		Script: `echo "<placeholder>"`,
		Placeholders: []models.Placeholder{
			models.Placeholder{Name: "<placeholder>"},
		},
	}

	stderr := bytes.NewBuffer([]byte{})

	_, err := ReplacePlaceholders(stderr, fs, baseDir, params)
	if !assert.NoError(t, err) {
		return
	}

	assert.Contains(t, stderr.String(), "Neither File nor Value are set for placeholder '<placeholder>'")
}
