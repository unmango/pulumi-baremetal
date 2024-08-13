package command_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/command"
)

var _ = Describe("Display", func() {
	It("should format bin", func() {
		c := &pb.Command{
			Bin:  pb.Bin_BIN_MKDIR,
			Args: []string{"some", "thing"},
		}

		res := command.Display(c)

		Expect(res).To(Equal("mkdir some thing"))
	})

	When("args are nil", func() {
		var args []string = nil

		It("should format bin", func() {
			c := &pb.Command{Bin: pb.Bin_BIN_CHMOD, Args: args}

			res := command.Display(c)

			Expect(res).To(Equal("chmod"))
		})
	})

	When("args are empty", func() {
		args := []string{}

		It("should format bin", func() {
			c := &pb.Command{Bin: pb.Bin_BIN_CHMOD, Args: args}

			res := command.Display(c)

			Expect(res).To(Equal("chmod"))
		})
	})

	When("bin is nil", func() {
		It("should format as nil", func() {
			c := &pb.Command{}

			res := command.Display(c)

			Expect(res).To(Equal("<nil>"))
		})

		It("should format args", func() {
			c := &pb.Command{Args: []string{"something", "else"}}

			res := command.Display(c)

			Expect(res).To(Equal("<nil> something else"))
		})
	})
})
