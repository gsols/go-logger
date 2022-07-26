package logger

import (
	"github.com/rs/zerolog"

	"sync"
)

// logger is a pointer to a `zerolog.Logger` struct.
var _logger *zerolog.Logger
var _once sync.Once

// Init is a function that takes a `Config` struct as an argument and returns nothing. It is a function
// that is called once and only once
func Init(config *Config) {
	_once.Do(func() {
		_logger = New(config)
	})
}

// New is a function that creates a new logger with the given configuration
func New(config *Config) *zerolog.Logger {
	var logLevel zerolog.Level = config.LogLevel

	if config.DebugMode {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	writer := ParseWriter(config.WriterConfig)

	logger := zerolog.New(writer).With().Stack().Timestamp().Caller().Logger()

	return &logger
}
