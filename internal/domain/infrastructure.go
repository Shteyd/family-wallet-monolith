package domain

type (
	LoggerArgs map[string]any

	Logger interface {
		Debug(msg string, args LoggerArgs)
		Error(err error, args LoggerArgs)
		Warn(msg string, args LoggerArgs)
		Info(msg string, args LoggerArgs)
	}
)
