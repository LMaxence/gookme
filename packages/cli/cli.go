package cli

import (
	"fmt"

	"github.com/LMaxence/gookme/packages/configuration"
	gitclient "github.com/LMaxence/gookme/packages/git-client"
	"github.com/LMaxence/gookme/packages/logging"
	"github.com/urfave/cli/v2"
)

var logger = logging.NewLogger("cli")

func validateHookType(value string) (configuration.HookType, error) {
	switch configuration.HookType(value) {
	case configuration.PreCommit, configuration.CommitMsg, configuration.PostCheckout, configuration.PostMerge, configuration.PostRewrite, configuration.PrePush, configuration.PreRebase, configuration.PostCommit:
		return configuration.HookType(value), nil
	default:
		return "", fmt.Errorf("invalid HookType: %s. Accepted values are: pre-commit, prepare-commit-msg, commit-msg,  post-commit, post-merge, post-rewrite, pre-rebase, post-checkout, pre-push", value)
	}
}

func NewCLI() *cli.App {
	app := &cli.App{
		Name:    "gookme",
		Version: "0.1.0",
		Usage:   "Run Git hooks with ease",
		Before: func(ctx *cli.Context) error {
			logger.Infof("Running gookme %s", ctx.App.Version)
			dir, err := gitclient.GetRepoPath(nil)

			if err != nil {
				logger.Errorf("Error while getting current working directory: %s", err)
				return err
			}
			logger.Infof("Working directory: %s", dir)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    string(RunCommand),
				Aliases: []string{"r"},
				Usage:   "load and run git hooks based on staged Git changes",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "type",
						Aliases: []string{"t"},
						Value:   "pre-commit",
						Usage:   "The type of Git hook to run. Default is pre-commit, but accepted values are: pre-commit, prepare-commit-msg, commit-msg,  post-commit, post-merge, post-rewrite, pre-rebase, post-checkout, pre-push",
					},
				},
				Action: func(cContext *cli.Context) error {
					hookType, err := validateHookType(cContext.String("type"))
					if err != nil {
						return err
					}
					args := RunCommandArguments{
						HookType: hookType,
					}
					return Run(args)
				},
			},
		},
	}

	return app
}
