package hooksscripts

/* ====================================================================

This file contains utility functions to read and write scripts in the .git/hooks
folder of a git repository.

===================================================================== */

import (
	"os"
	"path"

	"github.com/LMaxence/gookme/packages/configuration"
)

// LoadScriptFileContent loads the content of a file located in the .git folder of the
// provided directory, and named after the provided hook name.
func LoadScriptFileContent(gitFolderPath string, hookName configuration.HookType) (string, error) {
	// Check if the script file exists in the .git folder of the provided directory
	hookPath := path.Join(gitFolderPath, "hooks", string(hookName))

	content, err := os.ReadFile(hookPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WriteScriptFileContent writes a script in the .git folder, with the provided content
// and named after the provided hook name.
func WriteScriptFileContent(gitFolderPath string, hookName configuration.HookType, content string) error {
	// Check if the script file exists in the folder of the provided directory
	hookPath := path.Join(gitFolderPath, "hooks", string(hookName))

	err := os.WriteFile(hookPath, []byte(content), 0755)
	if err != nil {
		return err
	}

	return nil
}

func DeleteScriptFile(gitFolderPath string, hookName configuration.HookType) error {
	// Check if the script file exists in the folder of the provided directory
	hookPath := path.Join(gitFolderPath, "hooks", string(hookName))

	err := os.Remove(hookPath)
	if err != nil {
		return err
	}

	return nil
}
