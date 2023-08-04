package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Error(error, Args)
	Warn(string, Args)
	Info(string, Args)
	Debug(string, Args)
}

type logger struct {
	Stdout *zap.SugaredLogger
}

func New(zap *zap.SugaredLogger) Logger {
	return logger{Stdout: zap}
}

func (log logger) Debug(msg string, args Args) {
	log.Stdout.Debugw(msg, args.ParseArgs()...)
}

func (log logger) Error(err error, args Args) {
	log.Stdout.Errorw(err.Error(), args.ParseArgs()...)
}

func (log logger) Info(msg string, args Args) {
	log.Stdout.Infow(msg, args.ParseArgs()...)
}

func (log logger) Warn(msg string, args Args) {
	log.Stdout.Warnw(msg, args.ParseArgs()...)
}
