package sqlite

import (
	"testing"

	"fcupcntr/internal/counter"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := t.Context()
	database, err := Open(ctx, ":memory:")
	if err != nil {
		t.Fatalf("Database creation: %v", err)
	}
	defer database.Close()
	repository := NewCounterRepository(database)
	cntr, err := counter.New("test")
	if err != nil {
		t.Fatalf("Counter creation: %v", err)
	}

	created, err := repository.Create(ctx, cntr)
	if err != nil {
		t.Fatalf("DB record creation: %v", err)
	}
	assert.Equal(t, cntr.Name(), created.Name())
	assert.Equal(t, cntr.CreatedAt(), created.CreatedAt())
	assert.Equal(t, cntr.LastIncidentAt(), created.LastIncidentAt())
	assert.NotEqual(t, cntr.ID(), created.ID()) // TODO: Чекнуть через будущий Find
}
