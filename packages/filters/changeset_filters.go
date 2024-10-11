package filters

import (
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

// filterChangesetWithPrefix filters a list of paths with a given prefix
func filterChangesetWithPrefix(
	changedPaths []string,
	prefix string,
) []string {
	filtered := make([]string, 0)
	for _, path := range changedPaths {
		// We just need the path to start with the prefix
		if strings.HasPrefix(path, prefix) {
			filtered = append(filtered, path)
		}
	}

	return filtered
}

// filterChangesetWithPattern filters a list of paths with a given pattern
func filterChangesetWithPattern(
	changedPaths []string,
	pattern string,
) ([]string, error) {
	filtered := make([]string, 0)

	// If the pattern does not start with "**", we add it, so the pattern can match
	// at any level of the path
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
