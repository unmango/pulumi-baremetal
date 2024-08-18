package config

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/config"
)

func FromContext(ctx context.Context) config.Config {
	return infer.GetConfig[config.Config](ctx)
}
