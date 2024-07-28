package tests

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	baremetal "github.com/unmango/pulumi-baremetal/provider"
)

var _ = Describe("Provider", Ordered, func() {
	var server integration.Server
	var provisioner *testProvisioner

	BeforeAll(func(ctx context.Context) {
		By("creating a provisioner")
		prov, err := NewTestProvisioner(ctx, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		provisioner = prov

		By("creating a provider server")
		server = integration.NewServer(
			baremetal.Name,
			semver.MustParse("1.0.0"),
			baremetal.Provider(),
		)
	})

	It("should create a tee", func() {
		By("creating the resource")
		response, err := server.Create(p.CreateRequest{
			Urn: urn("Tee"),
			Properties: resource.PropertyMap{
				"stdin": resource.NewStringProperty("test"),
				"files": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewStringProperty("test"),
				}),
			},
			Preview: false,
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeNil())
	})

	AfterAll(func(ctx context.Context) {
		By("stopping the provisioner")
		err := provisioner.Stop(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
})

// urn is a helper function to build an urn for running integration tests.
func urn(typ string) resource.URN {
	return resource.NewURN("stack", "proj", "",
		tokens.Type("test:index:"+typ), "name")
}
