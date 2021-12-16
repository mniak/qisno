package folhacerta

type Config struct {
	Token string
}

func New(c Config) *_FolhaCertaClockManager {
	return &_FolhaCertaClockManager{
		config: c,
	}
}
