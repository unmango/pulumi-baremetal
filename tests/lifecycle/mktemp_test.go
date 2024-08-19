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

var _ = Describe("Mktemp", func() {
	var resource tokens.Type = "baremetal:coreutils:Mktemp"
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"tmpdir": true,
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"].V).NotTo(BeEmpty())
					Expect(output["exitCode"].V).To(BeEquivalentTo(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
				},
			},
			Updates: []integration.Operation{
				{ // Add a trigger
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"tmpdir": true,
						},
						"triggers": []string{"a trigger"},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"].V).NotTo(BeEmpty())
						Expect(output["exitCode"].V).To(BeEquivalentTo(0))
						Expect(output["triggers"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
							pr.NewProperty("a trigger"),
						})))
						Expect(inputs["args"]).To(Equal(output["args"]))
					},
				},
				{ // change a trigger
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"tmpdir": true,
						},
						"triggers": []string{"an updated trigger"},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"].V).NotTo(BeEmpty())
						Expect(output["exitCode"].V).To(BeEquivalentTo(0))
						Expect(output["triggers"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
							pr.NewProperty("an updated trigger"),
						})))
						Expect(inputs["args"]).To(Equal(output["args"]))
					},
				},
			},
		})
	})

	It("should not execute when unchanged", func(ctx context.Context) {
		var firstDir string

		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"tmpdir": true,
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"].V).NotTo(BeEmpty())
					Expect(output["exitCode"].V).To(BeEquivalentTo(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
					firstDir = output["stdout"].V.(string)
				},
			},
			Updates: []integration.Operation{
				{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"tmpdir": true,
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(firstDir).NotTo(BeEmpty())
						Expect(output["stdout"]).To(HavePropertyValue(firstDir))
						Expect(output["exitCode"].V).To(BeEquivalentTo(0))
						Expect(output["triggers"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
							pr.NewProperty("a trigger"),
						})))
						Expect(inputs["args"]).To(Equal(output["args"]))
					},
				},
			},
		})
	})

	It("should fail when template is invalid", func() {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"template": "does-not-have-enough-x",
					},
				}),
				ExpectFailure: true,
			},
		})
	})
})
