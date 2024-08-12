package tests

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/asset"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var _ = Describe("Tee", func() {
	var resource tokens.Type = "baremetal:coreutils:Tee"
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle with content", Pending, func(ctx context.Context) {
		file := containerPath("create.txt")
		newFile := containerPath("update.txt")

		By("creating a string asset")
		content, err := asset.FromText("Test lifecycle content")
		Expect(err).NotTo(HaveOccurred())

		By("creating a different string asset")
		newContent, err := asset.FromText("Updated content")
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"content": content,
						"files":   []string{file},
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(content.Text))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))

					data, err := provisioner.ReadFile(ctx, file)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(data)).To(Equal(content.Text))
				},
			},
			Updates: []integration.Operation{
				{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"content": content,
							"files":   []string{newFile},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(content.Text))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(ctx, file))

						data, err := provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(content.Text))
					},
				},
				{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"content": newContent,
							"files":   []string{newFile},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(newContent.Text))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))

						data, err := provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newContent.Text))
					},
				},
				{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"content": newContent,
							"files":   []string{file, newFile},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(newContent.Text))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(HaveLen(2)) // TODO: Make this better
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))

						data, err := provisioner.ReadFile(ctx, file)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newContent.Text))

						data, err = provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newContent.Text))
					},
				},
			},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, newFile))
		Expect(provisioner).NotTo(ContainFile(ctx, file))
	})

	It("should complete a full lifecycle with stdin", func(ctx context.Context) {
		file := containerPath("create.txt")
		newFile := containerPath("update.txt")
		stdin := "Test content stdin"
		newStdin := "Updated stdin"

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"stdin": stdin,
						"files": []string{file},
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(stdin))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))

					data, err := provisioner.ReadFile(ctx, file)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(data)).To(Equal(stdin))
				},
			},
			Updates: []integration.Operation{
				{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"stdin": stdin,
							"files": []string{newFile},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(stdin))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(ctx, file))

						data, err := provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(stdin))
					},
				},
				{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"stdin": newStdin,
							"files": []string{newFile},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(newStdin))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))

						data, err := provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))
					},
				},
				{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"stdin": newStdin,
							"files": []string{file, newFile},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(newStdin))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(HaveLen(2)) // TODO: Make this better
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))

						data, err := provisioner.ReadFile(ctx, file)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))

						data, err = provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))
					},
				},
			},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, newFile))
		Expect(provisioner).NotTo(ContainFile(ctx, file))
	})

	It("should fail when file doesn't exist", func() {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"stdin": "does not matter for this test",
						"files": []string{"/does/not/exist"},
					},
				}),
				ExpectFailure: true,
			},
		})
	})
})
