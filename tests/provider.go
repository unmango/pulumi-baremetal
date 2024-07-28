package tests

import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi-go-provider/integration"

	baremetal "github.com/unmango/pulumi-baremetal/provider"
)

func NewIntegrationProvider() integration.Server {
	return integration.NewServer(
		baremetal.Name,
		semver.MustParse("1.0.0"),
		baremetal.Provider(),
	)
}
