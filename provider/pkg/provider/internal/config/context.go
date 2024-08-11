package config

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider"
)

func FromContext(ctx context.Context) provider.Config {
	return infer.GetConfig[provider.Config](ctx)
}
