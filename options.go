package logger

import (
	"io"

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

// WithOutput is a function that takes an io.Writer and returns an option.
func WithOutput(w io.Writer) Option {
	return func(l *zerolog.Logger) {
		*l = l.Output(w)
	}
}

// WithSentry is a function that returns an option.
func WithSentry() Option {
	var sentryHook SentryHook
	return func(l *zerolog.Logger) {
		*l = l.Hook(sentryHook)
	}
}
