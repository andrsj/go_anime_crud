package zerolog

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andrsj/go_anime_crud/pkg/logger"
	"github.com/rs/zerolog"
)

type ZerologLogger struct {
	logger *zerolog.Logger
}

func New() logger.Interface {
	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    true,
		TimeFormat: time.RFC1123,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
	}

	logger := zerolog.New(output).
		With().
		Timestamp().
		Logger()

	return &ZerologLogger{logger: &logger}
}

func (l *ZerologLogger) Debug(message string) {
	l.logger.Debug().Msg(message)
}

func (l *ZerologLogger) Info(message string) {
	l.logger.Info().Msg(message)
}

func (l *ZerologLogger) Warn(message string) {
	l.logger.Warn().Msg(message)
}

func (l *ZerologLogger) Error(message string) {
	l.logger.Error().Msg(message)
}

func (l *ZerologLogger) Fatal(message string) {
	l.logger.Fatal().Msg(message)
}
