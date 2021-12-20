package main

import (
	"github.com/mniak/qisno/internal/config"
	"github.com/mniak/qisno/internal/folhacerta"
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
	return _Application{
		Title: conf.Title,
		ClockManager: folhacerta.New(folhacerta.Config{
			Token:   conf.Clock.Token,
			Verbose: true,
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
			Verbose:     true,
		}),
	}, nil
}
