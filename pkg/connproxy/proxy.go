package connproxy

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/awryme/sniproxy/pkg/dnsresolver"
	"github.com/awryme/sniproxy/pkg/logging"
	"github.com/awryme/sniproxy/pkg/networking"
)

type Proxy struct {
	resolver     *net.Resolver
	headerParser HeaderParser
	port         int
}

func New(headerParser HeaderParser, port int) *Proxy {
	return &Proxy{
		resolver:     dnsresolver.New(dnsresolver.DefaultResolver),
		headerParser: headerParser,
		port:         port,
	}
}

// HandleConn parses tls message, and tunnels conn based on SNI
//
// HandleConn doesn't close the connection
func (p *Proxy) HandleConn(ctx context.Context, logf logging.Logf, localConn net.Conn) (info ConnInfo, _ error) {
	hostname, receivedData, err := p.headerParser.ParseHeader(logf, localConn)
	if err != nil {
		return info, fmt.Errorf("parse request header: %w", err)
	}
	info.Hostname = hostname

	logf("parsed tls header", slog.String("sni_hostname", hostname))

	addrs, err := p.resolver.LookupHost(ctx, hostname)
	if err != nil {
		return info, fmt.Errorf("resolve remote hostname %s: %w", hostname, err)
	}

	remoteConn, err := networking.DialAnyTCP(addrs, p.port)
	if err != nil {
		return info, fmt.Errorf("dial remote hostname %s: %w", hostname, err)
	}
	defer remoteConn.Close()

	info.Addr = remoteConn.RemoteAddr().String()

	_, err = remoteConn.Write(receivedData)
	if err != nil {
		return info, fmt.Errorf("write to remote hostname %s: %w", hostname, err)
	}
	err = twoWayCopy(remoteConn, localConn, func(elapsed time.Duration) {
		logf("connection ongoing",
			info.ToSlog("conn_info"),
			slog.Duration("duration", elapsed))
	})
	if err != nil {
		return info, fmt.Errorf("two-way copy with remote hostname %s: %w", hostname, err)
	}
	return info, nil
}
