package utils

import (
	"fmt"
	"math"
	"time"
)

func FormatDuration(d time.Duration) string {
	hours := int(math.Trunc(d.Hours()))
	minutes := int(math.Trunc(d.Minutes())) - hours*60
	if hours == 0 {
		return fmt.Sprintf("%d min", minutes)
	}
	return fmt.Sprintf("%dh%02d", hours, minutes)
}
