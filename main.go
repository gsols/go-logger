package logger

import (
	"io"

	"github.com/rs/zerolog"
)

var logger = New(WithConsoleWriter()).With().Caller().Timestamp().Logger()

func New(w io.Writer, opts ...Option) zerolog.Logger {
	zerolog.TimeFieldFormat = TimeFieldFormat
	zerolog.CallerMarshalFunc = CallerMarshalFunc
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	l := zerolog.New(w)
	for _, opt := range opts {
		opt(&l)
	}

	return l
}

func WithOptions(opts ...Option) {
	for _, opt := range opts {
		opt(&logger)
	}
}
