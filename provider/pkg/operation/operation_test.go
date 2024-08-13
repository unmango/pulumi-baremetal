package operation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/operation"
)

var _ = Describe("Display", func() {
	It("should format better than just calling string", func() {
		o := &pb.Operation{
			Command: &pb.Command{
				Bin:  pb.Bin_BIN_CHMOD,
				Args: []string{"some", "thing"},
			},
			Result: &pb.Result{ExitCode: 69},
		}

		res := operation.Display(o)

		Expect(res).To(Equal("exitCode: 69, command: chmod some thing"))
	})
})
