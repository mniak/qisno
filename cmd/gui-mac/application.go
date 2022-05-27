package main

import (
	"github.com/mniak/qisno/internal/adp"
	"github.com/mniak/qisno/internal/config"
	"github.com/mniak/qisno/internal/keepass"
	"github.com/mniak/qisno/internal/wrappers"
	"github.com/mniak/qisno/pkg/qisno"
)

type _Application struct {
	Title        string
	ClockManager qisno.ClockManager
	OTPProvider  qisno.OTPProvider
	VPNProvider  qisno.VPNProvider
}

func initApplication() (_Application, error) {
	conf, err := config.Load()
	if err != nil {
		return _Application{}, err
	}
	adp := adp.New(adp.Config{
		Username: conf.Clock.Username,
		Password: conf.Clock.Password,
		Verbose:  true,
	})
	return _Application{
		Title:        conf.Title,
		ClockManager: &adp,
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
			Verbose:     true,
		}),
	}, nil
}
