package cli

import (
	"sync"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/LMaxence/gookme/packages/executor"
	gitclient "github.com/LMaxence/gookme/packages/git-client"
	"github.com/LMaxence/gookme/packages/scheduler"
	"github.com/urfave/cli/v2"
)

const (
	RunCommandName CommandName = "run"
)

type RunCommandArguments struct {
	HookType       configuration.HookType
	GitCommandArgs []string
	From           string
	To             string
}

func parseRunCommandArguments(cContext *cli.Context) (*RunCommandArguments, error) {
	hookType, err := validateHookType(cContext.String("type"))
	if err != nil {
		return nil, err
	}

	args := &RunCommandArguments{
		HookType:       hookType,
		GitCommandArgs: cContext.Args().Slice(),
		From:           cContext.String("from"),
		To:             cContext.String("to"),
	}
	return args, nil
}

func run(args RunCommandArguments) error {
	dir, err := gitclient.GetRepoPath(nil)
	if err != nil {
		logger.Errorf("Error while getting current working directory: %s", err)
		return err
	}

	logger.Info("Loading configurations")
	conf, err := configuration.LoadConfiguration(dir, args.HookType)
	if err != nil {
		logger.Errorf("Error while loading configuration: %s", err)
		return err
	}

	var delimiter *gitclient.GitRefDelimiter
	if args.From != "" && args.To != "" {
		logger.Debugf("Setting Git ref delimiter from %s to %s", args.From, args.To)
		delimiter = &gitclient.GitRefDelimiter{
			From: args.From,
			To:   args.To,
		}
	}

	changedPaths, err := gitclient.GetStagedFiles(&dir, delimiter)
	logger.Debugf("Staged files: %v", changedPaths)
	if err != nil {
		logger.Errorf("Error while getting staged files: %s", err)
		return err
	}

	conf.Hooks = scheduler.FilterHooksWithChangeset(changedPaths, conf.Hooks)

	nSteps := 0
	for _, hook := range conf.Hooks {
		nSteps += len(hook.Steps)
	}

	logger.Infof("Running %d hooks, %d steps", len(conf.Hooks), nSteps)
	executors := make([]*executor.HookExecutor, 0, len(conf.Hooks))

	for _, hook := range conf.Hooks {
		exec := executor.NewHookExecutor(&hook, args.GitCommandArgs)
		exec = exec.WithExitOnStepError()
		executors = append(executors, exec)
	}

	hooksWg := sync.WaitGroup{}
	for _, exec := range executors {
		hooksWg.Add(1)
		go func() {
			exec.Run()
			hooksWg.Done()
		}()
	}

	hooksWg.Wait()
	return nil
}

var RunCommand *cli.Command = &cli.Command{
	Name:    string(RunCommandName),
	Aliases: []string{"r"},
	Usage:   "load and run git hooks based on staged Git changes",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Value:   "pre-commit",
			Usage:   "The type of Git hook to run. Default is pre-commit, but accepted values are: pre-commit, prepare-commit-msg, commit-msg,  post-commit, post-merge, post-rewrite, pre-rebase, post-checkout, pre-push",
		},
		&cli.StringFlag{
			Name:    "from",
			Aliases: []string{"f"},
			Usage:   "An optional commit SHA-1 hash to compare to generate the staged changes from.",
		},
		&cli.StringFlag{
			Name:    "to",
			Aliases: []string{"o"},
			Usage:   "An optional commit SHA-1 hash to compare to generate the staged changes to.",
		},
	},
	Action: func(cContext *cli.Context) error {
		args, err := parseRunCommandArguments(cContext)

		if err != nil {
			return err
		}
		return run(*args)
	},
}
