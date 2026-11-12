package counter

import (
	"errors"
	"time"
)

type Counter struct {
	Name           string
	lastIncidentAt time.Time
}

func Days(since time.Time, now time.Time) int {
	duration := now.Sub(since)
	if duration < 0 {
		return 0
	}
	return int(duration / (24 * time.Hour))
}

func (c Counter) Days(now time.Time) int {
	return Days(c.lastIncidentAt, now)
}

func (c *Counter) Reset() {
	c.reset(time.Now())
}

func (c *Counter) LastIncidentAt() time.Time {
	return c.lastIncidentAt
}

func New(name string) (Counter, error) {
	if name == "" {
		return Counter{}, errors.New("name missing")
	}
	return Counter{
		Name:           name,
		lastIncidentAt: time.Now(),
	}, nil
}

func (c *Counter) reset(now time.Time) {
	c.lastIncidentAt = now
}
