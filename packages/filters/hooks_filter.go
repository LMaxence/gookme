package filters

/* ==================== Hooks filters ====================

Hook filters are components that will filter the hooks that should
be executed. They are based on the changeset that is calculated by
the resolving strategies.

With a given changeset, the filters will determine which hooks should
be executed, depending on their path.

=========================================================== */

import "github.com/LMaxence/gookme/packages/configuration"

// FilterHooksWithChangeset filters the hooks that should be executed based on the
// changeset and the `path` configuration field of the hooks.
func FilterHooksWithChangeset(
	changedPaths []string,
	hooks []configuration.Hook,
) []configuration.Hook {
	filtered := make([]configuration.Hook, 0)

	// For each hook, check if the hook directory is contained by one of the changeset's elements.
	// If it is the case, add the hook to the list of hooks to execute
	// If the hook path is not in the changeset, skip it

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
