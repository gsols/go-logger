package logger

import (
	"github.com/rs/zerolog"
)

const defaultLogLevel = zerolog.WarnLevel

var logger = zerolog.New(WithConsoleWriter()).Level(defaultLogLevel).With().Caller().Timestamp().Logger()

func init() {
	zerolog.CallerMarshalFunc = CallerMarshalFunc
}

func WithOptions(opts ...Option) {
	for _, opt := range opts {
		opt(&logger)
	}
}
