package counter_test

import (
	"testing"
	"time"

	"fcupcntr/internal/counter"
	"github.com/stretchr/testify/assert"
)

func TestNewEmptyName(t *testing.T) {
	_, err := counter.New("")

	assert.ErrorIs(t, err, counter.ErrEmptyName)
}

func TestNew(t *testing.T) {
	name := "Paralelepipedik"
	before := time.Now()
	cntr, err := counter.New(name)
	if err != nil {
		t.Fatalf("%v", err)
	}
	after := time.Now()

	assert.Equal(t, name, cntr.Name())
	assert.False(t, cntr.LastIncidentAt().Before(before))
	assert.False(t, cntr.LastIncidentAt().After(after))
}
