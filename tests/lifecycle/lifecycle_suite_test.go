package lifecycle

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/pulumi-baremetal/tests/services"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

var (
	provisioner *services.Provisioner
	sshServer   *services.Sshd
	clientCerts *util.CertBundle
)

var _ = BeforeSuite(func(ctx context.Context) {
	var err error

	By("generating client certs")
	clientCerts, err = util.NewCertBundle("ca", "pulumi")
	Expect(err).NotTo(HaveOccurred())

	By("creating a provisioner")
	prov, err := services.NewProvisioner("6969", clientCerts.Ca, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	By("starting the provisioner")
	err = prov.Start(ctx)
	Expect(err).NotTo(HaveOccurred())
	provisioner = prov

	By("creating an ssh server")
	ssh, err := services.NewSshd(ctx, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	By("starting the ssh server")
	err = ssh.Start(ctx)
	Expect(err).NotTo(HaveOccurred())
	sshServer = ssh
})

func TestLifecycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lifecycle Suite")
}

var _ = AfterSuite(func(ctx context.Context) {
	if provisioner != nil {
		By("stopping the provisioner")
		// TODO: Why does this error but the sshServer does not
		_ = provisioner.Stop(ctx)
	}

	if sshServer != nil {
		By("stopping the ssh server")
		err := sshServer.Stop(ctx)
		Expect(err).NotTo(HaveOccurred())
	}
})
