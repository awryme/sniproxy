package bufkeeper

import (
	"bytes"
	"io"
)

// BufKeeper takes a reader and keeps every read data in a buffer, accessible later
type BufKeeper struct {
	buf    *bytes.Buffer
	reader io.Reader
}

func New(reader io.Reader) *BufKeeper {
	return &BufKeeper{bytes.NewBuffer(nil), reader}
}

func (bk *BufKeeper) Read(p []byte) (n int, err error) {
	n, err = bk.reader.Read(p)
	bk.buf.Write(p[:n])
	return n, err
}

func (bk *BufKeeper) Bytes() []byte {
	return bk.buf.Bytes()
}

func (bk *BufKeeper) Buffer() *bytes.Buffer {
	return bk.buf
}

func (bk *BufKeeper) Len() int {
	return bk.buf.Len()
}
