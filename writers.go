package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

const FileNameTimeFormat = "2006-01-02_15_04_05.log"

func WithConsoleWriter() io.Writer {
	return zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
		w.Out = os.Stdout
	})
}

func WithFileWriter(path string) io.Writer {
	_ = os.MkdirAll(path, os.ModePerm)
	path = filepath.Join(path, time.Now().Format(FileNameTimeFormat))
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	return file
}

func MultiLevelWriter(writers ...io.Writer) io.Writer {
	return zerolog.MultiLevelWriter(writers...)
}
