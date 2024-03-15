package logger

import (
	"io"

	"github.com/rs/zerolog"
)

// Option is a function that takes a pointer to a Logger and sets its options.
// This is used to apply different configurations to the logger.
type Option func(logger *zerolog.Logger)

// WithVerbosity is a function that takes an integer value and returns an option.
// This option sets the logger's level based on the provided verbosity.
// The verbosity is subtracted from the Panic level to determine the log level.
func WithVerbosity(verbosity int) Option {
	return func(l *zerolog.Logger) {
		*l = l.Level(zerolog.Level(int(zerolog.WarnLevel) - verbosity))
	}
}

// WithOutput is a function that takes an io.Writer and returns an option.
// This option sets the logger's output destination to the provided io.Writer.
func WithOutput(w io.Writer) Option {
	return func(l *zerolog.Logger) {
		*l = l.Output(w)
	}
}

// WithSentry is a function that returns an option.
// This option adds a Sentry hook to the logger.
// The Sentry hook is used to send logs to a Sentry server for error tracking and monitoring.
func WithSentry() Option {
	var sentryHook SentryHook
	return func(l *zerolog.Logger) {
		*l = l.Hook(sentryHook)
	}
}

// WithWriter is a function that takes an io.Writer and returns an option.
// This option sets the logger's output destination to the provided io.Writer.
// This is similar to WithOutput, but can be used when you want to specify a different writer.
func WithWriter(w io.Writer) Option {
	return func(l *zerolog.Logger) {
		*l = l.Output(w)
	}
}
