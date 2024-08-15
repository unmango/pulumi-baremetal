package tests

import (
	"context"

	"github.com/docker/go-connections/nat"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

var _ = Describe("Bootstrap", Pending, func() {
	var (
		resource tokens.Type = "baremetal:index:Bootstrap"
		server   integration.Server
		host     string
		port     nat.Port
		version  = "0.0.1-test"
	)

	BeforeEach(func(ctx context.Context) {
		var err error
		By("fetching the connection details")
		host, port, err = sshServer.ConnectionDetails(ctx)
		Expect(err).NotTo(HaveOccurred())

		By("creating a provider server")
		server = util.NewServer()

		By("configuring the provider")
		err = util.ConfigureProvider(server).
			WithProvisioner("NA", "NA").
			Configure()
		Expect(err).NotTo(HaveOccurred())
	})

	It("should complete a full lifecycle", func(ctx context.Context) {
		run(server, integration.LifeCycleTest{
			Resource: resource,
			Create: integration.Operation{
				Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
					"version": version,
					"connection": map[string]interface{}{
						"host": host,
						"port": port.Int(),
					},
				}),
				Hook: func(inputs, output pr.PropertyMap) {
				},
			},
		})

		Expect(provisioner).To(ContainFile(ctx, "/usr/local/bin/provisioner"))
	})
})
