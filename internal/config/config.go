package config

type Config struct {
	Clock ClockConfig
	OTP   OTPConfig
}

type ClockConfig struct {
	Token string
}

type OTPConfig struct {
	Database string
	Password string
	Entry    string
}
