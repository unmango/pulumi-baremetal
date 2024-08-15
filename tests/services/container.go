package services

import (
	"context"

	tc "github.com/testcontainers/testcontainers-go"
)

func FileExists(ctx context.Context, ctr tc.Container, path string) (bool, error) {
	exitCode, _, err := ctr.Exec(ctx, []string{"stat", path})
	return err == nil && exitCode == 0, nil
}
