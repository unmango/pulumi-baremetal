package tests

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/pulumi-baremetal/tests/util"
	. "github.com/unmango/pulumi-baremetal/tests/util/expect"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

var _ = Describe("Tee", Ordered, func() {
	urn := util.Urn("Tee", "cmd")
	var server integration.Server

	BeforeAll(func(ctx context.Context) {
		By("creating a working directory for the tee test")
		err := provisioner.Exec(ctx, "mkdir", "-p", "/tmp/tee")
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeAll(func(ctx context.Context) {
		By("creating a provider server")
		server = util.NewIntegrationProvider()
	})

	BeforeAll(func(ctx context.Context) {
		By("configuring the provider")
		err := provisioner.ConfigureProvider(ctx, server)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("write stdin to a file", func() {
		stdin := "Test stdin"
		file := "/tmp/tee/create.txt"

		var teeId *string
		var stdout *string
		var stderr *string

		props := resource.PropertyMap{
			"create": resource.NewObjectProperty(resource.PropertyMap{
				"content": resource.NewStringProperty(stdin),
				"files": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewStringProperty(file),
				}),
			}),
		}

		It("should create", func(ctx context.Context) {
			response, err := server.Create(p.CreateRequest{
				Urn:        urn,
				Preview:    false,
				Properties: props,
			})

			Expect(err).NotTo(HaveOccurred())
			teeId = &response.ID

			out, ok := response.Properties["stderr"].V.(string)
			Expect(ok).To(BeTrueBecause("stderr was not a string"))
			stderr = &out

			out, ok = response.Properties["stdout"].V.(string)
			Expect(ok).To(BeTrueBecause("stdout was not a string"))
			stdout = &out

			Expect(provisioner).To(ContainFile(ctx, file))
		})

		It("should delete", func(ctx context.Context) {
			By("asserting the developer hasn't made an error")
			Expect(teeId).NotTo(BeNil())

			err := server.Delete(p.DeleteRequest{
				Urn: urn,
				ID:  *teeId,
				Properties: resource.PropertyMap{
					"createOpts": resource.NewObjectProperty(resource.PropertyMap{
						"content": resource.NewStringProperty(stdin),
						"files": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewStringProperty(file),
						}),
					}),
					"stdout": resource.NewStringProperty(*stdout),
					"stderr": resource.NewStringProperty(*stderr),
					"createdFiles": resource.NewArrayProperty([]resource.PropertyValue{
						resource.NewStringProperty(file),
					}),
				},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(provisioner).NotTo(ContainFile(ctx, file))
		})
	})
})
