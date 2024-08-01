package sdk_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pi "github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/unmango/pulumi-baremetal/tests/sdk"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

var (
	provisioner util.TestProvisioner
	test        *pi.ProgramTester
)

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

func TestSdk(t *testing.T) {
	opts := sdk.ProgramTestOptions(provisioner)
	test = sdk.Test(t, "dotnet", opts)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Dotnet Suite")
}

var _ = sdk.DescribeSdk("Sdk Test", test)
