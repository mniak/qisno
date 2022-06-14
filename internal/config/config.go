package config

type Config struct {
	Title string
	Clock ClockConfig
	OTP   OTPConfig
	VPN   VPNConfig
}

type ClockConfig struct {
	Username string
	Password string
}

type OTPConfig struct {
	Database string
	Password string
	Entry    string
}

type VPNConfig struct {
	UsePasswordManager bool

	Host        string
	Username    string
	Password    string
	TrustedCert string
}
