package main

import (
	"flag"
	"monolith/config"
	migrate "monolith/migrations"
	"monolith/pkg/slog"
)

const (
	actionUp   = "up"
	actionDown = "down"
)

var (
	actionFlag = flag.String("action", "up", "Choose action: up / down")
	helpFlag   = flag.Bool("help", false, "View default help message")
)

func main() {
	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
	}

	config := config.New(".")

	switch *actionFlag {
	case actionUp:
		migrate.Up(config.DatabaseDsn)
	case actionDown:
		migrate.Down(config.DatabaseDsn)
	default:
		slog.Fatal(
			"Broken action type",
			slog.String("action", *actionFlag),
		)
	}
}
