package proxyserver

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/awryme/sniproxy/pkg/connproxy"
	"github.com/awryme/sniproxy/pkg/logging"
	"github.com/oklog/ulid/v2"
)

func Start(ctx context.Context, logf logging.Logf, addr string, listenPort int, proxyType ProxyType) error {
	headerParser, ok := mapProxyTypeToHeaderParser(proxyType)
	if !ok {
		return fmt.Errorf("failed to find conn header parser by proxy type (proxy_type = %s)", proxyType)
	}
	listenAddr := fmt.Sprintf("%s:%d", addr, listenPort)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return fmt.Errorf("start listener %w", err)
	}

	proxy := connproxy.New(headerParser, listenPort)
	logf = logging.With(logf, slog.String("server_type", proxyType.String()))
	logf("started new server", slog.String("addr", addr), slog.Int("port", listenPort))

	for {
		conn, err := listener.Accept()
		if err != nil {
			logf("listener accept error",
				slog.String("addr", addr),
				slog.Int("port", listenPort),
				logging.Error(err),
			)
			continue
		}

		go handleConn(ctx, logf, proxy, conn)
	}
}

func handleConn(ctx context.Context, logf logging.Logf, proxy *connproxy.Proxy, conn net.Conn) {
	defer conn.Close()

	start := time.Now()
	connId := ulid.Make().String()

	logf = logging.With(logf, slog.String("conn_id", connId))

	logf("accepted connection, handling", slog.String("raddr", conn.RemoteAddr().String()))
	err := proxy.HandleConn(ctx, logf, conn)
	success := err == nil
	logf("connection finished",
		slog.Bool("success", success),
		logging.Error(err),
		slog.Duration("elapsed", time.Since(start)),
	)
}
