package main

import (
	"log/slog"
	"runtime/debug"

	"github.com/awryme/slogf"
)

func printBuildInfo(logf slogf.Logf) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		logf("no build info available")
		return
	}
	attrs := []slog.Attr{slog.String("go_version", info.GoVersion)}
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs", "vcs.revision":
			attrs = append(attrs, slog.String(setting.Key, setting.Value))
		}
	}
	logf("build info", attrs...)
}
