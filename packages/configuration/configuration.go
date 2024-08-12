package configuration

type OnModifiedFileBehaviour string

const (
	AddAndCommit OnModifiedFileBehaviour = "addAndCommit"
	Abort        OnModifiedFileBehaviour = "abort"
)

/*
Global configuration for the Gookme tool.
Behaviours defined here will be used for all hooks to run.

The Gookme global configuration is a JSON file named .gookme.json, located at the root of the repository.
*/
type GookmeGlobalConfiguration struct {
	// The behaviour to use when there are modified files at the end of the hook.
	OnModifiedFiles *OnModifiedFileBehaviour `json:"onModifiedFiles,omitempty"`
}

func DefaultGlobalConfiguration() *GookmeGlobalConfiguration {
	onModifiedFiles := Abort
	return &GookmeGlobalConfiguration{
		OnModifiedFiles: &onModifiedFiles,
	}
}

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

/*
Configuration for a hook.

The Gookme hook configuration is a JSON file named .{hookType}.json, located in a subfolder or at the root of the repository. The hookType is the type of Git hook to attach the configuration to, such as pre-commit, post-merge, etc.
*/
type HookConfiguration struct {
	Steps []StepConfiguration `json:"steps"`
}

type HookType string

const (
	PreCommit     HookType = "pre-commit"
	PrepareCommit HookType = "prepare-commit-msg"
	CommitMsg     HookType = "commit-msg"
	PostCommit    HookType = "post-commit"
	PostMerge     HookType = "post-merge"
	PostRewrite   HookType = "post-rewrite"
	PreRebase     HookType = "pre-rebase"
	PostCheckout  HookType = "post-checkout"
	PrePush       HookType = "pre-push"
)
