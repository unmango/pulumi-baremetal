package tests

import (
	"context"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/pulumi-baremetal/tests/util"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

var _ = Describe("Tee", Ordered, func() {
	var server integration.Server

	BeforeAll(func(ctx context.Context) {
		By("creating a working directory for the tee test")
		err := provisioner.Exec(ctx, "mkdir", "-p", "/tmp/tee")
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeEach(func(ctx context.Context) {
		By("creating a provider server")
		server = util.NewIntegrationProvider()

		By("configuring the provider")
		err := provisioner.ConfigureProvider(ctx, server)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Create", func() {
		It("should write stdin to a file", func(ctx context.Context) {
			stdin := "Test stdin"
			file := "/tmp/tee/create.txt"

			By("creating the resource")
			response, err := server.Create(p.CreateRequest{
				Urn:     util.Urn("Tee", "cmd"),
				Preview: false,
				Properties: resource.PropertyMap{
					"stdin": resource.NewStringProperty(stdin),
					"create": resource.NewObjectProperty(resource.PropertyMap{
						"files": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewStringProperty(file),
						}),
					}),
				},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(response.Properties["stderr"].V).To(BeEmpty())
			Expect(response.Properties["stdout"].V).To(Equal(stdin))

			By("attempting to copy the created file")
			reader, err := provisioner.Ctr().CopyFileFromContainer(ctx, file)
			Expect(err).NotTo(HaveOccurred())
			result, err := io.ReadAll(reader)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(result)).To(Equal(stdin))
		})
	})
})
