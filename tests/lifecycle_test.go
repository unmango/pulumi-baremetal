package tests

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

const work = "/tmp/lifecycle"

type lifecycleTest func(
	*testing.T,
	context.Context,
	util.TestProvisioner,
) integration.LifeCycleTest

func TestLifecycle(t *testing.T) {
	ctx := context.Background()
	prov, err := util.NewProvisioner("5000", os.Stdout)
	if err != nil {
		t.Fatalf("failed to create provisioner: %s", err)
	}

	if err = prov.Start(ctx); err != nil {
		t.Fatalf("failed to start provisioner: %s", err)
	}

	if err = prov.Exec(ctx, "mkdir", "-p", work); err != nil {
		t.Fatalf("failed to create workspace in container: %s", err)
	}

	suite := map[string]lifecycleTest{
		"tee": TeeTest,
	}

	for name, createTest := range suite {
		_ = t.Run(name, func(t *testing.T) {
			test := createTest(t, ctx, prov)
			server := util.NewServerWithContext(ctx)
			prov.ConfigureProvider(ctx, server)
			test.Run(t, server)
		})
	}
}

func TeeTest(t *testing.T, ctx context.Context, p util.TestProvisioner) integration.LifeCycleTest {
	stdin := "Test lifecycle stdin"
	file := containerPath("create.txt")

	return integration.LifeCycleTest{
		Resource: "baremetal:cmd:Tee",
		Create: integration.Operation{
			Inputs: resource.PropertyMap{
				"create": resource.NewObjectProperty(resource.PropertyMap{
					"content": resource.NewStringProperty(stdin),
					"files": resource.NewArrayProperty([]resource.PropertyValue{
						resource.NewStringProperty(file),
					}),
				}),
			},
			Hook: func(inputs, output resource.PropertyMap) {
				exists, err := p.FileExists(ctx, file)
				if err != nil {
					t.Fatalf("failed to check file's existence: %s", err)
				}

				if !exists {
					t.Fatalf("expected %s to exist", file)
				}
			},
		},
	}
}

func containerPath(name string) string {
	return path.Join(work, name)
}
