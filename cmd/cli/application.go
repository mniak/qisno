package main

import (
	"context"
	"time"

	"github.com/mniak/pismo/internal/config"
	"github.com/mniak/pismo/internal/folhacerta"
	"github.com/mniak/pismo/internal/keepass"
	"github.com/mniak/pismo/internal/wrappers"
	"github.com/mniak/pismo/pkg/pismo"
	"github.com/spf13/cobra"
)

type _Application struct {
	ClockManager pismo.ClockManager
	OTPProvider  pismo.OTPProvider
	VPNProvider  pismo.VPNProvider
}

func initApplication(cmd *cobra.Command) (_Application, error) {
	conf, err := config.Load()
	if err != nil {
		return _Application{}, err
	}
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return _Application{}, err
	}
	return _Application{
		ClockManager: folhacerta.New(folhacerta.Config{
			Token:   conf.Clock.Token,
			Verbose: verbose,
		}),
		OTPProvider: keepass.New(keepass.Config{
			Database: conf.OTP.Database,
			Password: conf.OTP.Password,
			OTPEntry: conf.OTP.Entry,
		}),
		VPNProvider: wrappers.NewOpenfortiVPN(wrappers.OpenfortiVPNConfig{
			Host:        conf.VPN.Host,
			Username:    conf.VPN.Username,
			Password:    conf.VPN.Password,
			TrustedCert: conf.VPN.TrustedCert,
		}),
	}, nil
}

func newContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return ctx
}
