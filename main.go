package logger

import (
	"sync"

	"github.com/rs/zerolog"
)

// logger is a pointer to a `zerolog.Logger` struct.
var _logger *zerolog.Logger
var _once sync.Once

// Init creates a new logger with the default configuration, and sets it as the default logger
func Init() {
	_once.Do(func() {
		_logger = New()
	})
}

// InitWithConfig creates a new logger with the given configuration, and sets it as the default logger
func InitWithConfig(config *Config) {
	_once.Do(func() {
		_logger = NewWithConfig(config)
	})
}

// New is a function that creates a new logger with the default configuration
func New() *zerolog.Logger {
	zerolog.SetGlobalLevel(ParseVerboseLevel(0))

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	writer := GetWriter(&WriterConfig{
		Writer: StdoutWriter,
	})

	logger := zerolog.New(writer).With().Stack().Timestamp().Caller().Logger()

	return &logger
}

// NewWithConfig is a function that creates a new logger with the given configuration
func NewWithConfig(config *Config) *zerolog.Logger {
	logLevel := ParseVerboseLevel(config.Verbosity)

	if config.Debug {
		logLevel = zerolog.TraceLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	writer := GetWriter(config.WriterConfig)

	logger := zerolog.New(writer).With().Stack().Timestamp().Caller().Logger()

	return &logger
}
