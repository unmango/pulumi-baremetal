package config

import (
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/kubeadm"
)

func builder(build func(*cmd.B)) *pb.Command {
	return kubeadm.Builder(func(b *cmd.B) {
		b.Arg("config")
		build(b)
	})
}
