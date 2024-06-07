package logging

import (
	"io"
	"log/slog"
	"time"
)

type Logger struct {
	handler slog.Handler
}

func NewLogger(output io.Writer) Logger {
	handler := slog.NewTextHandler(output, &slog.HandlerOptions{})
	return Logger{handler}
}

func (l Logger) Info(msg string, attrs ...slog.Attr) {
	r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, 0)
	r.AddAttrs(attrs...)
	l.handler.Handle(backgroundCtx, r)
}

func (l Logger) Error(err error, msg string, attrs ...slog.Attr) {
	if err == nil {
		return
	}
	attrs = append(attrs, slog.String("error", err.Error()))
	r := slog.NewRecord(time.Now(), slog.LevelError, msg, 0)
	r.AddAttrs(attrs...)
	l.handler.Handle(backgroundCtx, r)
}

func (l Logger) With(attrs ...slog.Attr) Logger {
	return Logger{l.handler.WithAttrs(attrs)}
}

func (l Logger) ToLogf() Logf {
	return func(msg string, attrs ...slog.Attr) {
		l.Info(msg, attrs...)
	}
}
