package filters

import "github.com/LMaxence/gookme/packages/configuration"

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
