package zap

import (
	"log"

	"go.uber.org/zap"
)

func New(isDebug bool) *zap.Logger {
	options := []zap.Option{
		zap.AddCallerSkip(2),
	}

	logger, err := zap.NewProduction(options...)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if isDebug {
		logger, err = zap.NewDevelopment(options...)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	return logger
}

func NewSugar(isDebug bool) *zap.SugaredLogger {
	return New(isDebug).Sugar()
}
