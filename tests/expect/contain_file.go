package expect

import (
	"context"
	"fmt"
	"reflect"

	"github.com/onsi/gomega"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

type containFile struct {
	Context context.Context
	File    string
}

// Match implements types.GomegaMatcher.
func (e *containFile) Match(actual interface{}) (success bool, err error) {
	ctx := context.Background()

	if host, ok := actual.(util.TestHost); ok {
		return host.FileExists(ctx, e.File)
	}

	if ctr, ok := actual.(tc.Container); ok {
		return util.FileExists(ctx, ctr, e.File)
	}

	return false, fmt.Errorf("unupported match target: %s", reflect.TypeOf(actual))
}

// FailureMessage implements types.GomegaMatcher.
func (e *containFile) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected file to exist in container: %s", e.File)
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (e *containFile) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected file not to exist in container: %s", e.File)
}

func ContainFile(ctx context.Context, file string) gomega.OmegaMatcher {
	return &containFile{ctx, file}
}
