package httpheaderparser

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"github.com/awryme/sniproxy/pkg/bufkeeper"
	"github.com/awryme/sniproxy/pkg/logging"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ParseHeader(logf logging.Logf, reader io.Reader) (hostname string, receivedData []byte, _ error) {
	readBuf := bufkeeper.New(reader)

	req, err := http.ReadRequest(bufio.NewReader(readBuf))
	if err != nil {
		return "", readBuf.Bytes(), fmt.Errorf("error reading http request: %w", err)
	}
	if req.Host == "" {
		return "", readBuf.Bytes(), fmt.Errorf("received http request, but host is empty")
	}

	return req.Host, readBuf.Bytes(), nil
}
