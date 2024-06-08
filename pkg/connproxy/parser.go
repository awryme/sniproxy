package connproxy

import (
	"io"

	"github.com/awryme/slogf"
)

type HeaderParser interface {
	ParseHeader(logf slogf.Logf, reader io.Reader) (hostname string, receivedData []byte, err error)
}
