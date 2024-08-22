package cli

import (
	"fmt"
	"strings"

	"github.com/LMaxence/gookme/packages/configuration"
)

func validateHookType(value string) (configuration.HookType, error) {
	switch configuration.HookType(value) {
	case configuration.PreCommit, configuration.CommitMsg, configuration.PostCheckout, configuration.PostMerge, configuration.PostRewrite, configuration.PrePush, configuration.PreRebase, configuration.PostCommit, configuration.PrepareCommit:
		return configuration.HookType(value), nil
	default:
		return "", fmt.Errorf("invalid HookType: %s. Accepted values are: %s", value, strings.Join(configuration.ALL_HOOKS_STRING, ", "))
	}
}
