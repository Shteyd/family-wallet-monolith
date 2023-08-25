package pgxpool

import (
	"context"
	"monolith/pkg/slog"

	"github.com/jackc/pgx/v5/pgxpool"
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
