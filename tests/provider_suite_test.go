package tests

import (
	"context"
	"os/exec"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var repoRoot string

func TestProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider Suite")
}

var _ = BeforeSuite(func(ctx context.Context) {
	out, err := exec.
		CommandContext(ctx, "git", "rev-parse", "--show-toplevel").
		Output()

	Expect(err).NotTo(HaveOccurred())
	repoRoot = strings.TrimSpace(string(out))
})
