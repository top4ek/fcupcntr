-- +goose Up
-- +goose StatementBegin
CREATE TABLE counters(
  id CHAR(36) PRIMARY KEY NOT NULL,
  created_at INTEGER NOT NULL,
  last_incident_at INTEGER NOT NULL,
  name VARCHAR(100) NOT NULL
);
-- +goose StatementEnd
