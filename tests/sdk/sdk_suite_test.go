package sdk_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/unmango/pulumi-baremetal/tests/services"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

var (
	provisioner *services.Provisioner
	newTester   func(integration.ProgramTestOptions) *integration.ProgramTester
	wd          string
)

type configureTest func(integration.ProgramTestOptions) integration.ProgramTestOptions

var _ = BeforeSuite(func(ctx context.Context) {
	By("configuring the working directory")
	cwd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())
	wd = path.Join(cwd, "..", "..")
	Expect(wd).NotTo(BeNil())

	By("generating client certs")
	clientCerts, err := util.NewCertBundle("ca", "pulumi")
	Expect(err).NotTo(HaveOccurred())

	By("creating a provisioner")
	prov, err := services.NewProvisioner("4200", clientCerts.Ca, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	By("starting the provisioner")
	err = prov.Start(ctx)
	Expect(err).NotTo(HaveOccurred())
	provisioner = prov
})

func TestSdk(t *testing.T) {
	newTester = func(test integration.ProgramTestOptions) *integration.ProgramTester {
		return integration.ProgramTestManualLifeCycle(t, &test)
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Sdk Suite")
}

var _ = DescribeSdk("dotnet", func(base integration.ProgramTestOptions) integration.ProgramTestOptions {
	return base.With(integration.ProgramTestOptions{
		Dir: path.Join("..", "..", "examples", "dotnet"),
		// DotNetBin:    path.Join(wd, "bin", "dotnet", "dotnet"),
		Dependencies: []string{"UnMango.Baremetal"},
	})
})

var _ = AfterSuite(func(ctx context.Context) {
	By("stopping the provisioner")
	err := provisioner.Stop(ctx)
	Expect(err).NotTo(HaveOccurred())
})

func DescribeSdk(sdk string, configure configureTest) bool {
	return Describe(fmt.Sprintf("%s SDK test", sdk), Label(sdk), Ordered, func() {
		var tester *integration.ProgramTester

		BeforeAll(func() {
			By("configuring the test")
			test := configure(baseOptions(GinkgoWriter))

			By("creating the program tester")
			tester = newTester(test)
		})

		It("TestLifeCyclePrepare", func() {
			err := tester.TestLifeCyclePrepare()
			Expect(err).NotTo(HaveOccurred())
			DeferCleanup(tester.TestCleanUp)
		})

		It("TestLifeCycleInitialize", func() {
			err := tester.TestLifeCycleInitialize()
			Expect(err).NotTo(HaveOccurred())
		})

		It("TestPreviewUpdateAndEdits", func() {
			err := tester.TestPreviewUpdateAndEdits()
			Expect(err).NotTo(HaveOccurred())
		})

		It("TestLifeCycleDestroy", func() {
			err := tester.TestLifeCycleDestroy()
			Expect(err).NotTo(HaveOccurred())
		})
	})
}

func baseOptions(out io.Writer) integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		RunUpdateTest: true,
		Stdout:        out,
		Stderr:        out,
		Config: map[string]string{
			"baremetal:address": "localhost",
			"baremetal:port":    "4200",
		},
		LocalProviders: []integration.LocalDependency{{
			Package: "baremetal",
			Path:    path.Join("..", "..", "bin"),
		}},
	}
}
