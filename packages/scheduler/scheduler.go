package scheduler

import (
	"strings"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/LMaxence/gookme/packages/logging"
	"github.com/bmatcuk/doublestar/v4"
)

var logger = logging.NewLogger("scheduler")

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
			return nil, err
		}

		if match {
			filtered = append(filtered, path)
		}
	}

	return filtered, nil
}

func FilterHooksWithChangeset(
	changedPaths []string,
	hooks []configuration.Hook,
) []configuration.Hook {
	filtered := make([]configuration.Hook, 0)

	// For each hook, check if the hook directory is contained by one of the changeset's elements.
	// If it is the case, add the hook to the list of hooks to execute
	// If the hook is not in the changeset, skip it

	for _, hook := range hooks {
		hook.Files = append(hook.Files, filterChangesetWithPrefix(changedPaths, hook.Path)...)

		if len(hook.Files) > 0 {
			filtered = append(filtered, hook)
		} else {
			logger.Debugf("Hook %s did not match any file, dropping", hook.Path)
		}
	}

	// Return the list of hooks to execute
	return filtered
}

func FilterStepsWithOnlyOn(
	changedPaths []string,
	hooks []configuration.Hook,
) []configuration.Hook {
	filtered := make([]configuration.Hook, 0)

	// For each hook, check if the hook directory is contained by one of the changeset's elements.
	// If it is the case, add the hook to the list of hooks to execute

	for _, hook := range hooks {
		steps := make([]configuration.Step, 0)
		changedPaths := filterChangesetWithPrefix(changedPaths, hook.Path)

		for _, step := range hook.Steps {
			onlyOn := step.OnlyOn

			if onlyOn == nil {
				steps = append(steps, step)
				continue
			}

			changedPathsWithPattern, err := filterChangesetWithPattern(changedPaths, *onlyOn)
			if err != nil {
				continue
			}
			step.Files = append(step.Files, changedPathsWithPattern...)

			if len(step.Files) > 0 {
				steps = append(steps, step)
			} else {
				logger.Debugf("Step %s:%s did not match any file using pattern %s, dropping", step.PackageRelativePath, step.Name, *onlyOn)
			}
		}

		if len(steps) > 0 {
			filtered = append(filtered, configuration.Hook{
				Path:  hook.Path,
				Steps: steps,
			})
		}
	}

	// Return the list of hooks to execute
	return filtered
}
