package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const TimeFormat = "2006-01-02 15:04:05.000 -0700"

type Writer int8

const (
	DiscardWriter Writer = iota
	StdoutWriter
	StderrWriter
	FileWriter
	DefaultWriter // Default is a combination of File and Stdout with a global level of zerolog.TraceLevel
)

// BuildConsoleWriter returns a new zerolog.ConsoleWriter with
// the given output writer and a custom time format
func BuildConsoleWriter(out io.Writer) io.Writer {
	return zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = out
		w.TimeFormat = TimeFormat
	})
}

// BuildFileWriter creates a directory if it doesn't exist,
// then returns a lumberjack.Logger that writes to a file in that directory
func BuildFileWriter(config *WriterConfig) io.Writer {
	err := os.MkdirAll(config.Directory, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", config.Directory, config.FileName),
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge, // days
		Compress:   true,          // disabled by default
	}
}

// GetWriter takes a string and returns an io.Writer
func GetWriter(config *WriterConfig) io.Writer {
	switch config.Writer {
	case DiscardWriter:
		return io.Discard
	case StdoutWriter:
		return BuildConsoleWriter(os.Stdout)
	case StderrWriter:
		return BuildConsoleWriter(os.Stderr)
	case FileWriter:
		return BuildFileWriter(config)
	case DefaultWriter:
		writer := zerolog.MultiLevelWriter(
			BuildFileWriter(config),
			&LevelWriter{
				Level:  zerolog.GlobalLevel(),
				Writer: BuildConsoleWriter(os.Stdout),
			})
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		return writer
	}

	return io.Discard
}

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
