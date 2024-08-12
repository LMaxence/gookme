package gitclient

import (
	"os/exec"
)

func execCommandAtPath(path *string, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	if path != nil {
		cmd.Dir = *path
	}

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	if len(out) == 0 {
		return "", nil
	}

	return string(out), nil
}
