package keepass

type Config struct {
	Database string
	Password string
	OTPEntry string
}

func New(c Config) *Keepass {
	return &Keepass{
		config: c,
	}
}
