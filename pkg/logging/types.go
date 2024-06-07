package logging

import "log/slog"

const keyError = "error"

func Error(err error) slog.Attr {
	if err == nil {
		return slog.Attr{}
	}
	return slog.String(keyError, err.Error())
}
