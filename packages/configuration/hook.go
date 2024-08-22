package configuration

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

type Hook struct {
	Path  string
	Steps []Step
}
