// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

var _ = internal.GetEnvOrDefault

func GetAddress(ctx *pulumi.Context) string {
	return config.Get(ctx, "baremetal:address")
}
func GetCaPem(ctx *pulumi.Context) string {
	return config.Get(ctx, "baremetal:caPem")
}
func GetCertPem(ctx *pulumi.Context) string {
	return config.Get(ctx, "baremetal:certPem")
}
func GetKeyPem(ctx *pulumi.Context) string {
	return config.Get(ctx, "baremetal:keyPem")
}
func GetPort(ctx *pulumi.Context) string {
	return config.Get(ctx, "baremetal:port")
}
