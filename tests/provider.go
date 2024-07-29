package tests

import (
	"strings"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	baremetal "github.com/unmango/pulumi-baremetal/provider"
)

func NewIntegrationProvider() integration.Server {
	return integration.NewServer(
		baremetal.Name,
		semver.MustParse("1.0.0"),
		baremetal.Provider(),
	)
}

func urn(typ string, mods ...string) resource.URN {
	if len(mods) == 0 {
		mods = []string{"index"}
	}

	tok := strings.Join(append(mods, typ), ":")
	return resource.NewURN("stack", "proj", "",
		tokens.Type("test:"+tok), "name")
}
