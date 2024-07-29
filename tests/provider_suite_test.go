package tests

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

var (
	provisioner util.TestProvisioner
)

func TestProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider Suite")
}

var _ = BeforeSuite(func(ctx context.Context) {
	By("creating a provisioner")
	prov, err := util.NewTestProvisioner(ctx, os.Stdout)
	Expect(err).NotTo(HaveOccurred())

	By("starting the provisioner")
	err = prov.Start(ctx)
	Expect(err).NotTo(HaveOccurred())
	provisioner = prov
})

var _ = AfterSuite(func(ctx context.Context) {
	By("stopping the provisioner")
	err := provisioner.Stop(ctx)
	Expect(err).NotTo(HaveOccurred())
})
