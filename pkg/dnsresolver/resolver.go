package dnsresolver

import (
	"context"
	"fmt"
	"net"
	"time"
)

const dnsPort = 53

const DefaultResolver = "1.1.1.1"

func New(dnsaddr string) *net.Resolver {
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Second * 15,
			}
			return d.DialContext(ctx, network, fmt.Sprintf("%s:%d", dnsaddr, dnsPort))
		},
	}
}
