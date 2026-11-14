package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"fcupcntr/internal/counter"

	"github.com/google/uuid"
)

type CounterRepository struct {
	database *sql.DB
}

var ErrNotFound = errors.New("not found")

func NewCounterRepository(database *sql.DB) *CounterRepository {
	return &CounterRepository{
		database: database,
	}
}

func (r *CounterRepository) UpdateIncident(ctx context.Context, cntr counter.Counter) error {
	query := "UPDATE counters SET last_incident_at = ? WHERE id = ?"

	lastIncidentAt := cntr.LastIncidentAt().Unix()
	result, err := r.database.ExecContext(ctx, query, lastIncidentAt, cntr.ID())
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *CounterRepository) Find(ctx context.Context, id string) (counter.Counter, error) {
	query := "SELECT id, name, created_at, last_incident_at FROM counters WHERE id = ?"

	var (
		uuid           string
		name           string
		createdAt      int64
		lastIncidentAt int64
	)
	err := r.database.QueryRowContext(ctx, query, id).Scan(&uuid, &name, &createdAt, &lastIncidentAt)
	if errors.Is(err, sql.ErrNoRows) {
		return counter.Counter{}, ErrNotFound
	}

	if err != nil {
		return counter.Counter{}, err
	}

	return counter.Restore(
		uuid,
		name,
		time.Unix(createdAt, 0),
		time.Unix(lastIncidentAt, 0),
	), nil
}

func (r *CounterRepository) Create(ctx context.Context, cntr counter.Counter) (counter.Counter, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return cntr, err
	}

	createdAt := cntr.CreatedAt().Unix()
	lastIncidentAt := cntr.LastIncidentAt().Unix()

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
