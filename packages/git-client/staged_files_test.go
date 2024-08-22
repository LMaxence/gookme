package gitclient

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	helpers "github.com/LMaxence/gookme/packages/test-helpers"

	"github.com/stretchr/testify/assert"
)

func TestGetStagedFiles(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Create a file
	_, err = execCommandAtPath(&tmpDir, "touch", "file1")
	assert.NoError(t, err)

	// Stage the file
	_, err = execCommandAtPath(&tmpDir, "git", "add", ".")
	assert.NoError(t, err)

	// Call the function
	files, err := GetStagedFiles(&tmpDir, nil)

	// Assert the results
	assert.NoError(t, err)

	assert.Contains(t, files, filepath.Join(tmpDir, "file1"))
}

func TestGetStagedFilesWithNoStagedFiles(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Create a file
	_, err = execCommandAtPath(&tmpDir, "touch", "file1")
	assert.NoError(t, err)

	// Call the function
	files, err := GetStagedFiles(&tmpDir, nil)

	// Assert the results
	assert.NoError(t, err)
	assert.NotContains(t, files, filepath.Join(tmpDir, "file1"))
}

func TestGetNotStagedFiles(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Create a file
	_, err = execCommandAtPath(&tmpDir, "touch", "file1")
	assert.NoError(t, err)

	// Stage the file and then modify it
	_, err = execCommandAtPath(&tmpDir, "git", "add", ".")
	assert.NoError(t, err)

	// Write "test" to the file
	file, err := os.OpenFile(tmpDir+"/file1", os.O_WRONLY, fs.ModePerm)
	assert.NoError(t, err)
	defer file.Close()

	_, err = file.WriteString("test")
	assert.NoError(t, err)

	// Call the function
	files, err := GetNotStagedFiles(&tmpDir, nil)

	// Assert the results
	assert.NoError(t, err)
	assert.Contains(t, files, filepath.Join(tmpDir, "file1"))
}

func TestGetNotStagedFilesWithNoNotStagedFiles(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Create a file
	_, err = execCommandAtPath(&tmpDir, "touch", "file1")
	assert.NoError(t, err)

	// Stage the file
	_, err = execCommandAtPath(&tmpDir, "git", "add", ".")
	assert.NoError(t, err)

	// Call the function
	files, err := GetNotStagedFiles(&tmpDir, nil)

	// Assert the results
	assert.NoError(t, err)
	assert.NotContains(t, files, filepath.Join(tmpDir, "file1"))
}
