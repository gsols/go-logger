package logger

import (
	"io"

	"github.com/rs/zerolog"
)

var logger = New(WithDefaultWriter()).With().Caller().Timestamp().Logger()

func New(w io.Writer) zerolog.Logger {
	return zerolog.New(w)
}
