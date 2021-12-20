package pismo

type (
	WaitFunc       func() error
	DisconnectFunc func() error
)

type VPNProvider interface {
	Connect() (WaitFunc, DisconnectFunc, error)
}
