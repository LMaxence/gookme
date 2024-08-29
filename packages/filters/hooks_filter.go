package filters

import "github.com/LMaxence/gookme/packages/configuration"

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
