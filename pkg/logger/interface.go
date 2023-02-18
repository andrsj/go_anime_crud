package logger

type Interface interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}
