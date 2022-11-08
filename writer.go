package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const TimeFormat = "2006-01-02 15:04:05.000 -0700"

type LevelWriter struct {
	Level zerolog.Level
	io.Writer
}

// WriteLevel is a custom writer that allows us to write to a file only
// if the log level is greater than or equal to the level we set.
func (lw *LevelWriter) WriteLevel(l zerolog.Level, p []byte) (n int, err error) {
	if l >= lw.Level {
		return lw.Writer.Write(p)
	}

	return len(p), nil
}

// ParseWriter takes a string and returns an io.Writer
func ParseWriter(writerConfig WriterConfig) io.Writer {
	var writer io.Writer

	switch writerConfig.Writer {
	case "discard":
		writer = io.Discard
	case "stdout":
		writer = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.Out = os.Stdout
			w.TimeFormat = TimeFormat
		})
	case "stderr":
		writer = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.Out = os.Stderr
			w.TimeFormat = TimeFormat
		})
	case "file":
		err := os.MkdirAll(writerConfig.Directory, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		fileWriter := &lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s", writerConfig.Directory, writerConfig.FileName),
			MaxSize:    writerConfig.MaxSize, // megabytes
			MaxBackups: writerConfig.MaxBackups,
			MaxAge:     writerConfig.MaxAge, // days
			Compress:   true,                // disabled by default
		}

		fileWriterLeveled := &LevelWriter{Writer: fileWriter, Level: zerolog.TraceLevel}

		consoleWriter := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.Out = os.Stdout
			w.TimeFormat = TimeFormat
		})
		writer = zerolog.MultiLevelWriter(fileWriterLeveled, consoleWriter)
	default:
		writer = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = TimeFormat
		})
	}

	return writer
}
