package tests

import (
	"context"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

const work = "/tmp/lifecycle"

var _ = Describe("tee", Ordered, func() {
	stdin := "Test lifecycle stdin"
	file := containerPath("create.txt")

	var server integration.Server
	var test integration.LifeCycleTest

	BeforeAll(func(ctx context.Context) {
		By("creating the lifecycle test")
		test = integration.LifeCycleTest{
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
				ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
					"exitCode":     0,
					"stdout":       stdin,
					"stderr":       "",
					"createdFiles": []string{file},
					"args": map[string]interface{}{
						"append":  false,
						"content": stdin,
						"files":   []string{file},
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					data, err := provisioner.ReadFile(context.Background(), file)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(data)).To(Equal(stdin))
				},
			},
		}

		By("creating an integration server")
		server = util.NewServer()

		By("configuring the provider")
		err := provisioner.ConfigureProvider(ctx, server)
		Expect(err).NotTo(HaveOccurred())

		By("creating a workspace in the container")
		err = provisioner.Exec(ctx, "mkdir", "-p", work)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should complete a full lifecycle", func() {
		util.TestLifecycle(server, test)
	})
})

func containerPath(name string) string {
	return path.Join(work, name)
}
