package sdk

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func baseOptions(root, rel string, out io.Writer) integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Dir:           path.Join(root, "examples", rel),
		Dependencies:  []string{},
		Config:        map[string]string{},
		RunUpdateTest: true,
		Stdout:        out,
		Stderr:        out,
	}
}

func Test(t *testing.T, dir string, opts integration.ProgramTestOptions) *integration.ProgramTester {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	root := path.Clean(path.Join(wd, "..", "..", ".."))
	opts = baseOptions(root, dir, ginkgo.GinkgoWriter).With(opts)

	return integration.ProgramTestManualLifeCycle(t, &opts)
}

func DescribeSdk(desc string, test *integration.ProgramTester) bool {
	return ginkgo.Describe(desc, ginkgo.Ordered, func() {
		ginkgo.It("TestLifeCyclePrepare", func() {
			err := test.TestLifeCyclePrepare()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			ginkgo.DeferCleanup(test.TestCleanUp)
		})

		ginkgo.It("TestLifeCycleInitialize", func() {
			err := test.TestLifeCycleInitialize()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("TestPreviewUpdateAndEdits", func() {
			err := test.TestPreviewUpdateAndEdits()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("TestLifeCycleDestroy", func() {
			err := test.TestLifeCycleDestroy()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
	})
}
