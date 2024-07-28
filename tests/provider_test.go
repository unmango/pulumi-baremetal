package tests

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var _ = Describe("Provider", Ordered, func() {})

func urn(typ string, mods ...string) resource.URN {
	if len(mods) == 0 {
		mods = []string{"index"}
	}

	tok := strings.Join(append(mods, typ), ":")
	return resource.NewURN("stack", "proj", "",
		tokens.Type("test:"+tok), "name")
}
