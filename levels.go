package logger

import (
	"fmt"

	"github.com/rs/zerolog"
)

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return _logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return _logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return _logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return _logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return _logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return _logger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return _logger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return _logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	_logger.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	_logger.Printf(format, v...)
}

// ParseLevel converts a level string into a zerolog Level value.
// returns an error if the input string does not match known values.
func ParseLevel(levelStr string) zerolog.Level {
	logLevel, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	return logLevel
}

// ParseVerboseLevel takes an integer defining the
// verbosity level and returns a zerolog.Level
func ParseVerboseLevel(verbose int) zerolog.Level {
	switch verbose {
	case 1:
		return zerolog.InfoLevel
	case 2:
		return zerolog.DebugLevel
	case 3:
		return zerolog.TraceLevel
	default:
		return zerolog.WarnLevel
	}
}
