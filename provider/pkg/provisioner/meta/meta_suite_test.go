package meta_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMeta(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provisioner Meta Service Suite")
}
