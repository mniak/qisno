package main

import (
	"context"
	"fmt"
	"os"

	"github.com/caseymrm/menuet"
	"golang.design/x/clipboard"
)

func (a _Application) generateOTPMenuItems(ctx context.Context) []menuet.MenuItem {
	code, err := a.OTPProvider.OTP(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to generate OTP:", err)
		return []menuet.MenuItem{
			{
				Text: "Failed to generate OTP",
			},
		}
	}

	return []menuet.MenuItem{
		{
			Text: "OTP (click to copy)",
		},
		{
			Text: code,
			Clicked: func() {
				clipboard.Write(clipboard.FmtText, []byte(code))
			},
		},
	}
}
