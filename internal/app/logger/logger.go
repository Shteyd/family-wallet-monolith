package logger

import (
	"monolith/internal/domain"

	"github.com/rs/zerolog"
)

type _Logger struct {
	Zerolog zerolog.Logger
}

func New(zerolog zerolog.Logger) domain.Logger {
	return _Logger{
		Zerolog: zerolog,
	}
}

func (logger _Logger) Debug(msg string, args domain.LoggerArgs) {
	logger.Zerolog.Debug().Fields(args).Msg(msg)
}

func (logger _Logger) Error(err error, args domain.LoggerArgs) {
	logger.Zerolog.Err(err).Fields(args).Send()
}

func (logger _Logger) Info(msg string, args domain.LoggerArgs) {
	logger.Zerolog.Info().Fields(args).Msg(msg)
}

func (logger _Logger) Warn(msg string, args domain.LoggerArgs) {
	logger.Zerolog.Warn().Fields(args).Msg(msg)
}
