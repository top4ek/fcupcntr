package counter

import (
	"time"
)

type Counter struct {
	Name           string
	LastIncidentAt time.Time
}

func Days(since time.Time, now time.Time) int {
	duration := now.Sub(since)
	if duration < 0 {
		return 0
	}
	return int(duration / (24 * time.Hour))
}

func (c *Counter) Reset() {
	c.reset(time.Now())
}

func (c *Counter) reset(now time.Time) {
	c.LastIncidentAt = now
}
