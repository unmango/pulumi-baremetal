package tests

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	provisioner testProvisioner
)

func TestProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider Suite")
}

var _ = BeforeSuite(func(ctx context.Context) {
	By("creating a provisioner")
	prov, err := NewTestProvisioner(ctx, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	err = prov.Start(ctx)
	Expect(err).NotTo(HaveOccurred())
	provisioner = prov
})

var _ = AfterSuite(func(ctx context.Context) {
	By("stopping the provisioner")
	err := provisioner.Stop(ctx)
	Expect(err).NotTo(HaveOccurred())
})
