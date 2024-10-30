package commands

import (
	"github.com/kevsmith/tsetse/internal"
	"github.com/kevsmith/tsetse/internal/pages"
	"github.com/urfave/cli/v2"
)

func AddRunTestsCommand(app *cli.App) {
	command := &cli.Command{
		Name:    "run-tests",
		Aliases: []string{"rt"},
		Usage:   "run Playwright test suite",
		Action: func(ctx *cli.Context) error {
			return runTests(ctx)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "admin-email",
				EnvVars:  []string{"SF_TAX_E2E_USER_SP"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "admin-pw",
				EnvVars:  []string{"SF_TAX_E2E_USER_SP_PASS"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "sp-email",
				EnvVars:  []string{"SF_TAX_E2E_USER_SP_NON_ADMIN"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "sp-pw",
				EnvVars:  []string{"SF_TAX_E2E_USER_SP_NON_ADMIN_PASS"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "client-email",
				EnvVars:  []string{"SF_TAX_E2E_USER_CLIENT"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "client-pw",
				EnvVars:  []string{"SF_TAX_E2E_USER_CLIENT_PASS"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "domain",
				EnvVars:  []string{"SF_TAX_E2E_DOMAIN"},
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "headless",
				Usage: "enable headless browser mode",
			},
		},
	}
	app.Commands = append(app.Commands, command)
}

func runTests(ctx *cli.Context) error {
	testConfig, err := internal.NewTestRunConfig(ctx.String("admin-email"),
		ctx.String("admin-pass"),
		ctx.String("sp-email"),
		ctx.String("sp-pass"),
		ctx.String("client-email"),
		ctx.String("client-pw"), ctx.String("domain"))
	if err != nil {
		return err
	}
	testRun, err := internal.NewTestRun(ctx.Bool("headless"), testConfig)
	if err != nil {
		return err
	}
	loginTest := pages.LoginPageTest{}
	return loginTest.Run(testRun)
}
