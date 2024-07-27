package tests

import (
	"io"

	tc "github.com/testcontainers/testcontainers-go"
)

type writerConsumer struct {
	io.Writer
}

func (w *writerConsumer) Accept(log tc.Log) {
	w.Write(log.Content)
}

func LogToWriter(w io.Writer) tc.LogConsumer {
	return &writerConsumer{Writer: w}
}
