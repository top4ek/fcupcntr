package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"io/fs"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Open(ctx context.Context, dsn string) (*sql.DB, error) {
	database, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	err = database.PingContext(ctx)
	if err != nil {
		closeErr := database.Close()
		return nil, errors.Join(err, closeErr)
	}

	err = runMigrations(ctx, database)
	if err != nil {
		closeErr := database.Close()
		return nil, errors.Join(err, closeErr)
	}
	return database, nil
}

func runMigrations(ctx context.Context, database *sql.DB) error {
	migrationFiles, err := fs.Sub(migrations, "migrations")
	if err != nil {
		return err
	}

	provider, err := goose.NewProvider(goose.DialectSQLite3, database, migrationFiles)
	if err != nil {
		return err
	}

	_, err = provider.Up(ctx)
	if err != nil {
		return err
	}
	return nil
}
