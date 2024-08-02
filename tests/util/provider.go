package util

import (
	"context"
	"strings"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"golang.org/x/exp/maps"

	baremetal "github.com/unmango/pulumi-baremetal/provider"
)

type ProviderBuilder interface {
	Configure() error
	WithProvisioner(address, port string) ProviderBuilder
	WithCerts(*HostCerts) ProviderBuilder
}

func NewServer() integration.Server {
	return integration.NewServer(
		baremetal.Name,
		semver.MustParse("1.0.0"),
		baremetal.Provider(),
	)
}

func NewServerWithContext(ctx context.Context) integration.Server {
	return integration.NewServerWithContext(ctx,
		baremetal.Name,
		semver.MustParse("1.0.0"),
		baremetal.Provider(),
	)
}

func Urn(typ string, mods ...string) resource.URN {
	if len(mods) == 0 {
		mods = []string{"index"}
	}

	tok := strings.Join(append(mods, typ), ":")
	return resource.NewURN("stack", "proj", "",
		tokens.Type("test:"+tok), "name")
}

func ConfigureProvider(server integration.Server) ProviderBuilder {
	return &configureBuilder{p.ConfigureRequest{}, server}
}

type configureBuilder struct {
	p.ConfigureRequest
	server integration.Server
}

// Configure implements ProviderBuilder.
func (c *configureBuilder) Configure() error {
	return c.server.Configure(c.ConfigureRequest)
}

// WithCerts implements ProviderBuilder.
func (c *configureBuilder) WithCerts(certs *HostCerts) ProviderBuilder {
	args := c.Args.Mappable()
	maps.Copy(args, map[string]interface{}{
		"caPath":   certs.CaPath,
		"certPath": certs.CertPath,
		"keyPath":  certs.KeyPath,
	})

	c.Args = resource.NewPropertyMapFromMap(args)

	return c
}

// WithProvisioner implements ProviderBuilder.
func (c *configureBuilder) WithProvisioner(address string, port string) ProviderBuilder {
	args := c.Args.Mappable()
	maps.Copy(args, map[string]interface{}{
		"address": address,
		"port":    port,
	})

	c.Args = resource.NewPropertyMapFromMap(args)

	return c
}

var _ = (ProviderBuilder)((*configureBuilder)(nil))
