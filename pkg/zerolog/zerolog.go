package zerolog

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

func NewStdout(isDebug bool) zerolog.Logger {
	var stdout io.Writer = os.Stdout

	if isDebug {
		stdout = zerolog.ConsoleWriter{
			Out: stdout,
		}
	}

	return zerolog.New(stdout).With().Timestamp().Logger()
}
