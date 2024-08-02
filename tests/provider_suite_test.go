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
	pcerts      *util.HostCerts
	sshServer   util.SshServer
)

func TestProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider Suite")
}

var _ = BeforeSuite(func(ctx context.Context) {
	By("creating a provisioner")
	prov, err := util.NewProvisioner("6969", os.Stdout)
	Expect(err).NotTo(HaveOccurred())

	By("starting the provisioner")
	err = prov.Start(ctx)
	Expect(err).NotTo(HaveOccurred())
	provisioner = prov

	By("fetching the provisioner IP")
	ip, err := prov.Ip(ctx)
	Expect(err).NotTo(HaveOccurred())

	By("generating certificates")
	bundle, err := util.NewCertBundle(ip, "provisioner")
	Expect(err).NotTo(HaveOccurred())

	By("copying certs into the provisioner")
	pcerts, err = prov.WithCerts(ctx, bundle)
	Expect(err).NotTo(HaveOccurred())

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
