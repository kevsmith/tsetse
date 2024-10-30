package main

import (
	"os"

	"github.com/kevsmith/tsetse/internal/commands"
	"github.com/kevsmith/tsetse/internal/flags"
	"github.com/urfave/cli/v2"
)

var VerbosityCount = 0

func newCliApp() *cli.App {
	app := &cli.App{
		Name:                   "tsetse",
		Usage:                  "The highly useful, never annoying, web test runner",
		UseShortOptionHandling: true,
		Flags:                  []cli.Flag{flags.NewVerboseFlag()},
	}
	commands.AddInstallPlaywrightCommand(app)
	commands.AddRunTestsCommand(app)
	return app
}

func main() {
	app := newCliApp()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
