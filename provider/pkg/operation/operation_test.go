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

	It("should handle nil result", func() {
		o := &pb.Operation{
			Command: &pb.Command{
				Bin:  pb.Bin_BIN_CHMOD,
				Args: []string{"some", "thing"},
			},
			Result: nil,
		}

		res := operation.Display(o)

		Expect(res).To(Equal("chmod some thing"))
	})
})

var _ = Describe("FromCreate", func() {
	It("should create an operation", func() {
		command := &pb.Command{Bin: pb.Bin_BIN_CAT}
		res := &pb.CreateResponse{
			Result:       &pb.Result{ExitCode: 0},
			CreatedFiles: []string{"thing"},
			MovedFiles:   map[string]string{"thing": "other-thing"},
		}

		o := operation.FromCreate(command, res)

		Expect(o).NotTo(BeNil())
		Expect(o.Command).To(Equal(command))
		Expect(o.Result).To(Equal(res.Result))
		Expect(o.CreatedFiles).To(Equal(res.CreatedFiles))
		Expect(o.MovedFiles).To(Equal(res.MovedFiles))
	})

	It("should handle nil command", func() {
		res := &pb.CreateResponse{
			Result:       &pb.Result{ExitCode: 0},
			CreatedFiles: []string{"thing"},
			MovedFiles:   map[string]string{"thing": "other-thing"},
		}

		o := operation.FromCreate(nil, res)

		Expect(o).NotTo(BeNil())
		Expect(o.Command).To(BeNil())
		Expect(o.Result).To(Equal(res.Result))
		Expect(o.CreatedFiles).To(Equal(res.CreatedFiles))
		Expect(o.MovedFiles).To(Equal(res.MovedFiles))
	})

	It("should handle nil response", func() {
		command := &pb.Command{Bin: pb.Bin_BIN_CAT}

		o := operation.FromCreate(command, nil)

		Expect(o).NotTo(BeNil())
		Expect(o.Command).To(Equal(command))
		Expect(o.Result).To(BeNil())
		Expect(o.CreatedFiles).To(BeEmpty())
		Expect(o.MovedFiles).To(BeEmpty())
	})
})
