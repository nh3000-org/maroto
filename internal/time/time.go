package time

import (
	"time"

	"github.com/nh3000-org/maroto/v2/pkg/metrics"
)

// GetTimeSpent returns a metrics.Time with the time spent in the closure.
func GetTimeSpent(closure func()) *metrics.Time {
	start := time.Now()
	closure()
	return &metrics.Time{
		Value: float64(time.Since(start).Nanoseconds()),
		Scale: metrics.Nano,
	}
}
