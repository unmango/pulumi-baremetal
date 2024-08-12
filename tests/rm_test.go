package tests

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var _ = Describe("Rm", func() {
	var resource tokens.Type = "baremetal:coreutils:Rm"
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		file := containerPath("rm.txt")

		By("creating a file to be removed")
		err := provisioner.WriteFile(ctx, file, []byte("some text"))
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"files": []string{file},
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"].V).To(BeEquivalentTo(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(context.Background(), file))
				},
			},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, file))
	})
})
