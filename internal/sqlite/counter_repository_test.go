package sqlite

import (
	"log"
	"testing"

	"fcupcntr/internal/counter"

	"github.com/stretchr/testify/assert"
)

func TestFindNotFound(t *testing.T) {
	ctx := t.Context()
	database, err := Open(ctx, ":memory:")
	if err != nil {
		t.Fatalf("Database creation: %v", err)
	}
	defer func() {
		err := database.Close()
		if err != nil {
			log.Printf("Database closing error: %v", err)
		}
	}()

	repository := NewCounterRepository(database)
	_, err = repository.Find(ctx, "unknown")
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestCreate(t *testing.T) {
	ctx := t.Context()
	database, err := Open(ctx, ":memory:")
	if err != nil {
		t.Fatalf("Database creation: %v", err)
	}
	defer func() {
		err := database.Close()
		if err != nil {
			log.Printf("Database closing error: %v", err)
		}
	}()

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

	found, err := repository.Find(ctx, created.ID())
	if err != nil {
		t.Fatalf("Counter finding: %v", err)
	}

	assert.Equal(t, created.ID(), found.ID())
	assert.Equal(t, created.Name(), found.Name())
	assert.True(t, created.CreatedAt().Equal(found.CreatedAt()), "created_at")
	assert.True(t, created.LastIncidentAt().Equal(found.LastIncidentAt()), "last_incident_at")

}
