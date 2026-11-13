package counter_test

import (
	"testing"
	"time"

	"fcupcntr/internal/counter"

	"github.com/stretchr/testify/assert"
)

func TestRestore(t *testing.T) {
	created := time.Date(2026, time.November, 10, 12, 0, 0, 0, time.UTC)
	last := created.Add(time.Hour)
	cntr := counter.Restore("id", "name", created, last)
	assert.Equal(t, "id", cntr.ID())
	assert.Equal(t, "name", cntr.Name())
	assert.Equal(t, created, cntr.CreatedAt())
	assert.Equal(t, last, cntr.LastIncidentAt())
	assert.Equal(t, 2, cntr.Days(last.Add(49*time.Hour)))
}
