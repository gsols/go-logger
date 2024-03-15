package logger_test

import (
	"bytes"
	"testing"

	"github.com/gsols/go-logger"
)

func TestNew(t *testing.T) {
	out := &bytes.Buffer{}
	l := logger.New(out)
	l.Log().Msg("")
	if got, want := string(out.Bytes()), "{}\n"; got != want {
		t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
	}
}

func TestGlobal(t *testing.T) {
	logger.Debug().Msg("")
}
