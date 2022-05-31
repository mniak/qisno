package main

import (
	"context"

	"github.com/caseymrm/menuet"
	"golang.design/x/clipboard"
)

func (a _Application) generateOTPMenuItems(ctx context.Context) []*menuet.MenuItem {
	otpItem := &menuet.MenuItem{
		Text: "Loading...",
	}

	result := []*menuet.MenuItem{
		{
			Text: "OTP (click to copy)",
		},
		otpItem,
	}

	go func() {
		code, err := a.OTPProvider.OTP(ctx)
		if err != nil {
			otpItem.Text = "Failed to generate OTP"
			return
		}

		otpItem.Text = code
		otpItem.Clicked = func() {
			clipboard.Write(clipboard.FmtText, []byte(code))
		}
	}()

	return result
}
