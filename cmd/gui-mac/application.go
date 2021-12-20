package main

import (
	"github.com/mniak/pismo/internal/config"
	"github.com/mniak/pismo/internal/folhacerta"
	"github.com/mniak/pismo/internal/keepass"
	"github.com/mniak/pismo/pkg/pismo"
)

type _Application struct {
	ClockManager pismo.ClockManager
	OTPProvider  pismo.OTPProvider
}

func initApplication() (_Application, error) {
	conf, err := config.Load()
	if err != nil {
		return _Application{}, err
	}
	return _Application{
		ClockManager: folhacerta.New(folhacerta.Config{
			Token:   conf.Clock.Token,
			Verbose: true,
		}),
		OTPProvider: keepass.New(keepass.Config{
			Database: conf.OTP.Database,
			Password: conf.OTP.Password,
			OTPEntry: conf.OTP.Entry,
		}),
	}, nil
}
