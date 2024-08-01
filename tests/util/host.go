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
	Ip(context.Context) (string, error)
	ReadFile(context.Context, string) ([]byte, error)
	WriteFile(context.Context, string, []byte) error

	Start(context.Context) error
	Stop(context.Context) error
}

type host struct {
	req tc.GenericContainerRequest
	ctr *tc.Container
}

func (h host) Exec(ctx context.Context, args ...string) error {
	ctr, err := h.ensureContainer(ctx)
	if err != nil {
		return err
	}

	code, output, err := ctr.Exec(ctx, args)
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
	ctr, err := h.ensureContainer(ctx)
	if err != nil {
		return false, err
	}

	return FileExists(ctx, ctr, path)
}

// ReadFile implements TestHost.
func (h *host) ReadFile(ctx context.Context, path string) ([]byte, error) {
	ctr, err := h.ensureContainer(ctx)
	if err != nil {
		return nil, err
	}

	reader, err := ctr.CopyFileFromContainer(ctx, path)
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	return io.ReadAll(reader)
}

// WriteFile implements TestHost.
func (h *host) WriteFile(ctx context.Context, path string, data []byte) error {
	ctr, err := h.ensureContainer(ctx)
	if err != nil {
		return err
	}

	return ctr.CopyToContainer(ctx, data, path, 0700)
}

func (h *host) Ip(ctx context.Context) (string, error) {
	ctr, err := h.ensureContainer(ctx)
	if err != nil {
		return "", err
	}

	return ctr.ContainerIP(ctx)
}

// Start implements TestHost.
func (h *host) Start(ctx context.Context) error {
	ctr, err := h.ensureContainer(ctx)
	if err != nil {
		return err
	}

	return ctr.Start(ctx)
}

// Stop implements TestHost.
func (h *host) Stop(ctx context.Context) error {
	ctr, err := h.ensureContainer(ctx)
	if err != nil {
		return err
	}

	timeout := time.Duration(10 * time.Second)
	return ctr.Stop(ctx, &timeout)
}

func (h *host) ensureContainer(ctx context.Context) (tc.Container, error) {
	if h.ctr == nil {
		ctr, err := tc.GenericContainer(ctx, h.req)
		if err != nil {
			return nil, fmt.Errorf("creating container: %w", err)
		}

		h.ctr = &ctr
	}

	return *h.ctr, nil
}

var _ = (TestHost)((*host)(nil))

func FileExists(ctx context.Context, ctr tc.Container, path string) (bool, error) {
	_, err := ctr.CopyFileFromContainer(ctx, path)
	return err == nil, nil
}
