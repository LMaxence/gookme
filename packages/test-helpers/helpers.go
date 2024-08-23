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

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

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
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
