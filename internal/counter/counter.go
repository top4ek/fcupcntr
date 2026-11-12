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

func Days(since time.Time, now time.Time) int {
	duration := now.Sub(since)
	if duration < 0 {
		return 0
	}
	return int(duration / (24 * time.Hour))
}

func (c Counter) Name() string {
	return c.name
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
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return Counter{}, ErrEmptyName
	}
	return Counter{
		name:           trimmed,
		lastIncidentAt: time.Now(),
	}, nil
}

func (c *Counter) reset(now time.Time) {
	c.lastIncidentAt = now
}
