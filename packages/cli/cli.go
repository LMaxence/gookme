package cli

import (
	gitclient "github.com/LMaxence/gookme/packages/git-client"
	"github.com/LMaxence/gookme/packages/logging"
	"github.com/LMaxence/gookme/packages/meta"
	"github.com/urfave/cli/v2"
)

var logger = logging.NewLogger("cli")

func NewCLI() *cli.App {
	app := &cli.App{
		Name:    "gookme",
		Version: meta.GOOKME_CLI_VERSION,
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
			RunCommand,
			InitCommand,
			CleanCommand,
		},
	}

	return app
}
