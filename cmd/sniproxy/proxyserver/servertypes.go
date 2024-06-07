package proxyserver

import (
	"github.com/awryme/sniproxy/pkg/connproxy"
	"github.com/awryme/sniproxy/pkg/headerparsers/httpheaderparser"
	"github.com/awryme/sniproxy/pkg/headerparsers/tlsheaderparser"
)

type ProxyType string

func (pt ProxyType) String() string {
	return string(pt)
}

const (
	ProxyTypeHTTP  = "http"
	ProxyTypeHTTPS = "https"
)

func mapProxyTypeToHeaderParser(proxyType ProxyType) (connproxy.HeaderParser, bool) {
	switch proxyType {
	case ProxyTypeHTTP:
		return httpheaderparser.New(), true
	case ProxyTypeHTTPS:
		return tlsheaderparser.New(), true
	default:
		return nil, false
	}
}
