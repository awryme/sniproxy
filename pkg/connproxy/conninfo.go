package connproxy

import (
	"log/slog"

	"github.com/awryme/sniproxy/pkg/defaultvalue"
)

type ConnInfo struct {
	Hostname string
	Addr     string
}

func (info ConnInfo) LogValue() slog.Value {
	const emptyInfo = "<empty>"
	hostname := defaultvalue.For(info.Hostname, emptyInfo)
	addr := defaultvalue.For(info.Addr, emptyInfo)

	return slog.GroupValue(slog.String("hostname", hostname), slog.String("addr", addr))
}
