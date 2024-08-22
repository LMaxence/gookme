package configuration

/*
Configuration for a single step in a hook.
*/
type StepConfiguration struct {
	// The named of the step. Displayed in the UI and used in it to index steps and hooks
	Name string `json:"name"`

	// The command that will be invoked in `execSync`
	Command string `json:"command"`

	// A pattern string describing which changed files will trigger this step
	OnlyOn *string `json:"onlyOn,omitempty"`

	// Should this step be awaited before starting the next one
	Serial *bool `json:"serial,omitempty"`

	// Does this step extend a shared step
	From *string `json:"from,omitempty"`
}

type Step struct {
	ID                  string
	Name                string
	Command             string
	ExecutedCommand     string
	OnlyOn              *string
	Serial              bool
	From                *string
	PackageRelativePath string
}
