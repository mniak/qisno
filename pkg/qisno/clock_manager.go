package qisno

import (
	"context"
	"time"
)

type ClockInfo struct {
	Running        bool
	Empty          bool
	FirstStartTime time.Time
	LastStartTime  time.Time
	LastEndTime    time.Time
	TotalTimeToday time.Duration
}

type ClockManager interface {
	Query(ctx context.Context) (ClockInfo, error)
	Clock(ctx context.Context) error
}
