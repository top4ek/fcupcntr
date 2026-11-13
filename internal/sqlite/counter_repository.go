package sqlite

import (
	"context"
	"database/sql"

	"fcupcntr/internal/counter"

	"github.com/google/uuid"
)

type CounterRepository struct {
	database *sql.DB
}

func NewCounterRepository(database *sql.DB) *CounterRepository {
	return &CounterRepository{
		database: database,
	}
}

func (r *CounterRepository) Create(ctx context.Context, cntr counter.Counter) (counter.Counter, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return cntr, err
	}

	createdAt := cntr.CreatedAt().UnixNano()
	lastIncidentAt := cntr.LastIncidentAt().UnixNano()

	query := "INSERT INTO counters (id, name, created_at, last_incident_at) VALUES (?, ?, ?, ?)"

	_, err = r.database.ExecContext(ctx, query, id.String(), cntr.Name(), createdAt, lastIncidentAt)
	if err != nil {
		return cntr, err
	}

	return counter.Restore(
		id.String(),
		cntr.Name(),
		cntr.CreatedAt(),
		cntr.LastIncidentAt(),
	), nil
}
