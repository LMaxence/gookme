package gitclient

import (
	"fmt"
	"strings"
)

func GetStagedFiles(dirPath *string) ([]string, error) {
	root, err := GetRepoPath(dirPath)
	if err != nil {
		return nil, err
	}

	out, err := execCommandAtPath(
		dirPath,
		"git",
		"diff",
		"--cached",
		"--name-only",
		"--diff-filter=d",
		fmt.Sprintf("--line-prefix=%s", root+"/"),
	)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}

func GetNotStagedFiles(dirPath *string) ([]string, error) {
	root, err := GetRepoPath(dirPath)
	if err != nil {
		return nil, err
	}

	out, err := execCommandAtPath(
		dirPath,
		"git",
		"diff",
		"--name-only",
		"--diff-filter=d",
		fmt.Sprintf("--line-prefix=%s", root+"/"),
	)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}
