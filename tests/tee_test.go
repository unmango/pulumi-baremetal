package tests

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

var _ = Describe("Tee", Ordered, func() {
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		By("creating a provider server")
		server = NewIntegrationProvider()

		By("configuring the provider")
		err := provisioner.ConfigureProvider(ctx, server)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a tee", func() {
		stdin := "Test stdin"
		By("generating expected data")

		By("creating the resource")
		response, err := server.Create(p.CreateRequest{
			Urn: urn("Tee", "cmd"),
			Properties: resource.PropertyMap{
				"stdin": resource.NewStringProperty(stdin),
				"create": resource.NewObjectProperty(resource.PropertyMap{
					"files": resource.NewArrayProperty([]resource.PropertyValue{
						resource.NewStringProperty("test"),
					}),
				}),
			},
			Preview: false,
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeNil())
		Expect(response.Properties["stdout"].V).To(Equal("op: OP_CREATE, cmd: COMMAND_TEE, args: []string{\"test\"}, flags: map[string]*baremetalv1alpha1.Flag(nil)"))
	})
})
