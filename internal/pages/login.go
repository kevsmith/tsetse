package pages

import (
	"fmt"

	"github.com/kevsmith/tsetse/internal"
	"github.com/playwright-community/playwright-go"
)

/*
import { Page, expect } from '@playwright/test';

export class LoginPage {
    constructor(private page: Page) {}

    private usernameInput = '#login-form_username';
    private passwordInput = '#login-form_password';
    private submitButton = 'button[type="submit"]';
    private catalogSelector = 'a[data-analytics-id="solutionCatalog_CatalogButton"]';

    async goto() {
        await this.page.goto('/login', { timeout: 10000 });
    }

    async login(username: string, password: string) {
        // Fill in username and submit
        await this.page.fill(this.usernameInput, username);
        await this.page.click(this.submitButton);

        // Wait for password field, fill in password, and submit
        await this.page.waitForSelector(this.passwordInput, { timeout: 10000 });
        await this.page.fill(this.passwordInput, password);
        await this.page.click(this.submitButton);

        // Confirm successful login by checking for the authenticated page element
        await this.page.waitForSelector(this.catalogSelector, { timeout: 20000 });
    }

    async verifyLoginSuccess() {
        const catalogButton = this.page.locator(this.catalogSelector);
        await expect(catalogButton).toBeVisible();
        console.log("Successfully logged in as service provider.");
    }
}
*/

type PageTestArgs map[string]string

type PageTest interface {
	Run(testRun *internal.TestRun, args PageTestArgs) error
}

const usernameInput = "#login-form_username"
const passwordInput = "#login-form_password"
const submitButton = "button[type=\"submit\"]"
const catalogSelector = "a[data-analytics-id=\"solutionCatalog_CatalogButton\"]"

type LoginPageTest struct{}

func (lpt *LoginPageTest) Run(testRun *internal.TestRun) error {
	page, err := testRun.NewPage()
	if err != nil {
		return err
	}
	_, err = page.Goto(testRun.Config.MakeURL("/login"), playwright.PageGotoOptions{})
	if err != nil {
		return err
	}
	err = page.Fill(usernameInput, testRun.Config.AdminUser.Email)
	if err != nil {
		return err
	}
	if testRun.Config.Verbose {
		fmt.Printf("Submitted admin user email '%s'\n", testRun.Config.AdminUser.Email)
	}
	err = page.Click(submitButton)
	if err != nil {
		return err
	}
	err = page.Fill(passwordInput, testRun.Config.AdminUser.Password)
	if err != nil {
		return err
	}
	if testRun.Config.Verbose {
		fmt.Printf("Submitted admin user password\n")
	}
	err = page.Click(submitButton)
	if err != nil {
		return err
	}
	if testRun.Config.Verbose {
		fmt.Printf("Login successful\n")
	}
	return nil
}
