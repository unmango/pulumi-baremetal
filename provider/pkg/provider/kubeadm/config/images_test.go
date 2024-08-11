package config_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/kubeadm/config"
)

var _ = Describe("Images command", func() {
	It("should properly format arguments", func() {
		args := config.ImagesArgs{Command: config.Pull}

		cmd := args.Cmd()

		Expect(cmd.Bin).To(Equal(pb.Bin_BIN_KUBEADM))
		Expect(cmd.Args).To(Equal([]string{
			"config",
			"images",
			"pull",
		}))
	})
})
