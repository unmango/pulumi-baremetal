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

var _ = Describe("Mv", func() {
	var resource tokens.Type = "baremetal:coreutils:Mv"
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		file := containerPath("mv.txt")
		firstFile := containerPath("mv-new.txt")
		secondFile := containerPath("mv-2.txt")

		By("creating a file to be moved")
		err := provisioner.WriteFile(ctx, file, []byte("some text"))
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"source":      []string{file},
						"destination": firstFile,
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
						file: firstFile,
					})))
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(ctx, file))
					Expect(provisioner).To(ContainFile(ctx, firstFile))
				},
			},
			Updates: []integration.Operation{{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"source":      []string{file},
						"destination": secondFile,
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
						file: secondFile,
					})))
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(ctx, file))
					Expect(provisioner).NotTo(ContainFile(ctx, firstFile))
					Expect(provisioner).To(ContainFile(ctx, secondFile))
				},
			}},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, secondFile))
		Expect(provisioner).NotTo(ContainFile(ctx, firstFile))
		Expect(provisioner).To(ContainFile(ctx, file))
	})

	It("should support custom updates", func(ctx context.Context) {
		source := containerPath("mv-custom1.txt")
		dest := containerPath("mv-custom2.txt")
		final := containerPath("mv-custom-final.txt")

		By("creating a file to be moved")
		err := provisioner.WriteFile(ctx, source, []byte("some text"))
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"source":      []string{source},
						"destination": dest,
					},
					"customUpdate": []string{"mv", dest, final},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
						source: dest,
					})))
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(ctx, source))
					Expect(provisioner).To(ContainFile(ctx, dest))
				},
			},
			Updates: []integration.Operation{{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"source":      []string{"this is kinda nonsensical so I might change how this works in the future"},
						"destination": dest,
					},
					"customUpdate": []string{"mv", dest, final},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(ctx, source))
					Expect(provisioner).NotTo(ContainFile(ctx, dest))
					Expect(provisioner).To(ContainFile(ctx, final))
				},
			}},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, source))
		Expect(provisioner).NotTo(ContainFile(ctx, dest))
		Expect(provisioner).To(ContainFile(ctx, final))
	})

	It("should support custom deletes", func(ctx context.Context) {
		source := containerPath("mv-custom1.txt")
		dest := containerPath("mv-custom2.txt")
		final := containerPath("mv-custom-final.txt")

		By("creating a file to be moved")
		err := provisioner.WriteFile(ctx, source, []byte("some text"))
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"source":      []string{source},
						"destination": dest,
					},
					"customDelete": []string{"mv", dest, final},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
						source: dest,
					})))
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(ctx, source))
					Expect(provisioner).To(ContainFile(ctx, dest))
				},
			},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, source))
		Expect(provisioner).NotTo(ContainFile(ctx, dest))
		Expect(provisioner).To(ContainFile(ctx, final))
	})

	It("should delete when file doesn't exist", func(ctx context.Context) {
		source := containerPath("mv-custom1.txt")
		dest := containerPath("mv-custom2.txt")

		By("creating a file to be moved")
		err := provisioner.WriteFile(ctx, source, []byte("some text"))
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"source":      []string{source},
						"destination": dest,
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
						source: dest,
					})))
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(ctx, source))
					Expect(provisioner).To(ContainFile(ctx, dest))
					_, err := provisioner.Exec(ctx, "rm", dest)
					Expect(err).NotTo(HaveOccurred())
					Expect(provisioner).NotTo(ContainFile(ctx, dest))
				},
			},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, source))
		Expect(provisioner).NotTo(ContainFile(ctx, dest))
	})

	It("should fail when source doesn't exist", func() {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"source": []string{"/does/not/exist"},
					},
				}),
				ExpectFailure: true,
			},
		})
	})
})
