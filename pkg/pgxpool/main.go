package pgxpool

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(ctx context.Context, databaseDsn string) *pgxpool.Pool {
	database, err := pgxpool.New(ctx, databaseDsn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := database.Ping(ctx); err != nil {
		log.Fatalln(err.Error())
	}

	return database
}
