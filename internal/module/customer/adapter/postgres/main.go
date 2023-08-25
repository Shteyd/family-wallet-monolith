package postgres

import (
	"context"
	"monolith/internal/domain"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgresAdapter interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type _PostgresAdapter struct {
	Database domain.Database
}

func NewPostgresAdapter(database domain.Database) PostgresAdapter {
	return &_PostgresAdapter{
		Database: database,
	}
}

func (adapter *_PostgresAdapter) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return adapter.Database.Exec(ctx, sql, args...)
}

func (adapter *_PostgresAdapter) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return adapter.Database.QueryRow(ctx, sql, args...)
}
