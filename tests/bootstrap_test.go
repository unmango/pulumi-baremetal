package tests

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

var _ = Describe("Bootstrap", Ordered, func() {
	urn := util.Urn("Bootstrap")
	var server integration.Server

	version := "0.0.1-test"
	props := resource.PropertyMap{
		"version": resource.NewStringProperty(version),
	}

	BeforeAll(func(ctx context.Context) {
		By("fetching the connection props")
		conn, err := sshServer.ConnectionProps(ctx)
		Expect(err).NotTo(HaveOccurred())
		props["connection"] = conn
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

	It("should construct", func() {
		_, err := server.Construct(p.ConstructRequest{
			URN:     urn,
			Preview: false,
			Construct: func(ctx context.Context, cf p.ConstructFunc) (p.ConstructResponse, error) {
				res := p.ConstructResponse{}
				_, err := cf(nil, provider.ConstructInputs{}, nil)
				if err != nil {
					return res, err
				}

				return res, nil
			},
		})

		Expect(err).NotTo(HaveOccurred())
		// binPath, ok := response.Properties["binPath"].V.(string)
		// Expect(ok).To(BeTrueBecause("binPath was not a string"))
		// Expect(binPath).To(Equal("/usr/local/bin/provisioner"))
	})
})
