package baremetal

import (
	"bytes"
	"encoding/gob"
)

func Marshal[T any](x T) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(x); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func UnMarshal[T any](data []byte, x T) error {
	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)

	return dec.Decode(x)
}
