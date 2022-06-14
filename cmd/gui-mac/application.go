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
	// wrappers.VPNConfigwrappers.OpenfortiVPNConfig
	// vpnConfigLoader := ifthen(condition, wrappers.VPNConfigwrappers.OpenfortiVPNConfig{
	// 	Host:        conf.VPN.Host,
	// 	Username:    conf.VPN.Username,
	// 	Password:    conf.VPN.Password,
	// 	TrustedCert: conf.VPN.TrustedCert,
	// })
	kp := keepass.New(keepass.Config{
		Database: conf.OTP.Database,
		Password: conf.OTP.Password,
		OTPEntry: conf.OTP.Entry,
	})
	var vpnConfig wrappers.OpenfortiVPNConfigLoader
	if conf.VPN.UsePasswordManager {
		vpnConfig = wrappers.VPNConfigFromPasswordManager(kp)
	} else {
		vpnConfig = wrappers.VPNConfigInline(conf.VPN)
	}
	return _Application{
		Title:        conf.Title,
		ClockManager: &adp,
		OTPProvider:  kp,
		VPNProvider:  wrappers.NewOpenfortiVPN(vpnConfig, false),
	}, nil
}
