package logger

import (
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type SentryHook struct{}

func (h SentryHook) Run(_ *zerolog.Event, level zerolog.Level, message string) {
	if level == zerolog.ErrorLevel || level == zerolog.FatalLevel || level == zerolog.PanicLevel {
		captured := &sentry.Event{
			Level:   sentry.Level(level.String()),
			Message: message,
		}
		sentry.CaptureEvent(captured)
	}
}
