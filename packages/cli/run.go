package cli

import (
	"sync"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/LMaxence/gookme/packages/executor"
	gitclient "github.com/LMaxence/gookme/packages/git-client"
)

type RunCommandArguments struct {
	HookType configuration.HookType
}

func Run(args RunCommandArguments) error {
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
		exec := executor.NewHookExecutor(&hook)
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
