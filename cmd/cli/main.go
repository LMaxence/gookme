package main

import (
	"os"

	"github.com/LMaxence/gookme/packages/cli"
	"github.com/LMaxence/gookme/packages/logging"
)

var logger = logging.NewLogger("cli")

func main() {
	app := cli.NewCLI()

	if err := app.Run(os.Args); err != nil {
		logger.Errorf("Error while running CLI: %s", err)
		os.Exit(1)
	}
}
