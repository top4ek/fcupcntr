package counter_test

import (
	"testing"
	"time"

	"fcupcntr/internal/counter"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyName(t *testing.T) {
	cases := []struct {
		name    string
		payload string
	}{
		{
			name:    "Empty name",
			payload: "",
		},
		{
			name:    "Spaced name",
			payload: "    ",
		},
		{name: "Special spaced name",
			payload: "\t\n\r",
		},
	}

	for _, example := range cases {
		t.Run(example.name, func(t *testing.T) {
			_, err := counter.New(example.payload)
			assert.ErrorIs(t, err, counter.ErrEmptyName)
		})
	}
}

func TestNewCreatedAt(t *testing.T) {
	before := time.Now()
	cntr, err := counter.New("Test")
	if err != nil {
		t.Fatalf("%v", err)
	}
	after := time.Now()

	assert.Equal(t, cntr.CreatedAt(), cntr.LastIncidentAt())
	assert.False(t, cntr.CreatedAt().Before(before))
	assert.False(t, cntr.CreatedAt().After(after))
}

func TestNew(t *testing.T) {
	name := "Paralelepipedik"
	before := time.Now()
	cntr, err := counter.New("\t" + name + "\n")
	if err != nil {
		t.Fatalf("%v", err)
	}
	after := time.Now()

	assert.Equal(t, name, cntr.Name())
	assert.False(t, cntr.LastIncidentAt().Before(before))
	assert.False(t, cntr.LastIncidentAt().After(after))
}
