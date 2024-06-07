package networking

import (
	"fmt"
	"net"
)

func DialAnyTCP(addrs []string, port int) (net.Conn, error) {
	var lastError error
	for _, addr := range addrs {
		clientConn, err := net.Dial("tcp", fmt.Sprintf(
			"%s:%d", addr, port,
		))
		if err == nil {
			return clientConn, nil
		}
		lastError = fmt.Errorf("dial addr %s: %w (port = %d)", addr, err, port)
	}
	return nil, lastError
}
