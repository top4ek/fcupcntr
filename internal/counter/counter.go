package counter

import (
	"errors"
	"strings"
	"time"
)

type Counter struct {
	name           string
	lastIncidentAt time.Time
}

var ErrEmptyName = errors.New("name missing")

func (c Counter) Name() string {
	return c.name
}

func (c Counter) Days(current time.Time) int {
	return days(c.lastIncidentAt, current)
}

func (c *Counter) Reset() {
	c.reset(time.Now())
}

func (c *Counter) LastIncidentAt() time.Time {
	return c.lastIncidentAt
}

func New(name string) (Counter, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return Counter{}, ErrEmptyName
	}
	return Counter{
		name:           trimmed,
		lastIncidentAt: time.Now(),
	}, nil
}

func days(since time.Time, now time.Time) int {
	duration := now.Sub(since)
	if duration < 0 {
		return 0
	}
	return int(duration / (24 * time.Hour))
}

func (c *Counter) reset(now time.Time) {
	c.lastIncidentAt = now
}
