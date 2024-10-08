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

var _ = Describe("Command", func() {
	var resource tokens.Type = "baremetal:index:Command"
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		file := containerPath("chmod.txt")
		expectedFile := "/tmp/exec-test.txt"

		By("creating a file to modify")
		err := provisioner.WriteFile(ctx, file, []byte("some text"))
		Expect(err).NotTo(HaveOccurred())

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"create": []string{"mv", file, expectedFile},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).To(ContainFile(ctx, expectedFile))
				},
			},
		})

		_, err = provisioner.Exec(ctx, "touch", "blah")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should execute whitelisted command", func(ctx context.Context) {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"create": []string{"perl", "--help"},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"].V).NotTo(BeEmpty())
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["args"]).To(Equal(inputs["args"]))
				},
			},
		})
	})

	It("should refuse to execute unknown bin", func() {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"create": []string{"really-hope-this-never-exists"},
				}),
				ExpectFailure: true,
			},
		})
	})

	It("should execute custom delete", func(ctx context.Context) {
		expectedFile := "/tmp/custom-delete-test.txt"

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"create": []string{"touch", expectedFile},
					"delete": []string{"rm", expectedFile},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["create"]).To(Equal(inputs["create"]))
					Expect(output["delete"]).To(Equal(inputs["delete"]))
				},
			},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, expectedFile))
	})

	It("should execute custom update", func(ctx context.Context) {
		expectedFile := "/tmp/custom-update-test.txt"

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"create":   []string{"touch", expectedFile},
					"update":   []string{"rm", expectedFile},
					"triggers": []string{"a"},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["create"]).To(Equal(inputs["create"]))
					Expect(output["delete"]).To(Equal(inputs["delete"]))
				},
			},
			Updates: []integration.Operation{{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"create":   []string{"touch", expectedFile},
					"update":   []string{"rm", expectedFile},
					"triggers": []string{"b"},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["create"]).To(Equal(inputs["create"]))
					Expect(output["delete"]).To(Equal(inputs["delete"]))
				},
			}},
		})

		Expect(provisioner).NotTo(ContainFile(ctx, expectedFile))
	})
})
