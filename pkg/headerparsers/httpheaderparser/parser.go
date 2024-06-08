package httpheaderparser

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"github.com/awryme/slogf"
	"github.com/awryme/sniproxy/pkg/bufkeeper"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ParseHeader(logf slogf.Logf, reader io.Reader) (hostname string, receivedData []byte, _ error) {
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
