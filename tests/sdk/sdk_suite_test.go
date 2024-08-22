package sdk_test

import (
	"context"
	"io"
	"os"
	"path"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

var (
	newTester    func(integration.ProgramTestOptions) *integration.ProgramTester
	workDir, sdk string
	sdkOptions   = integration.ProgramTestOptions{}
)

var _ = BeforeSuite(func(ctx context.Context) {
	By("configuring the working directory")
	cwd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())
	workDir = path.Join(cwd, "..", "..")
	Expect(workDir).NotTo(BeNil())

	sdkEnv, ok := os.LookupEnv("SDK")
	Expect(ok).To(BeTrueBecause("SDK env not set"))
	sdk = sdkEnv

	switch sdk {
	case "dotnet":
		sdkOptions = integration.ProgramTestOptions{
			Dir:          path.Join(workDir, "examples", "dotnet"),
			Dependencies: []string{"UnMango.Baremetal"},
		}
	case "nodejs":
		sdkOptions = integration.ProgramTestOptions{
			Dir:          path.Join(workDir, "examples", "nodejs"),
			Dependencies: []string{"@unmango/baremetal"},
		}
	}
})

func TestSdk(t *testing.T) {
	newTester = func(test integration.ProgramTestOptions) *integration.ProgramTester {
		return integration.ProgramTestManualLifeCycle(t, &test)
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Sdk Suite")
}

var _ = Describe("SDK test", Ordered, func() {
	var tester *integration.ProgramTester

	BeforeAll(func() {
		By("configuring the test")
		test := baseOptions(GinkgoWriter).With(sdkOptions)

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

func baseOptions(out io.Writer) integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		RunUpdateTest: true,
		Stdout:        out,
		Stderr:        out,
		Config: map[string]string{
			"baremetal:address": "provisioner-test",
			"baremetal:port":    "4200",
		},
		LocalProviders: []integration.LocalDependency{{
			Package: "baremetal",
			Path:    path.Join("..", "..", "bin"),
		}},
	}
}
