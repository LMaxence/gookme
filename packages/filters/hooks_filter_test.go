package filters

import (
	"testing"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/stretchr/testify/assert"
)

type HookFilterTestCase struct {
	ID        string
	Changeset []string
	Hooks     []configuration.Hook
	Expected  []configuration.Hook
}

func getHookFiltersTestCases() []HookFilterTestCase {
	cases := []HookFilterTestCase{}

	// Case 1: Hook is in the changeset

	hook := configuration.NewHookFixture("foo/bar")
	expectedHook := configuration.NewHookFixture("foo/bar").WithFiles("foo/bar/item.json")

	cases = append(cases, HookFilterTestCase{
		ID:        "Hook is in the changeset",
		Changeset: []string{"foo/bar/item.json", "bar/foo/item.go", "foo/main.go"},
		Hooks:     []configuration.Hook{hook.ToHook()},
		Expected:  []configuration.Hook{expectedHook.ToHook()},
	})

	// Case 2: Hook is not in the changeset

	hook = configuration.NewHookFixture("foo/bar").WithFiles()

	cases = append(cases, HookFilterTestCase{
		ID:        "Hook is not in the changeset",
		Changeset: []string{"bar/foo/item.go", "foo/main.go"},
		Hooks:     []configuration.Hook{hook.ToHook()},
		Expected:  []configuration.Hook{},
	})

	return cases
}

func TestFilterHooksWithChangeset(t *testing.T) {
	for _, tc := range getHookFiltersTestCases() {
		t.Run(tc.ID, func(t *testing.T) {
			res := FilterHooksWithChangeset(tc.Changeset, tc.Hooks)

			for i, hook := range res {
				assert.Equal(t, tc.Expected[i].Path, hook.Path)
				assert.Equal(t, tc.Expected[i].Files, hook.Files)
			}
		})
	}
}
