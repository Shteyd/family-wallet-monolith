package pgxpool

import (
	"context"
	"monolith/pkg/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Aliases
type (
	CommandTag   pgconn.CommandTag
	Rows         = pgx.Rows
	Row          = pgx.Row
	Batch        = pgx.Batch
	BatchResults = pgx.BatchResults
)

func NewDatabase(ctx context.Context, databaseDsn string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(databaseDsn)
	if err != nil {
		slog.Fatal(err.Error())
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		slog.Fatal(err.Error())
	}

	return pool
}
