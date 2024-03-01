package logger

import (
	"github.com/rs/zerolog"
)

// Option is a function that takes a pointer to a Logger and sets its options.
type Option func(logger *zerolog.Logger)

// WithDebug is a function that takes a boolean value and returns an option.
func WithDebug() Option {
	return func(l *zerolog.Logger) {
		*l = l.Level(zerolog.DebugLevel)
	}
}

// WithVerbosity is a function that takes an integer value and returns an option.
func WithVerbosity(verbosity int) Option {
	return func(l *zerolog.Logger) {
		*l = l.Level(zerolog.Level(int(zerolog.PanicLevel) - verbosity))
	}
}

func WithSentry() Option {
	var sentryHook SentryHook
	return func(l *zerolog.Logger) {
		*l = l.Hook(sentryHook)
	}
}
