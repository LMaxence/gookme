package testhelpers

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/LMaxence/gookme/packages/logging"
)

var logger = logging.NewLogger("test-helpers")

// directoryExists checks if a directory exists and is a directory
func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// CreateTemporaryDirectory creates a temporary directory and returns its path
func CreateTemporaryDirectory() (string, error) {
	// Create a temporary directory
	directory, err := exec.Command("mktemp", "-d").Output()

	var path string

	if err != nil {
		return "", err
	}

	path = strings.TrimSpace(string(directory))
	path = "/private" + path

	exists := directoryExists(path)

	// This is a workaround for creating temporary directories within GitHub Actions
	if !exists {
		logger.Warn("mktemp command failed, falling back to os.MkdirTemp")

		tmpDir := os.Getenv("TMPDIR")
		if tmpDir == "" {
			return "", errors.New("TMPDIR environment variable is not set")
		}

		path, err = os.MkdirTemp(tmpDir, "gookme-tmp-*")
		if err != nil {
			return "", err
		}

		exists = directoryExists(path)
		if !exists {
			return "", errors.New("failed to create temporary directory")
		}
	}

	return path, nil
}

// SetupTmpGit creates a temporary directory and initializes a git repository in it
// It returns the path to the temporary directory to be later on used within tests
func SetupTmpGit() (string, error) {
	path, err := CreateTemporaryDirectory()
	if err != nil {
		return "", err
	}

	// Initialize a git repository
	gitInitCmd := exec.Command("git", "init")
	gitInitCmd.Dir = path
	_, err = gitInitCmd.Output()

	if err != nil {
		return "", err
	}

	return path, nil
}

// WriteFile writes a file with the provided content in the provided directory
func WriteFile(directory, filename, content string) error {
	// Create directory if it does not exist
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return err
	}

	path := filepath.Join(directory, filename)
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return file.Close()
}
