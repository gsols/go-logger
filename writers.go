package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

const TimeFormat = "2006-01-02 15:04:05"
const FileNameTimeFormat = "logs/2006-01-02_15.log"

func WithConsoleWriter() io.Writer {
	return zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = TimeFormat
		w.Out = os.Stdout
	})
}

func WithFileWriter() io.Writer {
	_ = os.MkdirAll("logs", os.ModePerm)
	file, _ := os.OpenFile(time.Now().Format(FileNameTimeFormat), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	return file
}

func WithDefaultWriter() io.Writer {
	return zerolog.MultiLevelWriter(WithConsoleWriter(), WithFileWriter())
}
