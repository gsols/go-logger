package logger_test

import (
	"github.com/idbi/logger"

	"testing"
	"time"
)

var fakeMessage = "Test logging, but use a somewhat realistic message length."

func TestLogInit(t *testing.T) {
	logger.Init()
}

func TestLogInitWithConfig(t *testing.T) {
	logger.InitWithConfig(&logger.Config{
		Verbose:  false,
		LogLevel: logger.ParseLevel("warn"),
		WriterConfig: logger.WriterConfig{
			Writer:     "discard",
			Directory:  ".",
			FileName:   "test.log",
			MaxSize:    1,
			MaxBackups: 2,
			MaxAge:     24,
		}})
}

func BenchmarkLogEmpty(b *testing.B) {
	logger.InitWithConfig(&logger.Config{
		Verbose:  false,
		LogLevel: logger.ParseLevel("warn"),
		WriterConfig: logger.WriterConfig{
			Writer:     "discard",
			Directory:  ".",
			FileName:   "test.log",
			MaxSize:    1,
			MaxBackups: 2,
			MaxAge:     24,
		}})
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Log().Msg("")
		}
	})
}

func BenchmarkDisabled(b *testing.B) {
	logger.InitWithConfig(&logger.Config{
		Verbose:  false,
		LogLevel: logger.ParseLevel("disabled"),
		WriterConfig: logger.WriterConfig{
			Writer:     "discard",
			Directory:  ".",
			FileName:   "test.log",
			MaxSize:    1,
			MaxBackups: 2,
			MaxAge:     24,
		}})
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg(fakeMessage)
		}
	})
}

func BenchmarkInfo(b *testing.B) {
	logger.InitWithConfig(&logger.Config{
		Verbose:  false,
		LogLevel: logger.ParseLevel("warn"),
		WriterConfig: logger.WriterConfig{
			Writer:     "discard",
			Directory:  ".",
			FileName:   "test.log",
			MaxSize:    1,
			MaxBackups: 2,
			MaxAge:     24,
		}})
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg(fakeMessage)
		}
	})
}

func BenchmarkContextFields(b *testing.B) {
	logger.InitWithConfig(&logger.Config{
		Verbose:  false,
		LogLevel: logger.ParseLevel("warn"),
		WriterConfig: logger.WriterConfig{
			Writer:     "discard",
			Directory:  ".",
			FileName:   "test.log",
			MaxSize:    1,
			MaxBackups: 2,
			MaxAge:     24,
		}})
	logger.With().
		Str("string", "four!").
		Time("time", time.Time{}).
		Int("int", 123).
		Float32("float", -2.203230293249593).
		Logger()
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg(fakeMessage)
		}
	})
}

func BenchmarkContextAppend(b *testing.B) {
	logger.InitWithConfig(&logger.Config{
		Verbose:  false,
		LogLevel: logger.ParseLevel("warn"),
		WriterConfig: logger.WriterConfig{
			Writer:     "discard",
			Directory:  ".",
			FileName:   "test.log",
			MaxSize:    1,
			MaxBackups: 2,
			MaxAge:     24,
		}})
	logger.With().
		Str("foo", "bar").
		Logger()
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.With().Str("bar", "baz")
		}
	})
}

func BenchmarkLogFields(b *testing.B) {
	logger.InitWithConfig(&logger.Config{
		Verbose:  false,
		LogLevel: logger.ParseLevel("warn"),
		WriterConfig: logger.WriterConfig{
			Writer:     "discard",
			Directory:  ".",
			FileName:   "test.log",
			MaxSize:    1,
			MaxBackups: 2,
			MaxAge:     24,
		}})
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("string", "four!").
				Time("time", time.Time{}).
				Int("int", 123).
				Float32("float", -2.203230293249593).
				Msg(fakeMessage)
		}
	})
}
