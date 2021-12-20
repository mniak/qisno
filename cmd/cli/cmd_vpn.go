package main

import (
	"github.com/spf13/cobra"
)

var cmdVPN = &cobra.Command{
	Use: "vpn",
	RunE: func(cmd *cobra.Command, args []string) error {
		w, _, err := app.VPNProvider.Connect()
		if err != nil {
			return err
		}
		err = w()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cmdVPN)
}
