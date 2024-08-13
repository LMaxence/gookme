package scheduler

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/bmatcuk/doublestar/v4"
)

func filterChangesetWithPrefix(
	changedPaths []string,
	prefix string,
) []string {
	filtered := make([]string, 0)
	for _, path := range changedPaths {
		if strings.HasPrefix(path, prefix) {
			filtered = append(filtered, path)
		}
	}

	return filtered
}

func filterChangesetWithPattern(
	changedPaths []string,
	pattern string,
) ([]string, error) {
	filtered := make([]string, 0)

	if !strings.HasPrefix(pattern, "**") {
		pattern = "**/" + pattern
	}

	for _, path := range changedPaths {

		match, err := doublestar.Match(pattern, path)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		if match {
			filtered = append(filtered, path)
		}
	}

	return filtered, nil
}

func filterHooksWithChangeset(
	changedPaths []string,
	hooks []configuration.Hook,
) []configuration.Hook {
	filtered := make([]configuration.Hook, 0)

	// For each hook, check if the hook directory is contained by one of the changeset's elements.
	// If it is the case, add the hook to the list of hooks to execute
	// If the hook is not in the changeset, skip it
	for _, hook := range hooks {
		hookDir := filepath.Dir(hook.Path)
		matchedPaths := filterChangesetWithPrefix(changedPaths, hookDir)
		if len(matchedPaths) > 0 {
			filtered = append(filtered, hook)
		}
	}

	// Return the list of hooks to execute
	return filtered
}

func filterStepsWithOnlyOn(
	changedPaths []string,
	hooks []configuration.Hook,
) []configuration.Hook {
	filtered := make([]configuration.Hook, 0)

	// For each hook, check if the hook directory is contained by one of the changeset's elements.
	// If it is the case, add the hook to the list of hooks to execute

	for _, hook := range hooks {
		steps := make([]configuration.StepConfiguration, 0)
		hookDir := filepath.Dir(hook.Path)
		changedPaths := filterChangesetWithPrefix(changedPaths, hookDir)

		for _, step := range hook.Configuration.Steps {
			onlyOn := step.OnlyOn

			if onlyOn == nil {
				steps = append(steps, step)
				continue
			}

			changedPathsWithPattern, err := filterChangesetWithPattern(changedPaths, *onlyOn)
			if err != nil {
				continue
			}

			if len(changedPathsWithPattern) > 0 {
				steps = append(steps, step)
			}
		}

		if len(steps) > 0 {
			filtered = append(filtered, configuration.Hook{
				Path: hook.Path,
				Configuration: configuration.HookConfiguration{
					Steps: steps,
				},
			})
		}
	}

	// Return the list of hooks to execute
	return filtered
}
