package cli

import (
	"path"
	"strings"

	"github.com/LMaxence/gookme/packages/configuration"
	gitclient "github.com/LMaxence/gookme/packages/git-client"
	hooksscripts "github.com/LMaxence/gookme/packages/hooks-scripts"
	"github.com/LMaxence/gookme/packages/meta"
	"github.com/urfave/cli/v2"
)

const (
	InitCommandName CommandName = "init"
)

type InitCommandArguments struct {
	HookTypes []configuration.HookType
}

func parseInitCommandArguments(cContext *cli.Context) (*InitCommandArguments, error) {
	rawHookTypes := cContext.StringSlice("types")
	allHookTypes := cContext.Bool("all")

	hookTypes := make([]configuration.HookType, 0)

	if allHookTypes {
		hookTypes = configuration.ALL_HOOKS
	} else {
		for _, rawHookType := range rawHookTypes {
			hookType, err := validateHookType(rawHookType)
			if err != nil {
				return nil, err
			}
			hookTypes = append(hookTypes, hookType)
		}
	}

	args := &InitCommandArguments{
		HookTypes: hookTypes,
	}
	return args, nil
}

func createHookScriptFileOrAddScript(
	hookType string,
) error {
	repositoryPath, err := gitclient.GetRepoPath(nil)
	gitDirPath := path.Join(repositoryPath, ".git")

	if err != nil {
		return err
	}

	logger.Infof("Creating or updating %s hook script", hookType)

	var content string
	logger.Debugf("Checking if script %s file exists at path %s", string(hookType), gitDirPath)
	exists, err := hooksscripts.ScriptFileExists(gitDirPath, configuration.HookType(hookType))
	if err != nil {
		return err
	}

	if exists {
		logger.Debugf("Script file already exists for %s hook, loading content", hookType)
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
		logger.Debugf("Script file does not exist for %s hook", hookType)
		content = "#!/bin/sh\n\n"
		logger.Trace("Content of new script file:")
		logger.Trace(content)
	}

	scriptVariables := hooksscripts.ScriptVariables{
		HookType: configuration.HookType(hookType),
		Version:  meta.GOOKME_CLI_VERSION,
	}

	logger.Debugf("Adding gookme script to %s hook script", hookType)
	content, err = hooksscripts.AddGookmeScript(content, &scriptVariables)

	if err != nil {
		return err
	}

	logger.Tracef("Content of %s hook script after addition:", hookType)
	logger.Trace(content)

	logger.Debugf("Writing script to %s hook script file", hookType)
	err = hooksscripts.WriteScriptFileContent(gitDirPath, configuration.HookType(hookType), content)
	if err != nil {
		return err
	}

	logger.Infof("Successfully created or updated %s hook script", hookType)
	return nil
}

func assertHooksDir() error {
	repositoryPath, err := gitclient.GetRepoPath(nil)
	if err != nil {
		return err
	}

	hooksDirPath := path.Join(repositoryPath, "hooks")
	partialsDirPath := path.Join(hooksDirPath, "partials")
	partialGitKeepPath := path.Join(partialsDirPath, ".gitkeep")
	sharedDirPath := path.Join(hooksDirPath, "shared")
	sharedGitKeepPath := path.Join(sharedDirPath, ".gitkeep")

	logger.Debugf("Checking if hooks directory exists at path %s", partialsDirPath)

	err = hooksscripts.AssertFolder(hooksDirPath)
	if err != nil {
		return err
	}

	err = hooksscripts.AssertFolder(partialsDirPath)
	if err != nil {
		return err
	}

	err = hooksscripts.AssertFile(partialGitKeepPath)
	if err != nil {
		return err
	}

	err = hooksscripts.AssertFolder(sharedDirPath)
	if err != nil {
		return err
	}

	err = hooksscripts.AssertFile(sharedGitKeepPath)
	if err != nil {
		return err
	}

	return nil
}

func initFn(args InitCommandArguments) error {
	if len(args.HookTypes) == 0 {
		logger.Infof("No hook types provided, nothing to do")
		return nil
	}

	hookTypesString := make([]string, len(args.HookTypes))
	for i, hookType := range args.HookTypes {
		hookTypesString[i] = string(hookType)
	}

	logger.Infof("Initializing hooks: %v", strings.Join(hookTypesString, ", "))

	for _, hookType := range args.HookTypes {
		logger.Infof("Initializing %s hook", hookType)
		err := createHookScriptFileOrAddScript(string(hookType))
		if err != nil {
			return err
		}
		logger.Infof("Successfully initialized %s hook", hookType)
	}

	err := assertHooksDir()
	if err != nil {
		return err
	}

	return nil
}

var InitCommand *cli.Command = &cli.Command{
	Name:    string(InitCommandName),
	Aliases: []string{"i"},
	Usage:   "initialize Git hooks script files to run Gookme",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "types",
			Aliases: []string{"t"},
			Usage:   "The types of Git hooks to hook. Accepted values are: pre-commit, prepare-commit-msg, commit-msg,  post-commit, post-merge, post-rewrite, pre-rebase, post-checkout, pre-push",
		},
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Initialize all available hooks. Has precedence over the --types flag",
		},
	},
	Action: func(cContext *cli.Context) error {
		args, err := parseInitCommandArguments(cContext)
		if err != nil {
			return err
		}
		return initFn(*args)
	},
}
