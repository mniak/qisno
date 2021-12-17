package main

import (
	"context"
	"time"

	"github.com/mniak/pismo"
	"github.com/mniak/pismo/internal/config"
	"github.com/mniak/pismo/internal/folhacerta"
	"github.com/mniak/pismo/internal/keepass"
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
			Token: conf.Clock.Token,
		}),
		OTPProvider: keepass.New(keepass.Config{
			Database: conf.OTP.Database,
			Password: conf.OTP.Password,
			OTPEntry: conf.OTP.Entry,
		}),
	}, nil
}

func newContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return ctx
}
