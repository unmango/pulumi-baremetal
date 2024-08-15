package lifecycle

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var _ = Describe("Chmod", func() {
	var resource tokens.Type = "baremetal:coreutils:Chmod"
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		file := containerPath("chmod.txt")

		By("creating a file to modify")
		err := provisioner.WriteFile(ctx, file, []byte("some text"))
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"files":     []string{file},
						"octalMode": "0700",
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))
				},
			},
		})

		_, err = provisioner.Exec(ctx, "touch", "blah")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should fail when file doesn't exist", func() {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"files":     []string{"/does/not/exist"},
						"octalMode": "0700",
					},
				}),
				ExpectFailure: true,
			},
		})
	})
})
