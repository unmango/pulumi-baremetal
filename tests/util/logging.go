package util

import (
	"fmt"
	"io"

	tc "github.com/testcontainers/testcontainers-go"
)

type writerConsumer struct {
	io.Writer
}

func LogToWriter(w io.Writer) tc.LogConsumer {
	return &writerConsumer{Writer: w}
}

func NewLogger(w io.Writer) tc.Logging {
	return &writerConsumer{Writer: w}
}

// Accept implements testcontainers.LogConsumer.
func (w *writerConsumer) Accept(log tc.Log) {
	_, _ = w.Write(log.Content)
}

// Printf implements testcontainers.Logging.
func (w *writerConsumer) Printf(format string, v ...interface{}) {
	_, _ = fmt.Fprintf(w, format, v...)
}
