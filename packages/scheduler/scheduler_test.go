package scheduler

import (
	"fmt"
	"testing"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/stretchr/testify/assert"
)

var defaultSteps = []configuration.Step{
	{
		Command: "echo 'Hello World'",
		ID:      "step1",
	},
}

func TestFilterScheduledHooksWithChangeset(t *testing.T) {
	changedPaths := []string{"/packages/path1/somefile.go"}
	hooks := []configuration.Hook{
		{
			Path:  "/packages/path1",
			Steps: defaultSteps,
		},
		{
			Path:  "/packages/path2",
			Steps: defaultSteps,
		},
	}

	filteredHooks := FilterHooksWithChangeset(changedPaths, hooks)
	assert.Equal(t, 1, len(filteredHooks))
	assert.Equal(t, hooks[0], filteredHooks[0])
}

func TestFilterScheduledHooksWithLargerChangeset(t *testing.T) {
	changedPaths := []string{"/packages/path1/somefile.go", "/packages/path2/somefile.go"}
	hooks := []configuration.Hook{
		{
			Path:  "/packages/path1",
			Steps: defaultSteps,
		},
		{
			Path:  "/packages/path2",
			Steps: defaultSteps,
		},
	}

	filteredHooks := FilterHooksWithChangeset(changedPaths, hooks)
	assert.Equal(t, 2, len(filteredHooks))
	assert.Equal(t, hooks[0], filteredHooks[0])
	assert.Equal(t, hooks[1], filteredHooks[1])
}

func TestFilterScheduledHooksWithChangesetWithNoMatchingHooks(t *testing.T) {
	changedPaths := []string{"/packages/path1/somefile.go"}
	hooks := []configuration.Hook{
		{
			Path:  "/packages/path2",
			Steps: defaultSteps,
		},
	}

	filteredHooks := FilterHooksWithChangeset(changedPaths, hooks)
	assert.Equal(t, 0, len(filteredHooks))
}

var filterChangesetWithPrefixTestCases = []struct {
	changedPaths []string
	prefix       string
	expected     []string
}{
	{
		changedPaths: []string{"/packages/path1/somefile.go", "/packages/path2/somefile.go"},
		prefix:       "/packages/path1",
		expected:     []string{"/packages/path1/somefile.go"},
	},
	{
		changedPaths: []string{"/packages/path1/somefile.go", "/packages/path2/somefile.go"},
		prefix:       "/packages/path2",
		expected:     []string{"/packages/path2/somefile.go"},
	},
	{
		changedPaths: []string{"/packages/path1/somefile.go", "/packages/path2/somefile.go"},
		prefix:       "/packages/path3",
		expected:     []string{},
	},
}

func TestFilterChangesetWithPrefix(t *testing.T) {
	for _, testCase := range filterChangesetWithPrefixTestCases {
		filteredPaths := filterChangesetWithPrefix(testCase.changedPaths, testCase.prefix)
		assert.Equal(t, testCase.expected, filteredPaths)
	}
}

var filterChangesetWithPatternTestCases = []struct {
	changedPaths []string
	pattern      string
	expected     []string
}{
	{
		changedPaths: []string{"/packages/path1/somefile.go", "/packages/path1/somefile.json"},
		pattern:      "*.json",
		expected:     []string{"/packages/path1/somefile.json"},
	},
	{
		changedPaths: []string{"/packages/path1/somefile.go", "/packages/path1/somefile.json"},
		pattern:      "test/*.json",
		expected:     []string{},
	},
	{
		changedPaths: []string{"/packages/path1/somefile.go", "/packages/path1/somefile.json"},
		pattern:      "packages/path1/*.json",
		expected:     []string{"/packages/path1/somefile.json"},
	},
	{
		changedPaths: []string{"/packages/path1/somefile.go", "/packages/path1/somefile.json"},
		pattern:      "*.{go,json}",
		expected:     []string{"/packages/path1/somefile.go", "/packages/path1/somefile.json"},
	},
}

func TestFilterChangesetWithPatter(t *testing.T) {
	for _, testCase := range filterChangesetWithPatternTestCases {
		filteredPaths, err := filterChangesetWithPattern(testCase.changedPaths, testCase.pattern)
		assert.Nil(t, err)
		assert.Equal(t, testCase.expected, filteredPaths)
	}
}

func fixtureHookConfiguration(
	path string,
	onlyOn []*string,
) configuration.Hook {
	hook := configuration.Hook{
		Path:  path,
		Steps: []configuration.Step{},
	}

	for i, pattern := range onlyOn {
		hook.Steps = append(hook.Steps, configuration.Step{
			OnlyOn:  pattern,
			Command: "echo Hello world",
			ID:      fmt.Sprintf("step%d", i),
		})
	}

	return hook
}

var goStarPattern = "*.go"
var somefileStarPattern = "somefile.*"
var filterStepsWithOnlyOnTestCases = []struct {
	changedPaths []string
	hooks        []configuration.Hook
	expected     []configuration.Hook
}{
	{
		changedPaths: []string{"/packages/path1/somefile.go"},
		hooks: []configuration.Hook{
			fixtureHookConfiguration("/packages/path1", []*string{nil}),
		},
		expected: []configuration.Hook{
			fixtureHookConfiguration("/packages/path1", []*string{nil}),
		},
	},
	{
		changedPaths: []string{"/packages/path1/somefile.go"},
		hooks: []configuration.Hook{
			fixtureHookConfiguration("/packages/path1", []*string{&goStarPattern}),
		},
		expected: []configuration.Hook{
			fixtureHookConfiguration("/packages/path1", []*string{&goStarPattern}),
		},
	},
	{
		changedPaths: []string{"/packages/path1/somefile.json"},
		hooks: []configuration.Hook{
			fixtureHookConfiguration("/packages/path1", []*string{&goStarPattern}),
		},
		expected: []configuration.Hook{},
	},
	{
		changedPaths: []string{"/packages/path1/somefile.json"},
		hooks: []configuration.Hook{
			fixtureHookConfiguration("/packages/path1", []*string{&somefileStarPattern}),
		},
		expected: []configuration.Hook{
			fixtureHookConfiguration("/packages/path1", []*string{&somefileStarPattern}),
		},
	},
	{
		changedPaths: []string{"/packages/path1/some-other-file.json"},
		hooks: []configuration.Hook{
			fixtureHookConfiguration("/packages/path1", []*string{&somefileStarPattern}),
		},
		expected: []configuration.Hook{},
	},
}

func TestFilterStepsWithOnlyOn(t *testing.T) {
	for _, testCase := range filterStepsWithOnlyOnTestCases {
		filteredHooks := FilterStepsWithOnlyOn(testCase.changedPaths, testCase.hooks)

		assert.Equal(t, len(testCase.expected), len(filteredHooks))
		for i, expectedHook := range testCase.expected {
			filteredHook := filteredHooks[i]

			assert.Equal(t, expectedHook.Path, filteredHook.Path)
			assert.Equal(t, len(expectedHook.Steps), len(filteredHook.Steps))
		}
	}
}
