package migrate

import (
	"database/sql"
	"embed"
	"monolith/pkg/slog"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	defaultDriver = "postgres"
)

//go:embed *.sql
var embedMigrations embed.FS

var database *sql.DB

func setup(databaseDsn string) {
	var err error
	database, err = sql.Open(defaultDriver, databaseDsn)
	if err != nil {
		slog.Fatal(err.Error(), slog.String("driver-name", defaultDriver))
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(defaultDriver); err != nil {
		slog.Fatal(err.Error())
	}
}

func Up(databaseDsn string) {
	if database == nil {
		setup(databaseDsn)
	}

	if err := goose.Up(database, "."); err != nil {
		slog.Fatal(err.Error())
	}
}

func Down(databaseDsn string) {
	if database == nil {
		setup(databaseDsn)
	}

	if err := goose.Down(database, "."); err != nil {
		slog.Fatal(err.Error())
	}
}
