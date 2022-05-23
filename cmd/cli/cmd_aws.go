package main

import (
	"fmt"

	"github.com/mniak/qisno/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cmdAws)
	cmdAws.AddCommand(&cmdAwsConfig)

	cmdAws.AddCommand(&cmdAwsLogin)
	cmdAwsLogin.Flags().Bool("force", false, "Do not consider the login cache")
}

var cmdAws = cobra.Command{
	Use: "aws",
}

var cmdAwsConfig = cobra.Command{
	Use:  "config <envname>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envname := args[0]

		profile, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "Profile")
		handle(err, "failed to load the profile name")

		url, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "URL")
		handle(err, "failed to load the URL")

		oktaUsername, err := app.PasswordManager.Username("Provedores/Okta")
		handle(err, "failed to load the username")

		oktaPassword, err := app.PasswordManager.Password("Provedores/Okta")
		handle(err, "failed to load the password")

		err = utils.ExecInteractive(
			"saml2aws", "configure",
			"-a", profile,
			"--idp-provider=Okta", "--mfa=OKTA",
			fmt.Sprintf("--profile=%s", profile),
			fmt.Sprintf("--url=%s", url),
			fmt.Sprintf("--username=%s", oktaUsername),
			fmt.Sprintf("--password=%s", oktaPassword),
			"--skip-prompt",
		)
		handle(err)
	},
}

var cmdAwsLogin = cobra.Command{
	Use:  "login <envname>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envname := args[0]

		profile, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "Profile")
		handle(err, "failed to load the profile name")

		force, err := cmd.Flags().GetBool("force")
		handle(err)

		oktaUsername, err := app.PasswordManager.Username("Provedores/Okta")
		handle(err, "failed to load the username")

		oktaPassword, err := app.PasswordManager.Password("Provedores/Okta")
		handle(err, "failed to load the password")

		otpCode, err := app.OTPProvider.OTP(newContext())
		handle(err)

		saml2awsArgs := []string{
			"login",
			"-a", profile,
			"--idp-provider=Okta", "--mfa=OKTA",
			fmt.Sprintf("--profile=%s", profile),
			fmt.Sprintf("--mfa-token=%s", otpCode),
			fmt.Sprintf("--username=%s", oktaUsername),
			fmt.Sprintf("--password=%s", oktaPassword),
			"--skip-prompt",
		}
		if force {
			saml2awsArgs = append(saml2awsArgs, "--force")
		}
		err = utils.ExecInteractive("saml2aws", saml2awsArgs...)
		handle(err)
	},
}
