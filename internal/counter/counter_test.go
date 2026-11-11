package counter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	now := time.Date(2026, time.November, 10, 12, 0, 0, 0, time.UTC)
	counter := Counter{
		Name:           "Skynet",
		LastIncidentAt: now.Add(-48 * time.Hour),
	}
	counter.reset(now)
	assert.Equal(t, now, counter.LastIncidentAt)
	assert.Equal(t, 0, Days(counter.LastIncidentAt, now))
	assert.Equal(t, "Skynet", counter.Name)
}

func TestDays(t *testing.T) {
	now := time.Date(2026, time.November, 10, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		since  time.Time
		result int
	}{
		{
			since:  now.Add(-24 * time.Hour),
			result: 1,
		},
		{
			since:  now.Add(-23*time.Hour - 59*time.Minute),
			result: 0,
		},
		{
			since:  now.Add(48 * time.Hour),
			result: 0,
		},
	}

	for _, example := range cases {
		assert.Equal(t, example.result, Days(example.since, now))
	}
}
