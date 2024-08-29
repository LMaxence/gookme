package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var filterChangesetWithPrefixTestCases = []struct {
	ID        string
	prefix    string
	changeset []string
	expected  []string
}{
	{
		ID:        "Match prefix",
		prefix:    "foo/bar",
		changeset: []string{"foo/bar/item.json", "bar/foo/item.go", "foo/main.go"},
		expected:  []string{"foo/bar/item.json"},
	},
	{
		ID:        "No match",
		prefix:    "bar",
		changeset: []string{"foo/bar/item.json", "foo/item.go", "foo/main.go"},
		expected:  []string{},
	},
	{
		ID:        "No match - other prefix",
		prefix:    "fizz/buzz",
		changeset: []string{"foo/bar/item.json", "bar/foo/item.go", "foo/main.go"},
		expected:  []string{},
	},
	{
		ID:        "Empty changeset",
		prefix:    "foo/bar",
		changeset: []string{},
		expected:  []string{},
	},
}

func TestFilterChangesetWithPrefix(t *testing.T) {
	for _, tc := range filterChangesetWithPrefixTestCases {
		t.Run(tc.ID, func(t *testing.T) {
			res := filterChangesetWithPrefix(tc.changeset, tc.prefix)
			assert.Equal(t, tc.expected, res)
		})
	}
}

var filterChangesetWithPatternTestCases = []struct {
	ID        string
	pattern   string
	changeset []string
	expected  []string
	error     bool
}{
	{
		ID:        "Match pattern",
		pattern:   "foo/**",
		changeset: []string{"foo/bar/item.json", "bar/foo/item.go", "foo/main.go"},
		expected:  []string{"foo/bar/item.json", "bar/foo/item.go", "foo/main.go"},
		error:     false,
	},
	{
		ID:        "Match pattern - no prefix",
		pattern:   "*.json",
		changeset: []string{"foo/bar/item.json", "bar/foo/item.go", "foo/main.go"},
		expected:  []string{"foo/bar/item.json"},
		error:     false,
	},
	{
		ID:        "No match",
		pattern:   "*.js",
		changeset: []string{"foo/bar/item.json", "bar/foo/item.go", "foo/main.go"},
		expected:  []string{},
		error:     false,
	},
	{
		ID:        "Empty changeset",
		pattern:   "*.go",
		changeset: []string{},
		expected:  []string{},
		error:     false,
	},
}

func TestFilterChangesetWithPattern(t *testing.T) {
	for _, tc := range filterChangesetWithPatternTestCases {
		t.Run(tc.ID, func(t *testing.T) {
			res, err := filterChangesetWithPattern(tc.changeset, tc.pattern)
			if tc.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expected, res)
		})
	}
}
