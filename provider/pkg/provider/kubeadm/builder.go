package kubeadm

import (
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

func Builder(build func(*cmd.B)) *pb.Command {
	b := cmd.B{}
	b.Bin(pb.Bin_BIN_KUBEADM)
	build(&b)
	return b.Cmd()
}
