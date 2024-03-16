package logger

import (
	"github.com/rs/zerolog"
)

const DefaultLevel = zerolog.WarnLevel
const DebugLevel = zerolog.DebugLevel

var logger = zerolog.New(WithConsoleWriter()).Level(DefaultLevel).With().Caller().Timestamp().Logger()

func init() {
	zerolog.CallerMarshalFunc = CallerMarshalFunc
}

func WithOptions(opts ...Option) {
	for _, opt := range opts {
		opt(&logger)
	}
}
