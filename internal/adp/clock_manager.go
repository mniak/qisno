package adp

import (
	"context"
	"sync"

	"github.com/mniak/adpexpert"
	"github.com/mniak/qisno/pkg/qisno"
)

type ADPClockManager struct {
	config     Config
	client     adpexpert.Client
	loggedIn   bool
	loginMutex sync.Mutex
}

func New(c Config) ADPClockManager {
	return ADPClockManager{
		config: c,
		client: adpexpert.Client{},
	}
}

func (cm *ADPClockManager) ensureLogin() error {
	cm.loginMutex.Lock()
	defer cm.loginMutex.Unlock()

	if cm.loggedIn {
		return nil
	}

	err := cm.client.Login(cm.config.Username, cm.config.Password)
	if err != nil {
		return err
	}
	cm.loggedIn = true
	return nil
}

func (cm *ADPClockManager) Query(ctx context.Context) (qisno.ClockInfo, error) {
	if err := cm.ensureLogin(); err != nil {
		return qisno.ClockInfo{}, err
	}
	punches, err := cm.client.GetLastPunches()
	if err != nil {
		return qisno.ClockInfo{}, err
	}
	if punches.LastPunches == nil {
		return qisno.ClockInfo{
			Running: false,
		}, nil
	}

	// punchesToday := make([]models.Punch, 0)
	// for _, punch := range punches.LastPunches {
	// }

	return qisno.ClockInfo{
		Running: false,
	}, nil
}

func (cm *ADPClockManager) Clock(ctx context.Context) error {
	if err := cm.ensureLogin(); err != nil {
		return err
	}
	if err := cm.client.PunchIn(); err != nil {
		return err
	}
	return nil
}
