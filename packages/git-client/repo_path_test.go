package gitclient

import (
	"os/exec"
	"strings"
	"testing"

	helpers "github.com/LMaxence/gookme/packages/test-helpers"

	"github.com/stretchr/testify/assert"
)

func TestGetGookmeRepoPath(t *testing.T) {
	path, err := GetRepoPath(nil)

	// Assert the results
	assert.NoError(t, err)
	assert.Contains(t, path, "/gookme")
}

func TestGetRepoPathWithStart(t *testing.T) {
	temporaryDirectory, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	path, err := GetRepoPath(&temporaryDirectory)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, temporaryDirectory, path)
}

func TestGetRepoPathExitCode1(t *testing.T) {
	// Create a temporary directory that is not a git repository
	out, err := exec.Command("mktemp", "-d").Output()
	assert.NoError(t, err)
	temporaryDirectory := strings.TrimSpace(string(out))

	// Call the function
	path, err := GetRepoPath(&temporaryDirectory)

	// Assert the results
	assert.Error(t, err)
	assert.Equal(t, "", path)
}

func TestGetRepoPathInvalidPath(t *testing.T) {
	nonExistingPath := "/path/does/not/exist"

	// Call the function
	path, err := GetRepoPath(&nonExistingPath)

	// Assert the results
	assert.Error(t, err)
	assert.Equal(t, "", path)
}
