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

	kp := keepass.New(keepass.Config{
		Database: conf.OTP.Database,
		Password: conf.OTP.Password,
		OTPEntry: conf.OTP.Entry,
	})
	adp := adp.New(adp.Config{
		Username: conf.Clock.Username,
		Password: conf.Clock.Password,
		Verbose:  flagVerbose,
	})
	var vpnConfig wrappers.OpenfortiVPNConfigLoader
	if conf.VPN.UsePasswordManager {
		vpnConfig = wrappers.VPNConfigFromPasswordManager(kp)
	} else {
		vpnConfig = wrappers.VPNConfigInline(conf.VPN)
	}
	return _Application{
		ClockManager:    &adp,
		OTPProvider:     kp,
		VPNProvider:     wrappers.NewOpenfortiVPN(vpnConfig, flagVerbose),
		PasswordManager: kp,
	}, nil
}

func newContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return ctx
}
