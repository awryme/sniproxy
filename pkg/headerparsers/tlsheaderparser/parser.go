package tlsheaderparser

import (
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"

	"github.com/awryme/slogf"
	"github.com/paultag/sniff/parser"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ParseHeader(logf slogf.Logf, reader io.Reader) (hostname string, receivedData []byte, _ error) {
	headerBuf := make([]byte, tlsHeaderFullLen)
	headerLen, err := reader.Read(headerBuf)
	if err != nil {
		return "", nil, fmt.Errorf("read tls header: %w", err)
	}

	if headerLen != tlsHeaderFullLen {
		return "", nil, fmt.Errorf("received wrong tls header length: smaller than required (received = %d, required = %d)", headerLen, tlsHeaderFullLen)
	}

	headerType := int(headerBuf[0])
	if headerType != headerTypeHandshake {
		return "", nil, fmt.Errorf("received wrong tls header type (received = %d, required = %d)", headerType, headerTypeHandshake)
	}

	headerMessageLength := int(binary.BigEndian.Uint16(headerBuf[3:5]))

	logf("received tls request header", slog.Int("header_message_length", int(headerMessageLength)))

	messageBuf := make([]byte, headerMessageLength)
	// messageLength, err := reader.Read(messageBuf)
	// if err != nil {
	// 	return "", nil, fmt.Errorf("read tls header message: %w", err)
	// }

	messageLength, err := io.ReadAtLeast(reader, messageBuf, len(messageBuf))
	if err != nil {
		return "", nil, fmt.Errorf("read tls header message: %w", err)
	}

	if messageLength != headerMessageLength {
		logf("received incomplete tls header message", slog.Int("received", messageLength), slog.Int("expected", headerMessageLength))
	}

	messageBuf = messageBuf[:messageLength]
	fullBuff := append(headerBuf, messageBuf...)

	hostname, err = parser.GetHostname(fullBuff)
	if err != nil {
		return hostname, messageBuf, fmt.Errorf("parse sni from request: %w", err)
	}

	return hostname, fullBuff, nil
}
