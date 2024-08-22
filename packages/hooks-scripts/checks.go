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

func AssertFolder(path string) error {
	// Check if a folder exists, if not create it
	logger.Debugf("Checking if hooks folder exists at path %s", path)

	_, err := os.Stat(path)
	if err == nil {
		logger.Debugf("Folder %s exists", path)
		return nil
	}

	if os.IsNotExist(err) {
		logger.Debugf("Hooks folder does not exist, creating it")
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func AssertFile(path string) error {
	// Check if a file exists, if not create it
	logger.Debugf("Checking if file exists at path %s", path)

	_, err := os.Stat(path)
	if err == nil {
		logger.Debugf("File %s exists", path)
		return nil
	}

	if os.IsNotExist(err) {
		logger.Debugf("File does not exist, creating it")
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}
