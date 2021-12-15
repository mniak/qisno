package fake

import (
	"context"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type OTPProvider struct{}

func (o *OTPProvider) OTP(ctx context.Context) (string, error) {
	time.Sleep(900 * time.Millisecond)
	return gofakeit.Digit() +
		gofakeit.Digit() +
		gofakeit.Digit() +
		gofakeit.Digit() +
		gofakeit.Digit() +
		gofakeit.Digit(), nil
}
