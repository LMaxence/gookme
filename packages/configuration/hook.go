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
	PreCommitHookType     HookType = "pre-commit"
	PrepareCommitHookType HookType = "prepare-commit-msg"
	CommitMsgHookType     HookType = "commit-msg"
	PostCommitHookType    HookType = "post-commit"
	PostMergeHookType     HookType = "post-merge"
	PostRewriteHookType   HookType = "post-rewrite"
	PreRebaseHookType     HookType = "pre-rebase"
	PostCheckoutHookType  HookType = "post-checkout"
	PrePushHookType       HookType = "pre-push"
)

var ALL_HOOKS = []HookType{
	PreCommitHookType,
	PrepareCommitHookType,
	CommitMsgHookType,
	PostCommitHookType,
	PostMergeHookType,
	PostRewriteHookType,
	PreRebaseHookType,
	PostCheckoutHookType,
	PrePushHookType,
}

var ALL_HOOKS_STRING = []string{
	string(PreCommitHookType),
	string(PrepareCommitHookType),
	string(CommitMsgHookType),
	string(PostCommitHookType),
	string(PostMergeHookType),
	string(PostRewriteHookType),
	string(PreRebaseHookType),
	string(PostCheckoutHookType),
	string(PrePushHookType),
}

type Hook struct {
	Path  string
	Files []string
	Steps []Step
}
