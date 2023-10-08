package utilz

import (
	"time"
)

func MeasureTime(start time.Time) time.Duration {
	elapsed := time.Since(start)
	return elapsed
}
