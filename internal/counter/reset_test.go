package counter_test

import (
	"testing"
	"time"

	"fcupcntr/internal/counter"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	name := "Skynet"
	counter, err := counter.New(name)
	if err != nil {
		t.Fatalf("%v", err)
	}
	oldCreatedAt := counter.CreatedAt()
	before := time.Now()
	counter.Reset()
	after := time.Now()
	assert.Equal(t, name, counter.Name())
	assert.Equal(t, oldCreatedAt, counter.CreatedAt())
	assert.False(t, counter.LastIncidentAt().Before(before))
	assert.False(t, counter.LastIncidentAt().After(after))
}
