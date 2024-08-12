package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

type Hook struct {
	Path          string
	Configuration HookConfiguration
}

type GookmeConfiguration struct {
	Global GookmeGlobalConfiguration
	Hooks  []Hook
}

func LoadGlobalConfiguration(directory string) (*GookmeGlobalConfiguration, error) {
	matches, err := filepath.Glob(filepath.Join(directory, ".gookme.json"))
	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		return DefaultGlobalConfiguration(), nil
	}

	match := matches[0]

	// Open the file and unmarshal its content into a `GookmeGlobalConfiguration` instance
	content, err := os.ReadFile(match)
	if err != nil {
		return nil, err
	}

	var global GookmeGlobalConfiguration
	err = json.Unmarshal(content, &global)
	if err != nil {
		return nil, err
	}

	return &global, nil
}

func LoadHooksConfiguration(directory string, hookType HookType) ([]Hook, error) {
	fsys := os.DirFS(directory)
	matches, err := doublestar.Glob(fsys, fmt.Sprintf("**/%s.json", string(hookType)))
	if err != nil {
		return nil, err
	}

	hooks := make([]Hook, 0, len(matches))

	for _, match := range matches {
		path := filepath.Join(directory, match)
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		var hook HookConfiguration
		err = json.Unmarshal(content, &hook)
		if err != nil {
			return nil, err
		}

		hooks = append(hooks, Hook{
			Path:          filepath.Dir(path),
			Configuration: hook,
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
