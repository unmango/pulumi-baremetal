package sdk_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/unmango/pulumi-baremetal/tests/sdk"
)

var test *integration.ProgramTester

func TestSdk(t *testing.T) {
	test = sdk.Test(t, "dotnet", integration.ProgramTestOptions{})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Dotnet Suite")
}

var _ = Describe("Sdk Test", func() {
	sdk.DescribeSdk("Sdk Test Inner", test)
})
