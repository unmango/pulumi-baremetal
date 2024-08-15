package expect

import (
	"context"
	"fmt"
	"reflect"

	"github.com/onsi/gomega"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/unmango/pulumi-baremetal/tests/services"
)

type containFile struct {
	Context context.Context
	File    string
}

// Match implements types.GomegaMatcher.
func (e *containFile) Match(actual interface{}) (success bool, err error) {
	if host, ok := actual.(*services.Host); ok {
		return host.FileExists(e.Context, e.File)
	}

	if host, ok := actual.(*services.Provisioner); ok {
		return host.FileExists(e.Context, e.File)
	}

	if ctr, ok := actual.(tc.Container); ok {
		return services.FileExists(e.Context, ctr, e.File)
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
