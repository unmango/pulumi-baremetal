package config

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type ImagesCommand string

var (
	List ImagesCommand = "list"
	Pull ImagesCommand = "pull"
)

var _ = (infer.Enum[ImagesCommand])((*ImagesCommand)(nil))

func (*ImagesCommand) Values() []infer.EnumValue[ImagesCommand] {
	return []infer.EnumValue[ImagesCommand]{
		{Value: List, Description: "Print a list of images kubeadm will use. The configuration file is used in case any images or image repositories are customized"},
		{Value: Pull, Description: "Pull images used by kubeadm"},
	}
}

type ImagesArgs struct {
	cmd.ArgsBase

	Command           ImagesCommand     `pulumi:"command"`
	Config            string            `pulumi:"config,optional"`
	CriSocket         string            `pulumi:"criSocket,optional"`
	FeatureGates      map[string]string `pulumi:"featureGates,optional"`
	ImageRepository   string            `pulumi:"imageRepository,optional"`
	KubernetesVersion string            `pulumi:"kubernetesVersion,optional"`
	Kubeconfig        string            `pulumi:"kubeconfig,optional"`
	Rootfs            string            `pulumi:"rootfs,optional"`
}

func (a ImagesArgs) Cmd() (*pb.Command, error) {
	return builder(func(b *cmd.B) {
		b.Arg("images")
		b.Arg(string(a.Command))
		b.Opv(a.Config, "--config")
		b.Opv(a.CriSocket, "--cri-socket")
		b.Opv(a.ImageRepository, "--image-repository")
		b.Opv(a.KubernetesVersion, "--kubernetes-version")
		b.Opv(a.Kubeconfig, "--kubeconfig")
		b.Opv(a.Rootfs, "--rootfs")
	}), nil
}

type Images struct{}

type ImagesState = cmd.State[ImagesArgs]

// Create implements infer.CustomCreate.
func (Images) Create(ctx context.Context, name string, inputs cmd.CommandArgs[ImagesArgs], preview bool) (string, ImagesState, error) {
	state := ImagesState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("images: %w", err)
	}

	return name, state, nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[ImagesArgs], ImagesState])((*Images)(nil))
