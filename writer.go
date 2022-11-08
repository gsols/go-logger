package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

// ParseWriter takes a string and returns an io.Writer
func ParseWriter(writerConfig WriterConfig) io.Writer {
	var writer io.Writer

	switch writerConfig.Writer {
	case "stderr":
		writer = os.Stderr
	case "file":
		err := os.MkdirAll(writerConfig.Directory, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		writer = io.MultiWriter(
			&lumberjack.Logger{
				Filename:   fmt.Sprintf("%s/%s", writerConfig.Directory, writerConfig.FileName),
				MaxSize:    writerConfig.MaxSize, // megabytes
				MaxBackups: writerConfig.MaxBackups,
				MaxAge:     writerConfig.MaxAge, // days
				Compress:   true,                // disabled by default
			},
			zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
				w.TimeFormat = "01/02/06 15:04:05.000 -0700"
			}),
		)
	case "discard":
		writer = io.Discard
	case "stdout":
	default:
		writer = os.Stdout
	}

	return writer
}
