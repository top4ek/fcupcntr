package counter

import (
	"time"
)

func Days(since time.Time, now time.Time) int {
	duration := now.Sub(since)
	if duration < 0 {
		return 0
	}
	return int(duration / (24 * time.Hour))
}
