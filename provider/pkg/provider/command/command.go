package command

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type Bin string

var (
	Chmod     Bin = "chmod"
	Mkdir     Bin = "mkdir"
	Mktemp    Bin = "mktemp"
	Mv        Bin = "mv"
	Rm        Bin = "rm"
	Systemctl Bin = "systemctl"
	Tar       Bin = "tar"
	Tee       Bin = "tee"
	Wget      Bin = "wget"
)

// Values implements infer.Enum.
func (b *Bin) Values() []infer.EnumValue[Bin] {
	return []infer.EnumValue[Bin]{
		{Value: Chmod},
		{Value: Mkdir},
		{Value: Mktemp},
		{Value: Mv},
		{Value: Rm},
		{Value: Systemctl},
		{Value: Tar},
		{Value: Tee},
		{Value: Wget},
	}
}

var _ = (infer.Enum[Bin])((*Bin)(nil))

type CommandArgs struct {
	cmd.ArgsBase

	Bin  Bin      `pulumi:"bin"`
	Args []string `pulumi:"args"`
}

func (a CommandArgs) Cmd() (*pb.Command, error) {
	b := cmd.B{}

	b.Arg(string(a.Bin))
	for _, a := range a.Args {
		b.Arg(a)
	}

	return b.Cmd(), nil
}

type Command struct{}

type CommandState = cmd.State[CommandArgs]

// Create implements infer.CustomCreate.
func (Command) Create(ctx context.Context, name string, inputs cmd.CommandArgs[CommandArgs], preview bool) (string, CommandState, error) {
	state := CommandState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("kubeadm: %w", err)
	}

	return name, state, nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[CommandArgs], CommandState])((*Command)(nil))
