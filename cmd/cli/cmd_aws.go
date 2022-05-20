package main

import (
	"fmt"

	"github.com/mniak/qisno/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cmdAws)
	cmdAws.AddCommand(&cmdAwsConfig)
}

var cmdAws = cobra.Command{
	Use: "aws",
}

var cmdAwsConfig = cobra.Command{
	Use:  "config <envname>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// password=$(ask_password "Type-in the Password Manager password: ")

		envname := args[0]

		// profile=$(passmgr-secret $password "AWS/$envname" Profile)
		profile, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "Profile")
		handle(err, "failed to load the profile")

		// url=$(passmgr-secret $password "AWS/$envname" URL)
		url, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "URL")
		handle(err, "failed to load the URL")

		// okta_username=$(passmgr-secret $password "Provedores/Okta" UserName)
		oktaUsername, err := app.PasswordManager.Username("Provedores/Okta")
		handle(err, "failed to load the username")

		// okta_password=$(passmgr-secret $password "Provedores/Okta" Password)
		oktaPassword, err := app.PasswordManager.Password("Provedores/Okta")
		handle(err, "failed to load the password")

		// saml2aws configure -a "$profile" "--idp-provider=Okta" "--mfa=OKTA" "--profile=$profile" "--url=$url" "--username=$okta_username" "--password=$okta_password" --skip-prompt
		result, err := utils.ExecSimple(
			"saml2aws", "configure",
			"-a", profile,
			"--idp-provider=Okta", "--mfa=OKTA",
			fmt.Sprintf("--profile=%s", profile),
			fmt.Sprintf("--url=%s", url),
			fmt.Sprintf("--username=$okta_username", oktaUsername),
			fmt.Sprintf("--password=%s", oktaPassword),
			"--skip-prompt",
		)
		handle(err)

		fmt.Println(result)
	},
}
