package logger

import (
	"github.com/rs/zerolog"
)

// Option is a function that takes a pointer to a Logger and sets its options.
type Option func(logger *zerolog.Logger)

// WithDebug is a function that takes a boolean value and returns an option.
func WithDebug() Option {
	return func(logger *zerolog.Logger) {
		*logger = logger.Level(zerolog.DebugLevel)
	}
}

// WithVerbosity is a function that takes an integer value and returns an option.
func WithVerbosity(verbosity int) Option {
	return func(logger *zerolog.Logger) {
		*logger = logger.Level(zerolog.Level(int(zerolog.PanicLevel) - verbosity))
	}
}
