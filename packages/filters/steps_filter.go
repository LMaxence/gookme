package filters

/* ===================== Steps filters =====================

Step filters are components that will filter the steps that should
be executed in hooks. They are based on the changeset that
is calculated by the resolving strategies.

With a given changeset, the filters will determine which steps of the hooks
should be executed, depending on their path or configuration.

=========================================================== */

import "github.com/LMaxence/gookme/packages/configuration"

// FilterStepsWithOnlyOn filters the steps that should be executed based on the
// changeset and the `onlyOn` configuration filed of the steps.
func FilterStepsWithOnlyOn(
	changedPaths []string,
	hooks []configuration.Hook,
) []configuration.Hook {
	filtered := make([]configuration.Hook, 0)

	// For each hook, check if the hook directory is contained by one of the changeset's elements.
	// If it is the case, add the hook to the list of hooks to execute

	for _, hook := range hooks {
		steps := make([]configuration.Step, 0)
		packagePaths := filterChangesetWithPrefix(changedPaths, hook.Path)

		for _, step := range hook.Steps {
			onlyOn := step.OnlyOn

			// If the step does not have an `onlyOn` field, add it to the list of steps to execute
			if onlyOn == nil {
				steps = append(steps, step)
				continue
			}

			// If the step has an `onlyOn` field, filter the changeset with the pattern
			packageChangedPathsWithPattern, err := filterChangesetWithPattern(packagePaths, *onlyOn)
			if err != nil {
				continue
			}
			step.Files = append(step.Files, packageChangedPathsWithPattern...)

			// Only add the step to the list of steps to execute if it has remaining matched files
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
