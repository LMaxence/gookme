package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/LMaxence/gookme/packages/logging"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/google/uuid"
)

var logger = logging.NewLogger("configuration")

type GookmeConfiguration struct {
	Global GookmeGlobalConfiguration
	Hooks  []Hook
}

func LoadGlobalConfiguration(directory string) (*GookmeGlobalConfiguration, error) {
	logger.Infof("Loading global configuration from %s", directory)

	matches, err := filepath.Glob(filepath.Join(directory, ".gookme.json"))
	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		logger.Info("No global configuration found, using default values")
		return DefaultGlobalConfiguration(), nil
	}

	match := matches[0]
	logger.Infof("Loading configuration from %s", match)

	// Open the file and unmarshal its content into a `GookmeGlobalConfiguration` instance
	content, err := os.ReadFile(match)
	if err != nil {
		logger.Warnf("Error while reading global configuration file %s: %s", match, err)
		return nil, err
	}

	var global GookmeGlobalConfiguration
	err = json.Unmarshal(content, &global)
	if err != nil {
		logger.Warnf("Error while reading global configuration file %s: %s", match, err)
		return nil, err
	}

	return &global, nil
}

func LoadHooksConfiguration(directory string, hookType HookType) ([]Hook, error) {
	pattern := fmt.Sprintf("**/hooks/%s.json", string(hookType))
	logger.Infof("Loading hooks configuration from %s with pattern %s", directory, pattern)

	fsys := os.DirFS(directory)
	matches, err := doublestar.Glob(fsys, pattern)
	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		logger.Infof("No hooks configuration found for hook type %s", hookType)
		return nil, nil
	}

	hooks := make([]Hook, 0, len(matches))

	for _, match := range matches {
		path := filepath.Join(directory, match)

		logger.Infof("Loading hook configuration from %s", path)
		content, err := os.ReadFile(path)
		if err != nil {
			logger.Warnf("Error while reading hook configuration file %s: %s", path, err)
			return nil, err
		}

		var hook HookConfiguration
		err = json.Unmarshal(content, &hook)
		if err != nil {
			logger.Warnf("Error while reading hook configuration file %s: %s", path, err)
			return nil, err
		}

		packageRelativePath := strings.Replace(filepath.Dir(path), directory, "", 1)
		packageRelativePath = strings.TrimPrefix(packageRelativePath, "/")
		packageRelativePath = strings.TrimSuffix(packageRelativePath, "/hooks")

		logger.Debugf("Package relative path: %s", packageRelativePath)

		steps := make([]Step, 0, len(hook.Steps))
		for _, step := range hook.Steps {
			steps = append(steps, Step{
				Name:                step.Name,
				PackageRelativePath: packageRelativePath,
				Command:             step.Command,
				ExecutedCommand:     exec.Command(step.Command).String(),
				Serial:              step.Serial != nil && *step.Serial,
				OnlyOn:              step.OnlyOn,
				From:                step.From,
				ID:                  uuid.NewString(),
			})
		}

		hooks = append(hooks, Hook{
			Path:  filepath.Dir(filepath.Dir(path)),
			Steps: steps,
		})
	}

	return hooks, nil
}

func LoadConfiguration(directory string, hookType HookType) (*GookmeConfiguration, error) {
	global, err := LoadGlobalConfiguration(directory)
	if err != nil {
		return nil, err
	}

	hooks, err := LoadHooksConfiguration(directory, hookType)
	if err != nil {
		return nil, err
	}

	return &GookmeConfiguration{
		Global: *global,
		Hooks:  hooks,
	}, nil
}
