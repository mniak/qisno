package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mniak/pismo/domain"
)

type ClockManager struct{}

func (c *ClockManager) Query(ctx context.Context) (domain.ClockInfo, error) {
	time.Sleep(200 * time.Millisecond)
	return domain.ClockInfo{
		Running:        gofakeit.Bool(),
		Empty:          gofakeit.Bool(),
		FirstStartTime: gofakeit.DateRange(time.Now().Add(-20*time.Hour), time.Now().Add(-2*time.Hour)),
		LastEndTime:    gofakeit.DateRange(time.Now().Add(-1*time.Hour), time.Now()),
		TotalTimeToday: time.Duration(gofakeit.IntRange(15, 60*12)) * time.Minute,
	}, nil
}

func (c *ClockManager) Clock(ctx context.Context) error {
	fmt.Println("o ponto foi marcado")
	return nil
}
