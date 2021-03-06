package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var otpCmd = &cobra.Command{
	Use: "otp",
	Run: func(cmd *cobra.Command, args []string) {
		code, err := app.OTPProvider.OTP(newContext())
		handle(err)

		print, err := cmd.Flags().GetBool("print")
		handle(err)
		if print {
			fmt.Println(code)
		} else {
			clipboard.Write(clipboard.FmtText, []byte(code))
			fmt.Println("OTP copied to clipboard!")
		}
	},
}

func init() {
	rootCmd.AddCommand(otpCmd)
	otpCmd.Flags().Bool("print", false, "Print the value instead of copying to clipboard")
}
