package executor

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/LMaxence/gookme/packages/logging"
)

var logger = logging.NewLogger("executor")

type StepExecutionStatus string

const (
	StepExecutionStatusPending StepExecutionStatus = "pending"
	StepExecutionStatusRunning StepExecutionStatus = "running"
	StepExecutionStatusSuccess StepExecutionStatus = "success"
	StepExecutionStatusFailed  StepExecutionStatus = "failed"
)

type HookExecutor struct {
	hook        *configuration.Hook
	tempDirPath string
	files       map[string]*os.File
	exitOnError bool
}

func NewHookExecutor(hook *configuration.Hook) *HookExecutor {
	h := &HookExecutor{
		hook:        hook,
		files:       make(map[string]*os.File),
		tempDirPath: os.TempDir(),
		exitOnError: false,
	}

	for _, step := range hook.Steps {
		h.onStepStatusChange(&step, StepExecutionStatusPending)
		p := path.Join(h.tempDirPath, step.ID+".log")

		logger.Debugf("Creating file for %s:%s at %s", step.PackageRelativePath, step.Name, p)

		file, err := os.Create(p)
		if err != nil {
			logger.Warnf("Error while creating file for %s (%s): %s", step.Name, step.ID, err)
			continue
		} else {
			h.files[step.ID] = file
		}
	}

	return h
}

func (h *HookExecutor) WithExitOnStepError() *HookExecutor {
	h.exitOnError = true

	return h
}

func (h *HookExecutor) getLogFile(id string) (*os.File, error) {
	file, ok := h.files[id]

	if !ok {
		return nil, fmt.Errorf("file for step %s not found, skipping", id)
	}

	return file, nil
}

func (h *HookExecutor) onError(step *configuration.Step, recordedError error) error {
	logger.Tracef("Writing error in file for %s: %s", step.ID, recordedError)
	file, err := h.getLogFile(step.ID)
	if err != nil {
		return err
	}

	// Write the error to the file
	logger.Tracef("Writing error to %s", file.Name())
	_, err = file.WriteString(recordedError.Error())
	if err != nil {
		return err
	}

	if h.exitOnError {
		logger.Errorf("Exiting because of error during step %s: %s", step.ID, recordedError)
		logger.Errorf("For more details, please refer to the step logs at path %s", file.Name())
		os.Exit(1)
	}

	return nil
}

func (h *HookExecutor) onLogRecord(step *configuration.Step, record string) error {
	logger.Tracef("Writing log record in file for %s: %s", step.ID, record)
	file, err := h.getLogFile(step.ID)
	if err != nil {
		logger.Tracef("Error while getting log file for %s: %s", step.ID, err)
		return err
	}

	// Write the log record to the file
	logger.Tracef("Writing %d chars to %s (%s)", len(record), file.Name(), step.Name)
	_, err = file.WriteString(record)
	if err != nil {
		return err
	}

	return nil
}

func (h *HookExecutor) onStepStatusChange(step *configuration.Step, status StepExecutionStatus) error {
	logger.Tracef("Step %s (%s) status changed to %s", step.Name, step.ID, status)
	if status == StepExecutionStatusFailed {
		logger.Errorf("❌ Step %s failed", step.Name)
	} else if status == StepExecutionStatusSuccess {
		logger.Infof("✅ Step %s:%s succeeded", step.PackageRelativePath, step.Name)
	} else if status == StepExecutionStatusRunning {
		logger.Infof("🏃 Step %s:%s is running", step.PackageRelativePath, step.Name)
	} else if status == StepExecutionStatusPending {
		logger.Infof("⏳ Step %s:%s is pending", step.PackageRelativePath, step.Name)
	}

	// If the status is success or failed, close the file
	if status == StepExecutionStatusSuccess || status == StepExecutionStatusFailed {
		file, err := h.getLogFile(step.ID)
		if err != nil {
			return err
		}

		logger.Debugf("Closing log file %s", file.Name())
		err = file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *HookExecutor) RunStep(step *configuration.Step) {
	h.onStepStatusChange(step, StepExecutionStatusRunning)

	commandParts := strings.Split(step.Command, " ")
	commandName := commandParts[0]
	commandArgs := commandParts[1:]

	command := exec.Command(commandName, commandArgs...)
	command.Dir = h.hook.Path

	l := logger.WithFields(map[string]string{
		"Step ID": step.ID,
		"Command": command.String(),
	})

	l.Debug("Running command: ", command.String())
	l.Debug("Path: ", command.Dir)

	stdout, err := command.StdoutPipe()
	if err != nil {
		l.Error("Error while creating stdout pipe: ", err)
		h.onError(step, err)
		return
	}

	err = command.Start()
	if err != nil {
		l.Error("Error while starting command: ", err)
		h.onError(step, err)
		return
	}
	l.Debug("Command started")

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line := scanner.Text()
		h.onLogRecord(step, line)
	}

	err = command.Wait()
	if err != nil {
		l.Error("Error while waiting for command to finish: ", err)
		h.onError(step, err)
	} else {
		l.Debug("Command finished")
		h.onStepStatusChange(step, StepExecutionStatusSuccess)
	}
}

func (h *HookExecutor) Run() {
	steps := h.hook.Steps

	wg := sync.WaitGroup{}

	for _, step := range steps {
		logger.Debugf("Running step %s (%s)", step.Name, step.ID)

		wg.Add(1)

		if step.Serial {
			logger.Debug("Step is blocking")
			h.RunStep(&step)
			wg.Done()
		} else {
			logger.Debug("Step is not blocking")
			go func() {
				h.RunStep(&step)
				wg.Done()
			}()
		}
	}

	wg.Wait()
}
