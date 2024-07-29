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
			"stdin": resource.NewStringProperty(stdin),
			"create": resource.NewObjectProperty(resource.PropertyMap{
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

			e, ok := response.Properties["stderr"].V.(string)
			Expect(ok).To(BeTrueBecause("stderr was not a string"))
			stderr = &e

			o, ok := response.Properties["stdout"].V.(string)
			Expect(ok).To(BeTrueBecause("stdout was not a string"))
			stdout = &o

			By("attempting to copy the created file")
			reader, err := provisioner.Ctr().CopyFileFromContainer(ctx, file)
			Expect(err).NotTo(HaveOccurred())
			result, err := io.ReadAll(reader)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(result)).To(Equal(stdin))
		})

		It("should delete", func(ctx context.Context) {
			Expect(teeId).NotTo(BeNil())
			err := server.Delete(p.DeleteRequest{
				Urn: urn,
				ID:  *teeId,
				Properties: resource.PropertyMap{
					"stdin": resource.NewStringProperty(stdin),
					"create": resource.NewObjectProperty(resource.PropertyMap{
						"files": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewStringProperty(file),
						}),
					}),
					"stdout": resource.NewStringProperty(*stdout),
					"stderr": resource.NewStringProperty(*stderr),
					"created_files": resource.NewArrayProperty([]resource.PropertyValue{
						resource.NewStringProperty(file),
					}),
				},
			})

			Expect(err).NotTo(HaveOccurred())

			By("attempting to copy the created file")
			_, err = provisioner.Ctr().CopyFileFromContainer(ctx, file)
			Expect(err).To(MatchError(ContainSubstring("Could not find the file /tmp/tee/create.txt")))
		})
	})
})
