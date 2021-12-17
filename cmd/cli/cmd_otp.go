package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var otpCmd = &cobra.Command{
	Use: "otp",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
		code, err := app.OTPProvider.OTP(ctx)
		cobra.CheckErr(err)

		print, err := cmd.Flags().GetBool("print")
		cobra.CheckErr(err)
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
