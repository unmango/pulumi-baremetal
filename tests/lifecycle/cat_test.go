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

var _ = Describe("Cat", func() {
	var resource tokens.Type = "baremetal:coreutils:Cat"
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		expectedText := "Test text ja"
		file := containerPath("cat.txt")

		By("creating a file to modify")
		err := provisioner.WriteFile(ctx, file, []byte(expectedText))
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
					Expect(output["stdout"]).To(HavePropertyValue(expectedText))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))
				},
			},
		})

		Expect(err).NotTo(HaveOccurred())
	})

	It("should fail when file doesn't exist", func() {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"files": []string{"/does/not/exist"},
					},
				}),
				ExpectFailure: true,
			},
		})
	})
})
