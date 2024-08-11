package kubeadm

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type KubeadmArgs struct {
	cmd.ArgsBase

	Commands []string `pulumi:"commands"`
}

func (a KubeadmArgs) Cmd() (*pb.Command, error) {
	return Builder(func(b *cmd.B) {
		for _, c := range a.Commands {
			b.Arg(c)
		}
	}), nil
}

type Kubeadm struct{}

type KubeadmState = cmd.State[KubeadmArgs]

// Create implements infer.CustomCreate.
func (Kubeadm) Create(ctx context.Context, name string, inputs cmd.CommandArgs[KubeadmArgs], preview bool) (string, KubeadmState, error) {
	state := KubeadmState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("kubeadm: %w", err)
	}

	return name, state, nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[KubeadmArgs], KubeadmState])((*Kubeadm)(nil))
