package util

import (
	"context"
	"fmt"
	"io"
	"time"

	tc "github.com/testcontainers/testcontainers-go"
)

type TestHost interface {
	Exec(context.Context, ...string) error
	FileExists(context.Context, string) (bool, error)
	ReadFile(context.Context, string) ([]byte, error)
	WriteFile(context.Context, string, []byte) error

	Start(context.Context) error
	Stop(context.Context) error
}

type host struct {
	ctr tc.Container
}

func (h host) Exec(ctx context.Context, args ...string) error {
	code, output, err := h.ctr.Exec(ctx, args)
	if err != nil {
		return err
	}

	out, err := io.ReadAll(output)
	if err != nil {
		return err
	}

	if code != 0 {
		return fmt.Errorf("unexpected return code: %d, output: %s", code, out)
	}

	return nil
}

// FileExists implements TestHost.
func (h *host) FileExists(ctx context.Context, path string) (bool, error) {
	return FileExists(ctx, h.ctr, path)
}

// ReadFile implements TestHost.
func (h *host) ReadFile(ctx context.Context, path string) ([]byte, error) {
	reader, err := h.ctr.CopyFileFromContainer(ctx, path)
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	return io.ReadAll(reader)
}

// WriteFile implements TestHost.
func (h *host) WriteFile(ctx context.Context, path string, data []byte) error {
	return h.ctr.CopyToContainer(ctx, data, path, 0700)
}

// Start implements TestHost.
func (h *host) Start(ctx context.Context) error {
	return h.ctr.Start(ctx)
}

// Stop implements TestHost.
func (h *host) Stop(ctx context.Context) error {
	timeout := time.Duration(10 * time.Second)
	return h.ctr.Stop(ctx, &timeout)
}

var _ = (TestHost)((*host)(nil))

func FileExists(ctx context.Context, ctr tc.Container, path string) (bool, error) {
	_, err := ctr.CopyFileFromContainer(ctx, path)
	return err == nil, nil
}
