package config

type Config struct {
	Clock ClockConfig
	OTP   OTPConfig
	VPN   VPNConfig
}

type ClockConfig struct {
	Token string
}

type OTPConfig struct {
	Database string
	Password string
	Entry    string
}

type VPNConfig struct {
	Host        string
	Username    string
	Password    string
	TrustedCert string
}
