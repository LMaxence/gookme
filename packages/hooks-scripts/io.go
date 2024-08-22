package hooksscripts

import (
	"os"
	"path"

	"github.com/LMaxence/gookme/packages/configuration"
)

func LoadScriptFileContent(gitFolderPath string, hookName configuration.HookType) (string, error) {
	// Check if the script file exists in the .git folder of the provided directory
	hookPath := path.Join(gitFolderPath, "hooks", string(hookName))

	content, err := os.ReadFile(hookPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

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
