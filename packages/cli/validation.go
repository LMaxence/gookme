package cli

import (
	"fmt"
	"strings"

	"github.com/LMaxence/gookme/packages/configuration"
)

func validateHookType(value string) (configuration.HookType, error) {
	switch configuration.HookType(value) {
	case configuration.PreCommitHookType, configuration.CommitMsgHookType, configuration.PostCheckoutHookType, configuration.PostMergeHookType, configuration.PostRewriteHookType, configuration.PrePushHookType, configuration.PreRebaseHookType, configuration.PostCommitHookType, configuration.PrepareCommitHookType:
		return configuration.HookType(value), nil
	default:
		return "", fmt.Errorf("invalid HookType: %s. Accepted values are: %s", value, strings.Join(configuration.ALL_HOOKS_STRING, ", "))
	}
}
