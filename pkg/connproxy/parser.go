package connproxy

import (
	"io"

	"github.com/awryme/sniproxy/pkg/logging"
)

type HeaderParser interface {
	ParseHeader(logf logging.Logf, reader io.Reader) (hostname string, receivedData []byte, err error)
}
