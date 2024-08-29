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

func GetChangedFilesBetweenRefs(
	dirPath *string,
	from string,
	to string,
) ([]string, error) {
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
		fmt.Sprintf("%s...%s", from, to),
	)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}

func getCommitsToBePushed(dirPath *string) ([]string, error) {
	out, err := execCommandAtPath(
		dirPath,
		"git",
		"rev-list",
		"@{push}^..",
	)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}

func GetFilesToBePushed(dirPath *string) ([]string, error) {
	commits, err := getCommitsToBePushed(dirPath)
	if err != nil {
		return nil, err
	}

	if len(commits) == 0 {
		logger.Warn("Commits to push list is empty")
		return []string{}, nil
	}

	start, end := commits[0], commits[len(commits)-1]

	return GetChangedFilesBetweenRefs(dirPath, start, end)
}

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

func GetFilesChangedNCommitsBefore(dirPath *string, n int) ([]string, error) {
	out, err := execCommandAtPath(
		dirPath,
		"git",
		"diff-tree",
		"--no-commit-id",
		"--name-only", "-r",
		fmt.Sprintf("HEAD~%d", n),
	)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}
