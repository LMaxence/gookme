package cli

import (
	"sync"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/LMaxence/gookme/packages/executor"
	gitclient "github.com/LMaxence/gookme/packages/git-client"
	"github.com/urfave/cli/v2"
)

const (
	RunCommandName CommandName = "run"
)

type RunCommandArguments struct {
	HookType       configuration.HookType
	GitCommandArgs []string
}

func parseRunCommandArguments(cContext *cli.Context) (*RunCommandArguments, error) {
	hookType, err := validateHookType(cContext.String("type"))
	if err != nil {
		return nil, err
	}
	args := &RunCommandArguments{
		HookType:       hookType,
		GitCommandArgs: cContext.Args().Slice(),
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
	},
	Action: func(cContext *cli.Context) error {
		args, err := parseRunCommandArguments(cContext)

		if err != nil {
			return err
		}
		return run(*args)
	},
}
