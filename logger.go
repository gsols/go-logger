package logger

import (
	"context"
	"fmt"
	"io"

	"github.com/rs/zerolog"
)

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return logger.Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return logger.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return logger.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	return logger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	return logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return logger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return logger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// Ctx returns the zerolog.Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

// WithContext returns a copy of ctx with the receiver attached. The Logger
// attached to the provided Context (if any) will not be effected.  If the
// receiver's log level is Disabled it will only be attached to the returned
// Context if the provided Context has a previously attached Logger. If the
// provided Context has no attached Logger, a Disabled Logger will not be
// attached.
func WithContext(ctx context.Context) context.Context {
	return logger.WithContext(ctx)
}

// UpdateContext updates the internal logger's context.
//
// Caution: This method is not concurrency safe.
// Use the With method to create a child logger before modifying the context from concurrent goroutines.
func UpdateContext(update func(c zerolog.Context) zerolog.Context) {
	logger.UpdateContext(update)
}

// GetLevel is a function that returns the current logging level of the logger.
// The logging level is of type zerolog.Level, which is an integer type where
// higher values represent lower severity levels.
//
// This function is useful when you want to check the current logging level
// programmatically, for example to only execute certain code if the logging
// level is set to a certain severity.
func GetLevel() zerolog.Level {
	return logger.GetLevel()
}

// SetGlobalLevel is a function that sets the global logging level.
// The logging level is of type zerolog.Level, which is an integer type where
// higher values represent lower severity levels. The levels are as follows:
// 0=panic, 1=fatal, 2=error, 3=warn, 4=info, 5=debug, 6=trace.
//
// This function is useful when you want to set the global logging level
// programmatically, for example to control the overall verbosity of your application.
//
// Parameters:
//
//	l (zerolog.Level): The logging level to set as the global level.
func SetGlobalLevel(l zerolog.Level) {
	zerolog.SetGlobalLevel(l)
}

// GlobalLevel is a function that returns the current global logging level.
// The logging level is of type zerolog.Level, which is an integer type where
// higher values represent lower severity levels.
//
// This function is useful when you want to check the current global logging level
// programmatically, for example to only execute certain code if the global logging
// level is set to a certain severity.
func GlobalLevel() zerolog.Level {
	return zerolog.GlobalLevel()
}
