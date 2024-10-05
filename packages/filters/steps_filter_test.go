package filters

import (
	"testing"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/stretchr/testify/assert"
)

type StepFilterTestCase struct {
	ID        string
	Changeset []string
	Hooks     []configuration.Hook
	Expected  []configuration.Hook
}

func getStepFiltersTestCases() []StepFilterTestCase {
	cases := []StepFilterTestCase{}

	// Case 1: One of the step has a onlyOn pattern that matches the changeset

	stepJson := configuration.NewStepFixture().WithOnlyOn("*.json")
	stepGo := configuration.NewStepFixture().WithOnlyOn("*.go")

	hook := configuration.NewHookFixture("foo/bar").WithFiles().WithStep(stepJson, stepGo)
	expectedHook := configuration.NewHookFixture("foo/bar").WithFiles("foo/bar/item.json").WithStep(stepJson.Copy().WithFiles("foo/bar/item.json"))

	cases = append(cases, StepFilterTestCase{
		ID:        "One of the step has a onlyOn pattern that matches the changeset",
		Changeset: []string{"foo/bar/item.json", "bar/foo/item.go", "foo/main.go"},
		Hooks:     []configuration.Hook{hook.ToHook()},
		Expected:  []configuration.Hook{expectedHook.ToHook()},
	})

	// Case 2: None of the steps have a onlyOn pattern that matches the changeset

	stepJson = configuration.NewStepFixture().WithOnlyOn("*.json")
	stepAll := configuration.NewStepFixture()

	hook = configuration.NewHookFixture("foo/bar").WithFiles().WithStep(stepJson, stepAll)
	expectedHook = configuration.NewHookFixture("foo/bar").WithFiles("foo/bar/item.go").WithStep(stepAll.Copy())

	cases = append(cases, StepFilterTestCase{
		ID:        "Match all files",
		Changeset: []string{"foo/bar/item.go", "bar/foo/item.go", "foo/main.go"},
		Hooks:     []configuration.Hook{hook.ToHook()},
		Expected:  []configuration.Hook{expectedHook.ToHook()},
	})

	return cases
}

func TestFilterStepsWithOnlyOn(t *testing.T) {
	for _, tc := range getStepFiltersTestCases() {
		t.Run(tc.ID, func(t *testing.T) {
			res := FilterStepsWithOnlyOn(tc.Changeset, tc.Hooks)

			for i, hook := range res {
				assert.Equal(t, tc.Expected[i].Path, hook.Path)
				assert.Equal(t, len(tc.Expected[i].Steps), len(hook.Steps))

				for j, step := range hook.Steps {
					assert.Equal(t, tc.Expected[i].Steps[j].Name, step.Name)
					assert.Equal(t, tc.Expected[i].Steps[j].Files, step.Files)
				}
			}
		})
	}
}
