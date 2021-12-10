package qisno

import "context"

type OTPProvider interface {
	OTP(ctx context.Context) (string, error)
}
