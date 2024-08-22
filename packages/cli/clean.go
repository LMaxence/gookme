package cli

import (
	"path"
	"strings"

	"github.com/LMaxence/gookme/packages/configuration"
	gitclient "github.com/LMaxence/gookme/packages/git-client"
	hooksscripts "github.com/LMaxence/gookme/packages/hooks-scripts"
	"github.com/urfave/cli/v2"
)

const (
	CleanCommandName CommandName = "clean"
)

func cleanHookScriptFile(
	hookType string,
) error {
	repositoryPath, err := gitclient.GetRepoPath(nil)
	gitDirPath := path.Join(repositoryPath, ".git")

	if err != nil {
		return err
	}

	logger.Infof("Cleaning %s hook script", hookType)

	var content string
	logger.Debugf("Checking if script %s file exists at path %s", string(hookType), gitDirPath)
	exists, err := hooksscripts.ScriptFileExists(gitDirPath, configuration.HookType(hookType))
	if err != nil {
		return err
	}

	if exists {
		logger.Debugf("Script file exists for %s hook, loading content", hookType)
		content, err = hooksscripts.LoadScriptFileContent(gitDirPath, configuration.HookType(hookType))
		logger.Tracef("Loaded content of %s hook script:", hookType)
		logger.Trace(content)

		if err != nil {
			return err
		}

		// Remove the existing gookme script from the content if it exists
		content = hooksscripts.RemoveExistingGookmeScript(content)
		logger.Tracef("Content of %s hook script after deletion:", hookType)
		logger.Trace(content)
	} else {
		logger.Infof("Script file %s does not exist", hookType)
	}

	logger.Debugf("Writing script to %s hook script file", hookType)

	if strings.Trim(content, " \n") == "#!/bin/sh" {
		logger.Infof("Script file %s is empty, removing it", hookType)
		err = hooksscripts.DeleteScriptFile(gitDirPath, configuration.HookType(hookType))
		if err != nil {
			return err
		}
	} else {
		err = hooksscripts.WriteScriptFileContent(gitDirPath, configuration.HookType(hookType), content)

		if err != nil {
			return err
		}
	}

	logger.Infof("Successfully created or updated %s hook script", hookType)
	return nil
}

func clean() error {

	logger.Info("Cleaning hooks")

	for _, hookType := range configuration.ALL_HOOKS {
		logger.Infof("Cleaning %s hook", hookType)
		err := cleanHookScriptFile(string(hookType))
		if err != nil {
			return err
		}
		logger.Infof("Successfully cleaned %s hook", hookType)
	}

	return nil
}

var CleanCommand *cli.Command = &cli.Command{
	Name:    string(CleanCommandName),
	Aliases: []string{"c"},
	Usage:   "Clean Git hooks scripts configured by Gookme",
	Action: func(cContext *cli.Context) error {
		return clean()
	},
}
