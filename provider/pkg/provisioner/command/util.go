package cmd

import (
	"io"
	"slices"
	"strings"
)

func stdinReader(stdin *string) io.Reader {
	if stdin == nil {
		return nil
	}

	return strings.NewReader(*stdin)
}

func prepend[T any](x T, xs []T) []T {
	return slices.Insert(xs, 0, x)
}
