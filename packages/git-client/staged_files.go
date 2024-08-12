package gitclient

import (
	"strings"
)

func GetStagedFiles(dirPath *string) ([]string, error) {
	out, err := execCommandAtPath(dirPath, "git", "diff", "--cached", "--name-only", "--diff-filter=d")

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}

func GetNotStagedFiles(dirPath *string) ([]string, error) {
	out, err := execCommandAtPath(dirPath, "git", "diff", "--name-only", "--diff-filter=d")

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}
