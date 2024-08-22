package cli

import (
	"fmt"

	"github.com/LMaxence/gookme/packages/configuration"
)

func validateHookType(value string) (configuration.HookType, error) {
	switch configuration.HookType(value) {
	case configuration.PreCommit, configuration.CommitMsg, configuration.PostCheckout, configuration.PostMerge, configuration.PostRewrite, configuration.PrePush, configuration.PreRebase, configuration.PostCommit:
		return configuration.HookType(value), nil
	default:
		return "", fmt.Errorf("invalid HookType: %s. Accepted values are: pre-commit, prepare-commit-msg, commit-msg,  post-commit, post-merge, post-rewrite, pre-rebase, post-checkout, pre-push", value)
	}
}
