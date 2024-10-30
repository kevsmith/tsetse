package commands

import (
	"github.com/kevsmith/tsetse/internal/flags"
	"github.com/playwright-community/playwright-go"
	"github.com/urfave/cli/v2"
)

func AddInstallPlaywrightCommand(app *cli.App) {
	command := &cli.Command{
		Name:    "install-playwright",
		Aliases: []string{"ip"},
		Usage:   "install Playwright and its dependencies",
		Action: func(ctx *cli.Context) error {
			return installPlaywright()
		},
	}
	app.Commands = append(app.Commands, command)
}

func installPlaywright() error {
	return playwright.Install(&playwright.RunOptions{
		Browsers: []string{"chromium"},
		Verbose:  flags.IsVerbose(),
	})
}
