package logging

import (
	"context"
	"io"
	"log/slog"
	"time"

	"github.com/awryme/sniproxy/pkg/extslices"
)

var backgroundCtx = context.Background()

type Logf func(msg string, attrs ...slog.Attr)

func NewLogf(output io.Writer) Logf {
	handler := slog.NewTextHandler(output, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if len(groups) == 0 && a.Key == slog.LevelKey {
				return slog.Attr{}
			}
			return a
		},
	})

	return func(msg string, attrs ...slog.Attr) {
		level := slog.LevelInfo

		lastValue, ok := extslices.LastElem(attrs)
		if ok && lastValue.Key == keyError {
			level = slog.LevelError
		}

		r := slog.NewRecord(time.Now(), level, msg, 0)
		r.AddAttrs(attrs...)
		handler.Handle(backgroundCtx, r)
	}
}

func LogError(err error, logf Logf, msg string, attrs ...slog.Attr) {
	if err == nil {
		return
	}
	attrs = append(attrs, slog.String(keyError, err.Error()))
	logf(msg, attrs...)
}

func With(logf Logf, attrs ...slog.Attr) Logf {
	return func(msg string, origAttrs ...slog.Attr) {
		logf(msg, append(attrs, origAttrs...)...)
	}
}
