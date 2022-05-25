package main

import (
	"fmt"

	"github.com/mniak/qisno/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cmdK8s)
	cmdK8s.AddCommand(&cmdKubeConfig)
}

var cmdK8s = cobra.Command{
	Use:     "k8s",
	Aliases: []string{"kubernetes", "kube", "k"},
}

var cmdKubeConfig = cobra.Command{
	Use:  "config <envname>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envname := args[0]

		if flagAutoLogin {
			awsLogin(envname, false)
		}

		profile, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "Profile")
		handle(err, "failed to load the profile name")

		region, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "Region")
		handle(err, "failed to load the region name")

		cluster, err := app.PasswordManager.Attribute(fmt.Sprintf("AWS/%s", envname), "ClusterName")
		handle(err, "failed to load the cluster name")

		err = utils.ExecInteractive(
			"aws", "eks", "update-kubeconfig",
			"--profile", profile,
			"--region", region,
			"--name", cluster,
		)
		handle(err)
	},
}
