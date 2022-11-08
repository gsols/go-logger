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
	logLevel := zerolog.InfoLevel
	zerolog.SetGlobalLevel(logLevel)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	writer := ParseWriter(WriterConfig{
		Writer: "stdout",
	})

	logger := zerolog.New(writer).With().Stack().Timestamp().Caller().Logger()

	return &logger
}

// NewWithConfig is a function that creates a new logger with the given configuration
func NewWithConfig(config *Config) *zerolog.Logger {
	logLevel := config.LogLevel

	if config.DebugMode {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	writer := ParseWriter(config.WriterConfig)

	logger := zerolog.New(writer).With().Stack().Timestamp().Caller().Logger()

	return &logger
}
