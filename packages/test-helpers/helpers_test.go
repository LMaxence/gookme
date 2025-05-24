package testhelpers

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupTmpGit(t *testing.T) {
	directory, err := SetupTmpGit()
	assert.NoError(t, err)

	// Check if the directory exists
	_, err = os.Stat(directory)
	assert.NoError(t, err)

	// Check if the directory is a git repository
	cmd := exec.Command("git", "status")
	cmd.Dir = directory
	err = cmd.Run()
	assert.NoError(t, err)
}

func TestWriteFile(t *testing.T) {
	directory, err := SetupTmpGit()
	assert.NoError(t, err)

	filename := "test.txt"
	content := "Hello, World!"

	err = WriteFile(directory, filename, content)
	assert.NoError(t, err)

	// Check if the file exists
	path := directory + "/" + filename
	_, err = os.Stat(path)
	assert.NoError(t, err)

	// Check the content of the file
	file, err := os.Open(path)
	assert.NoError(t, err)

	stat, err := file.Stat()
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), stat.Size())

	readContent := make([]byte, len(content))
	_, err = file.Read(readContent)
	assert.NoError(t, err)
	assert.Equal(t, content, string(readContent))

	err = file.Close()
	assert.NoError(t, err)
}
