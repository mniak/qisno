package keepass

import (
	"context"
	"net/url"
	"time"

	"github.com/pquerna/otp/totp"
)

func (k *Keepass) OTP(ctx context.Context) (string, error) {
	otp, err := k.Attribute(k.config.OTPEntry, "otp")
	if err != nil {
		return "", err
	}

	otpurl, err := url.Parse(otp)
	if err != nil {
		return "", err
	}

	secret := otpurl.Query().Get("secret")
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return code, nil
}
