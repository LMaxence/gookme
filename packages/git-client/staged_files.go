package gitclient

import (
	"fmt"
	"strings"

	"github.com/LMaxence/gookme/packages/logging"
)

var logger = logging.NewLogger("git-client")

type GitRefDelimiter struct {
	From string
	To   string
}

func GetStagedFiles(dirPath *string, delimiter *GitRefDelimiter) ([]string, error) {
	root, err := GetRepoPath(dirPath)
	if err != nil {
		return nil, err
	}

	var out string
	if delimiter != nil {
		out, err = execCommandAtPath(
			dirPath,
			"git",
			"diff",
			"--name-only",
			"--diff-filter=d",
			fmt.Sprintf("--line-prefix=%s", root+"/"),
			fmt.Sprintf("%s...%s", delimiter.From, delimiter.To),
		)
	} else {
		out, err = execCommandAtPath(
			dirPath,
			"git",
			"diff",
			"--cached",
			"--name-only",
			"--diff-filter=d",
			fmt.Sprintf("--line-prefix=%s", root+"/"),
		)
	}

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}

func GetNotStagedFiles(dirPath *string, delimiter *GitRefDelimiter) ([]string, error) {
	root, err := GetRepoPath(dirPath)
	if err != nil {
		return nil, err
	}

	var from string
	var to string
	if delimiter == nil {
		from = ""
		to = ""
	} else {
		from = delimiter.From
		to = delimiter.To
	}

	out, err := execCommandAtPath(
		dirPath,
		"git",
		"diff",
		"--name-only",
		"--diff-filter=d",
		fmt.Sprintf("--line-prefix=%s", root+"/"),
		from,
		to,
	)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}
