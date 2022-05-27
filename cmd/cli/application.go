package main

import (
	"context"
	"time"

	"github.com/mniak/qisno/internal/adp"
	"github.com/mniak/qisno/internal/config"
	"github.com/mniak/qisno/internal/keepass"
	"github.com/mniak/qisno/internal/wrappers"
	"github.com/mniak/qisno/pkg/qisno"
	"github.com/spf13/cobra"
)

type _Application struct {
	ClockManager    qisno.ClockManager
	OTPProvider     qisno.OTPProvider
	VPNProvider     qisno.VPNProvider
	PasswordManager qisno.PasswordManager
}

func initApplication(cmd *cobra.Command) (_Application, error) {
	conf, err := config.Load()
	if err != nil {
		return _Application{}, err
	}

	keepass := keepass.New(keepass.Config{
		Database: conf.OTP.Database,
		Password: conf.OTP.Password,
		OTPEntry: conf.OTP.Entry,
	})
	adp := adp.New(adp.Config{
		Username: conf.Clock.Username,
		Password: conf.Clock.Password,
		Verbose:  flagVerbose,
	})
	return _Application{
		ClockManager: &adp,
		OTPProvider:  keepass,
		VPNProvider: wrappers.NewOpenfortiVPN(wrappers.OpenfortiVPNConfig{
			Host:        conf.VPN.Host,
			Username:    conf.VPN.Username,
			Password:    conf.VPN.Password,
			TrustedCert: conf.VPN.TrustedCert,
			Verbose:     flagVerbose,
		}),
		PasswordManager: keepass,
	}, nil
}

func newContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return ctx
}
