package testhelpers

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func SetupTmpGit() (string, error) {
	// Create a temporary directory
	directory, err := exec.Command("mktemp", "-d").Output()
	if err != nil {
		return "", err
	}

	path := strings.TrimSpace(string(directory))
	path = "/private" + path

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
