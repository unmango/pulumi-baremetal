package tests

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/pulumi-baremetal/tests/util"
)

var (
	provisioner util.TestProvisioner
	sshServer   util.SshServer
	clientCerts *util.CertBundle
)

func TestProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider Suite")
}

var _ = BeforeSuite(func(ctx context.Context) {
	var err error

	By("generating client certs")
	clientCerts, err = util.NewCertBundle("ca", "pulumi")
	Expect(err).NotTo(HaveOccurred())

	By("creating a provisioner")
	prov, err := util.NewProvisioner("6969", clientCerts.Ca, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	By("starting the provisioner")
	err = prov.Start(ctx)
	Expect(err).NotTo(HaveOccurred())
	provisioner = prov

	By("creating an ssh server")
	ssh, err := util.NewSshServer(ctx)
	Expect(err).NotTo(HaveOccurred())

	// By("starting the ssh server")
	// err = ssh.Start(ctx)
	// Expect(err).NotTo(HaveOccurred())
	sshServer = ssh
})

var _ = AfterSuite(func(ctx context.Context) {
	By("stopping the provisioner")
	err := provisioner.Stop(ctx)
	Expect(err).NotTo(HaveOccurred())

	// By("stopping the ssh server")
	// err = sshServer.Stop(ctx)
	// Expect(err).NotTo(HaveOccurred())
})
