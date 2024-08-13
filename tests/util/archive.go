package util

import (
	"archive/tar"
	"fmt"
)

func WriteTarContents(t *tar.Writer, members map[string]string) error {
	for k, v := range members {
		if err := t.WriteHeader(&tar.Header{Name: k, Size: int64(len(v))}); err != nil {
			return fmt.Errorf("writing header for %s: %w", k, err)
		}

		if _, err := t.Write([]byte(v)); err != nil {
			return fmt.Errorf("writing contents for %s: %w", k, err)
		}
	}

	return nil
}
