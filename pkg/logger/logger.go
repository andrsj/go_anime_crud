package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const skipFrameCount = 3

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *zerolog.Logger
}

// New -.
func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	logger := zerolog.
		New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			FormatLevel: func(i interface{}) string {
				return strings.ToUpper(fmt.Sprintf("[%s]", i))
			},
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s", i)
			},
			// FormatCaller: func(i interface{}) string {
			// 	return filepath.Base(fmt.Sprintf("%s", i))
			// },
		}).
		With().
		Timestamp().
		// Caller(). // ?
		// CallerWithSkipFrameCount(-3). // runtime/extern.go func runtime.Caller
		// CallerWithSkipFrameCount(-2). // zerolog/event.go func hook.Run -> caller > runtime.Caller
		// CallerWithSkipFrameCount(-1). // zerolog/event.go func hook.Run -> caller
		// CallerWithSkipFrameCount(0). // zerolog/event.go func hook.Run
		// CallerWithSkipFrameCount(1). // zerolog/event.go func Msg
		// CallerWithSkipFrameCount(2). // pkg/logger/logger.go func log
		// CallerWithSkipFrameCount(3). // pkg/logger/logger.go func Info
		// CallerWithSkipFrameCount(4). // internal/app/app.go
		// CallerWithSkipFrameCount(5). // actual our caller => zerolog.CallerSkipFrameCount + skipFrameCount
		// CallerWithSkipFrameCount(6). // proc.go
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).
		Logger()

	return &Logger{
		logger: &logger,
	}
}

var _ Interface = (*Logger)(nil)

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.msg("error", message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)

	os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
