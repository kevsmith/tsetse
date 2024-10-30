package internal

import (
	"net/url"

	"github.com/kevsmith/tsetse/internal/flags"
	"github.com/playwright-community/playwright-go"
)

type UserInfo struct {
	Email    string
	Password string
}

type TestRunConfig struct {
	AdminUser    *UserInfo
	NonAdminUser *UserInfo
	ClientUser   *UserInfo
	Domain       *url.URL
	Verbose      bool
}

type TestRun struct {
	Config         *TestRunConfig
	pw             *playwright.Playwright
	browser        playwright.Browser
	browserContext playwright.BrowserContext
}

func NewTestRunConfig(adminEmail, adminPass, nonAdminEmail, nonAdminPass, clientEmail, clientPass, domain string) (*TestRunConfig, error) {
	parsed, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	return &TestRunConfig{
		AdminUser:    &UserInfo{Email: adminEmail, Password: adminPass},
		NonAdminUser: &UserInfo{Email: nonAdminEmail, Password: nonAdminPass},
		ClientUser:   &UserInfo{Email: clientEmail, Password: clientPass},
		Domain:       parsed,
		Verbose:      flags.IsVerbose(),
	}, nil
}

func (trc *TestRunConfig) MakeURL(path string) string {
	updated := trc.Domain.JoinPath(path)
	return updated.String()
}

func NewTestRun(headless bool, config *TestRunConfig) (*TestRun, error) {
	pw, err := playwright.Run(&playwright.RunOptions{
		SkipInstallBrowsers: true,
		Verbose:             flags.IsVerbose(),
		Browsers:            []string{"chromium"},
	})
	if err != nil {
		return nil, err
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: &headless,
	})
	if err != nil {
		return nil, err
	}
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{})
	if err != nil {
		browser.Close()
		return nil, err
	}
	return &TestRun{
		Config:         config,
		pw:             pw,
		browser:        browser,
		browserContext: context,
	}, nil
}

func (tr *TestRun) Close() {
	if tr.browserContext != nil {
		tr.browserContext.Close()
		tr.browserContext = nil
	}
	if tr.browser != nil {
		tr.browser.Close()
		tr.browser = nil
	}
}

func (tr *TestRun) NewPage() (playwright.Page, error) {
	return tr.browserContext.NewPage()
}
