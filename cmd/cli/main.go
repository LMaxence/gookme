package main

import (
	"os"
	"sync"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/LMaxence/gookme/packages/executor"
	gitclient "github.com/LMaxence/gookme/packages/git-client"
	"github.com/LMaxence/gookme/packages/logging"
)

var logger = logging.NewLogger("cli")

func main() {
	dir, err := gitclient.GetRepoPath(nil)

	if err != nil {
		logger.Errorf("Error while getting current working directory: %s", err)
		os.Exit(1)
		return
	}

	logger.Infof("Running gookme in %s", dir)
	logger.Info("Loading configurations")
	conf, err := configuration.LoadConfiguration(dir, configuration.PreCommit)

	if err != nil {
		logger.Errorf("Error while loading configuration: %s", err)
		os.Exit(1)
		return
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
}
