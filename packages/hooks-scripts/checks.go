package hooksscripts

import (
	"os"
	"path"

	"github.com/LMaxence/gookme/packages/configuration"
)

// Check if git hook script for the provided hook type exists in the provided directory
func ScriptFileExists(gitFolderPath string, hookName configuration.HookType) (bool, error) {
	// Check if the script file exists in the .git folder of the provided directory
	hookPath := path.Join(gitFolderPath, "hooks", string(hookName))
	_, err := os.Stat(hookPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
