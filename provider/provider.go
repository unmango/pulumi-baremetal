package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/coreutils"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/kubeadm"
)

const Name string = "baremetal"

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			PluginDownloadURL: "github://api.github.com/unmango",
			LanguageMap: map[string]any{
				"csharp": map[string]any{
					"rootNamespace": "UnMango",
					"packageReferences": map[string]any{
						"Pulumi":         "[3.65.0.0,4)",
						"Pulumi.Command": "1.0.*",
					},
				},
				"go": map[string]any{
					"importBasePath": "github.com/unmango/pulumi-baremetal/sdk/go/baremetal",
					"generics":       "side-by-side",
				},
				"nodejs": map[string]any{
					"packageName": "@unmango/baremetal",
					"dependencies": map[string]any{
						"@pulumi/command": "^1.0.0",
					},
				},
				"python": map[string]any{
					"packageName": "unmango_baremetal",
					"requires": map[string]any{
						"pulumi-command": ">=1.0.0,<2.0.0",
					},
				},
			},
		},
		Config: infer.Config[provider.Config](),
		Resources: []infer.InferredResource{
			infer.Resource[coreutils.Mktemp](),
			infer.Resource[coreutils.Tar](),
			infer.Resource[coreutils.Wget](),
			infer.Resource[coreutils.Chmod](),
			infer.Resource[coreutils.Mkdir](),
			infer.Resource[coreutils.Mv](),
			infer.Resource[coreutils.Rm](),
			infer.Resource[coreutils.Tee](),
			infer.Resource[kubeadm.Kubeadm](),
		},
		Components: []infer.InferredComponent{
			// Consuming external resources is no bueno atm
			// infer.Component[*provider.Bootstrap](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
	})
}
