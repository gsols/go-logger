package logger

import (
	"io"
	"sync"

	"github.com/rs/zerolog"
)

var mux = sync.RWMutex{}

// Option is a function that takes a pointer to a Logger and sets its options.
// This is used to apply different configurations to the logger.
type Option func(logger *zerolog.Logger)

// WithVerbosity is a function that takes an integer value and returns an option.
// This option sets the logger's level based on the provided verbosity.
// The verbosity is subtracted from the Panic level to determine the log level.
func WithVerbosity(verbosity int) Option {
	return func(l *zerolog.Logger) {
		mux.Lock()
		defer mux.Unlock()
		*l = l.Level(DefaultLevel - zerolog.Level(verbosity))
	}
}

// WithMaxPathParents is a function that takes an integer value and returns an Option.
// This option sets the maximum number of parent directories that will be included in the log's path.
//
// This function is useful when you want to limit the verbosity of the file paths in your logs.
// For example, if you set max to 2, only the last two directories in the path will be included.
func WithMaxPathParents(max int) Option {
	return func(_ *zerolog.Logger) {
		mux.Lock()
		defer mux.Unlock()
		MaxPathNumber = max
	}
}

// WithSentry is a function that returns an option.
// This option adds a Sentry hook to the logger.
// The Sentry hook is used to send logs to a Sentry server for error tracking and monitoring.
func WithSentry() Option {
	var sentryHook SentryHook
	return func(l *zerolog.Logger) {
		mux.Lock()
		defer mux.Unlock()
		*l = l.Hook(sentryHook)
	}
}

// WithWriter is a function that takes an io.Writer and returns an option.
// This option sets the logger's output destination to the provided io.Writer.
// This is similar to WithOutput, but can be used when you want to specify a different writer.
func WithWriter(w io.Writer) Option {
	return func(l *zerolog.Logger) {
		mux.Lock()
		defer mux.Unlock()
		*l = l.Output(w)
	}
}
