package gitclient

import (
	"os/exec"
)

func execCommandAtPath(path *string, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	if path != nil {
		cmd.Dir = *path
	}

	logger.Debugf("Executing command %s in %s", cmd.String(), cmd.Dir)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return string(out), err
	}

	if len(out) == 0 {
		return "", nil
	}

	return string(out), nil
}
