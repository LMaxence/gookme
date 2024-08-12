package gitclient

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStagedFiles(t *testing.T) {
	// Create a temporary directory
	temporaryDirectory, err := setupTmpGit()
	assert.NoError(t, err)

	// Create a file
	_, err = execCommandAtPath(&temporaryDirectory, "touch", "file1")
	assert.NoError(t, err)

	// Stage the file
	_, err = execCommandAtPath(&temporaryDirectory, "git", "add", ".")
	assert.NoError(t, err)

	// Call the function
	files, err := GetStagedFiles(&temporaryDirectory)

	// Assert the results
	assert.NoError(t, err)
	assert.Contains(t, files, "file1")
}

func TestGetStagedFilesWithNoStagedFiles(t *testing.T) {
	// Create a temporary directory
	temporaryDirectory, err := setupTmpGit()
	assert.NoError(t, err)

	// Create a file
	_, err = execCommandAtPath(&temporaryDirectory, "touch", "file1")
	assert.NoError(t, err)

	// Call the function
	files, err := GetStagedFiles(&temporaryDirectory)

	// Assert the results
	assert.NoError(t, err)
	assert.NotContains(t, files, "file1")
}

func TestGetNotStagedFiles(t *testing.T) {
	// Create a temporary directory
	temporaryDirectory, err := setupTmpGit()
	assert.NoError(t, err)

	// Create a file
	_, err = execCommandAtPath(&temporaryDirectory, "touch", "file1")
	assert.NoError(t, err)

	// Stage the file and then modify it
	_, err = execCommandAtPath(&temporaryDirectory, "git", "add", ".")
	assert.NoError(t, err)

	// Write "test" to the file
	file, err := os.OpenFile(temporaryDirectory+"/file1", os.O_WRONLY, fs.ModePerm)
	assert.NoError(t, err)
	defer file.Close()

	_, err = file.WriteString("test")
	assert.NoError(t, err)

	// Call the function
	files, err := GetNotStagedFiles(&temporaryDirectory)

	// Assert the results
	assert.NoError(t, err)
	assert.Contains(t, files, "file1")
}

func TestGetNotStagedFilesWithNoNotStagedFiles(t *testing.T) {
	// Create a temporary directory
	temporaryDirectory, err := setupTmpGit()
	assert.NoError(t, err)

	// Create a file
	_, err = execCommandAtPath(&temporaryDirectory, "touch", "file1")
	assert.NoError(t, err)

	// Stage the file
	_, err = execCommandAtPath(&temporaryDirectory, "git", "add", ".")
	assert.NoError(t, err)

	// Call the function
	files, err := GetNotStagedFiles(&temporaryDirectory)

	// Assert the results
	assert.NoError(t, err)
	assert.NotContains(t, files, "file1")
}
