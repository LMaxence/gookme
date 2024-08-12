package gitclient

import (
	"os"
	"strings"
)

// GetRepoPath returns the absolute path of the .git repository
// for the current working directory.
func GetRepoPath(start *string) (string, error) {
	// If the start path is not provided, use the current working directory
	if start == nil {
		path, err := os.Getwd()
		if err != nil {
			return "", err
		}

		start = &path
	}

	// Run the git-rev-parse command from the start path
	out, err := execCommandAtPath(start, "git", "rev-parse", "--show-toplevel")

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
