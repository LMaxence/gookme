package configuration

import (
	"fmt"

	"github.com/google/uuid"
)

type StepFixture struct {
	id     string
	name   string
	onlyOn *string
	files  []string
}

func (sf *StepFixture) ToStep() *Step {
	return &Step{
		Name:    sf.name,
		OnlyOn:  sf.onlyOn,
		Command: "echo Hello world",
		ID:      sf.id,
		Files:   sf.files,
	}
}

func (sf *StepFixture) Copy() *StepFixture {
	return &StepFixture{
		id:     uuid.NewString(),
		name:   sf.name,
		onlyOn: sf.onlyOn,
		files:  append([]string{}, sf.files...),
	}
}

func NewStepFixture() *StepFixture {
	return &StepFixture{
		id:     uuid.NewString(),
		name:   uuid.NewString(),
		onlyOn: nil,
		files:  make([]string, 0),
	}
}

func (sf *StepFixture) WithOnlyOn(onlyOn string) *StepFixture {
	sf.onlyOn = &onlyOn
	return sf
}

func (sf *StepFixture) WithFiles(files ...string) *StepFixture {
	sf.files = files
	return sf
}

type HookFixture struct {
	path  string
	files []string
	steps []*StepFixture
}

func NewHookFixture(path string) *HookFixture {
	return &HookFixture{
		path:  path,
		steps: make([]*StepFixture, 0),
		files: make([]string, 0),
	}
}

func (hf *HookFixture) Copy() *HookFixture {
	steps := make([]*StepFixture, 0)
	for _, step := range hf.steps {
		steps = append(steps, step.Copy())
	}

	return &HookFixture{
		path:  hf.path,
		steps: steps,
		files: append([]string{}, hf.files...),
	}
}

func (hf *HookFixture) ToHook() Hook {
	hook := Hook{
		Path:  hf.path,
		Steps: []Step{},
		Files: hf.files,
	}

	for i, step := range hf.steps {
		hook.Steps = append(hook.Steps, Step{
			OnlyOn:  step.onlyOn,
			Command: "echo Hello world",
			ID:      fmt.Sprintf("step%d", i),
			Files:   step.files,
		})
	}

	return hook
}

func (hf *HookFixture) WithFiles(files ...string) *HookFixture {
	hf.files = append(hf.files, files...)
	return hf
}

func (hf *HookFixture) WithStep(step ...*StepFixture) *HookFixture {
	hf.steps = append(hf.steps, step...)
	return hf
}
