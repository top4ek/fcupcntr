package counter_test

import (
	"testing"
	"time"

	"fcupcntr/internal/counter"

	"github.com/stretchr/testify/assert"
)

func TestDays(t *testing.T) {
	counter, err := counter.New("Test")
	if err != nil {
		t.Fatalf("%v", err)
	}
	cases := []struct {
		name    string
		current time.Time
		result  int
	}{
		{
			name:    "One day in the past",
			current: counter.LastIncidentAt().Add(-24 * time.Hour),
			result:  0,
		},
		{
			name:    "Almost one day but not",
			current: counter.LastIncidentAt().Add(23*time.Hour + 59*time.Minute),
			result:  0,
		},
		{
			name:    "In 1 days",
			current: counter.LastIncidentAt().Add(24 * time.Hour),
			result:  1,
		},
		{
			name:    "In 2 days",
			current: counter.LastIncidentAt().Add(48 * time.Hour),
			result:  2,
		},
	}

	for _, example := range cases {
		t.Run(example.name, func(t *testing.T) {
			assert.Equal(t, example.result, counter.Days(example.current))
		})
	}
}
