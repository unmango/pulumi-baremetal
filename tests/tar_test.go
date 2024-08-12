package tests

import (
	"bytes"
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

var _ = Describe("Tar", func() {
	var resource tokens.Type = "baremetal:coreutils:Tar"
	var server integration.Server
	work := containerPath("tar")

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		fileName := "someFile.txt"
		contents := "Some text that really doesn't matter"
		archive := containerPath("tar", "test-archive.tar.gz")
		dest := containerPath("tar", "destination")
		expectedFile := containerPath("tar", "destination", fileName)

		By("ensuring container directories exist")
		_, err := provisioner.Exec(ctx, "mkdir", "-p", work, dest)
		Expect(err).NotTo(HaveOccurred())

		By("creating an archive to operate on")
		buf := &bytes.Buffer{}
		err = util.CreateTarArchive(buf, map[string]string{
			fileName: contents,
		})
		Expect(err).NotTo(HaveOccurred())

		By("writing the archive to the container")
		err = provisioner.WriteFile(ctx, archive, buf.Bytes())
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"extract":   true,
						"file":      archive,
						"directory": dest,
						"args":      []string{fileName},
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["createdFiles"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
						pr.NewProperty(expectedFile),
					})))
					Expect(output["movedFiles"].V).To(Equal(pr.PropertyMap{}))
					Expect(output["args"].V).To(Equal(inputs["args"].V))
					Expect(provisioner).To(ContainFile(context.Background(), expectedFile))
				},
			},
		})

		Expect(provisioner).To(ContainFile(ctx, archive))
		Expect(provisioner).NotTo(ContainFile(ctx, expectedFile))
	})
})
