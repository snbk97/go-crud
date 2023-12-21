package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type LogContext struct {
	LogInfo func(message string, rest ...string)
	LogErr  func(err error, message string, rest ...string)
}

func CreateLogger(logContext string) (context *LogContext) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	_log := zerolog.New(os.Stdout).With().Timestamp().Str("logContext", logContext).Logger()

	context = &LogContext{
		LogInfo: func(message string, rest ...string) {
			m := []string{message}
			if len(rest) > 0 {
				m = append(m, rest...)
			}
			_log.Info().Msg(strings.Join(m, ":"))
		},
		LogErr: func(err error, message string, rest ...string) {
			m := []string{message}
			if len(rest) > 0 {
				m = append(m, rest...)
			}
			_log.Error().Caller().Stack().Err(err).Msg(strings.Join(m, ":"))
		},
	}

	return
}
