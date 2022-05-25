package main

import (
	"encoding/json"
	"fmt"

	"github.com/mniak/qisno/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cmdMetabase)
}

var cmdMetabase = cobra.Command{
	Use:  "metabase <envname>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envname := args[0]

		if flagAutoLogin {
			awsLogin(envname, false)
		}

		profile, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "Profile")
		handle(err, "failed to load the profile name")

		region, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "MetabaseRegion")
		handle(err, "failed to load the region name")

		target, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "MetabaseTarget")
		handle(err, "failed to load the target value")

		localPort, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "MetabasePort")
		handle(err, "failed to load the local port number")

		jsonParams, err := json.Marshal(map[string]any{
			"portNumber":      []string{"80"},
			"localPortNumber": []string{localPort},
		})
		handle(err, "could not build JSON parameters for metabase command")

		err = utils.ExecInteractive(
			"aws",
			"--profile", profile,
			"ssm", "start-session",
			"--region", region,
			"--target", target,
			"--document-name", "AWS-StartPortForwardingSession",
			"--parameters", string(jsonParams),
		)
		handle(err)
	},
}
