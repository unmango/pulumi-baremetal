package tests

import (
	"context"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var _ = Describe("Wget", Ordered, func() {
	var resource tokens.Type = "baremetal:coreutils:Wget"
	var server integration.Server
	dir := containerPath("wget")

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		url := "https://raw.githubusercontent.com/unmango/pulumi-baremetal/main/README.md"
		file := path.Join(dir, "README.md")

		By("creating a workspace for wget in the container")
		_, err := provisioner.Exec(ctx, "mkdir", "-p", dir)
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"directoryPrefix": dir,
						"urls":            []string{url},
						"quiet":           true,
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["createdFiles"].V).To(ContainElement(pr.NewProperty(file)))
					Expect(output["movedFiles"].V).To(BeEmpty())

					args := output["args"].ObjectValue()
					Expect(args["directoryPrefix"]).To(HavePropertyValue(dir))
					Expect(args["urls"].V).To(ContainElement(pr.NewProperty(url)))
					Expect(args["quiet"]).To(HavePropertyValue(true))

					_, err := provisioner.ReadFile(ctx, file)
					Expect(err).NotTo(HaveOccurred())
				},
			},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, file))
	})
})
