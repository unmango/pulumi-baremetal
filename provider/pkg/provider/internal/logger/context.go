package logger

import (
	"context"

	provider "github.com/pulumi/pulumi-go-provider"
)

func FromContext(ctx context.Context) provider.Logger {
	return provider.GetLogger(ctx)
}
