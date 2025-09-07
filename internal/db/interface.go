package db

import (
	"context"
	"database/sql"
)

// IDatabase defines the database interface.
type IDatabase interface {
	GetDB() *sql.DB
	WithTransaction(function func() error) error
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) *sql.Row
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	PingContext(ctx context.Context) error
	Close() error
}
