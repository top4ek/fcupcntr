package counter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	name := "Skynet"
	now := fixedTime()
	counter := Counter{
		name:           name,
		lastIncidentAt: now.Add(-48 * time.Hour),
	}
	counter.reset(now)
	assert.Equal(t, now, counter.LastIncidentAt())
	assert.Equal(t, name, counter.Name())
}

func TestDaysStruct(t *testing.T) {
	now := fixedTime()
	counter := Counter{
		name:           "Skynet",
		lastIncidentAt: now.Add(-49 * time.Hour),
	}
	assert.Equal(t, 2, counter.Days(now))
	counter.reset(now)
	assert.Equal(t, 0, counter.Days(now))
}

func fixedTime() time.Time {
	return time.Date(2026, time.November, 10, 12, 0, 0, 0, time.UTC)
}
