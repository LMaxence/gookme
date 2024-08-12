package gitclient

import (
	"os/exec"
	"strings"
)

func setupTmpGit() (string, error) {
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
