package logger

import "github.com/rs/zerolog"

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return _logger.With()
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return _logger.Hook(h)
}

// GroupID returns a pointer to a zerolog logger with the `groupId` field set to the value of the `id` parameter
func GroupID(id string) *zerolog.Logger {
	l := With().Str("groupId", id).Logger()
	return &l
}
