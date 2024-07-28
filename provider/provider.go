package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "baremetal"

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			LanguageMap: map[string]any{
				"csharp": map[string]string{"rootNamespace": "UnMango"},
				"go":     map[string]string{"importBasePath": "github.com/unmango/pulumi-baremetal/sdk/go/baremetal"},
				"nodejs": map[string]string{"packageName": "@unmango/baremetal"},
				"python": map[string]string{"packageName": "unmango_baremetal"},
			},
		},
		Resources: []infer.InferredResource{
			infer.Resource[cmd.Tee](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Config: infer.Config[provider.Config](),
	})
}
