package main

import (
	"github.com/mniak/qisno/internal/config"
	"github.com/mniak/qisno/internal/folhacerta"
	"github.com/mniak/qisno/internal/keepass"
	"github.com/mniak/qisno/pkg/qisno"
)

type _Application struct {
	Title        string
	ClockManager qisno.ClockManager
	OTPProvider  qisno.OTPProvider
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
	}, nil
}
