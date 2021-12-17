package folhacerta

type Config struct {
	Token   string
	Verbose bool
}

func New(c Config) *_FolhaCertaClockManager {
	return &_FolhaCertaClockManager{
		config: c,
	}
}
