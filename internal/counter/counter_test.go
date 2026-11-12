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

func TestDays(t *testing.T) {
	now := fixedTime()
	cases := []struct {
		name   string
		since  time.Time
		result int
	}{
		{
			name:   "One day",
			since:  now.Add(-24 * time.Hour),
			result: 1,
		},
		{
			name:   "Almost one day but not",
			since:  now.Add(-23*time.Hour - 59*time.Minute),
			result: 0,
		},
		{
			name:   "Accident from future",
			since:  now.Add(48 * time.Hour),
			result: 0,
		},
	}

	for _, example := range cases {
		t.Run(example.name, func(t *testing.T) {
			assert.Equal(t, example.result, Days(example.since, now))
		})
	}
}

func fixedTime() time.Time {
	return time.Date(2026, time.November, 10, 12, 0, 0, 0, time.UTC)
}
