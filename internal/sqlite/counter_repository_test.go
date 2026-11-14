package sqlite

import (
	"log"
	"testing"
	"time"

	"fcupcntr/internal/counter"

	"github.com/stretchr/testify/assert"
)

func TestUpdateIncidentNotFound(t *testing.T) {
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
	cntr := counter.Restore("test", "name", time.Now(), time.Now().Add(-24*time.Hour))
	_, err = repository.Create(ctx, cntr)
	if err != nil {
		t.Fatalf("DB record creation: %v", err)
	}

	absent := counter.Restore("not_found", "name", time.Now(), time.Now())
	err = repository.UpdateIncident(ctx, absent)

	assert.ErrorIs(t, err, ErrNotFound)
}

func TestUpdateIncident(t *testing.T) {
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
	cntr := counter.Restore("test", "name", time.Now(), time.Now().Add(-24*time.Hour))
	created, err := repository.Create(ctx, cntr)
	if err != nil {
		t.Fatalf("DB record creation: %v", err)
	}

	updated := created
	updated.Reset()
	err = repository.UpdateIncident(ctx, updated)
	if err != nil {
		t.Fatalf("Counter updating: %v", err)
	}

	found, err := repository.Find(ctx, created.ID())
	if err != nil {
		t.Fatalf("Counter finding: %v", err)
	}

	assert.Equal(t, updated.ID(), found.ID())
	assert.Equal(t, updated.Name(), found.Name())
	assert.Equal(t, updated.CreatedAt().Unix(), found.CreatedAt().Unix(), "created_at")
	assert.Equal(t, updated.LastIncidentAt().Unix(), found.LastIncidentAt().Unix(), "last_incident_at")
}

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
	assert.Equal(t, created.CreatedAt().Unix(), found.CreatedAt().Unix(), "created_at")
	assert.Equal(t, created.LastIncidentAt().Unix(), found.LastIncidentAt().Unix(), "last_incident_at")

}
