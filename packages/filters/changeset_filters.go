package filters

import (
	"strings"

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
			return nil, err
		}

		if match {
			filtered = append(filtered, path)
		}
	}

	return filtered, nil
}
