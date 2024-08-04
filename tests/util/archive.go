package util

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
)

func CreateTarArchive(w io.Writer, members map[string]string) error {
	g := gzip.NewWriter(w)
	defer func() {
		_ = g.Close()
	}()

	t := tar.NewWriter(g)
	defer func() {
		_ = t.Close()
	}()

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
