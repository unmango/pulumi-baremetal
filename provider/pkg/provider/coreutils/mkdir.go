package coreutils

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type MkdirArgs struct {
	cmd.ArgsBase

	Directory []string `pulumi:"directory"`
	Mode      *string  `pulumi:"mode,optional"`
	Parents   *bool    `pulumi:"parents,optional"`
	Verbose   *bool    `pulumi:"verbose,optional"`
	Help      *bool    `pulumi:"help,optional"`
	Version   *bool    `pulumi:"version,optional"`
}

// Cmd implements CommandArgs.
func (m MkdirArgs) Cmd() (*pb.Command, error) {
	b := cmd.B{Args: m.Directory}

	b.OpvP(m.Mode, "--mode")
	b.OpP(m.Parents, "--parents")
	b.OpP(m.Verbose, "--verbose")
	b.OpP(m.Help, "--help")
	b.OpP(m.Version, "--version")

	return &pb.Command{
		Bin:  pb.Bin_BIN_MKDIR,
		Args: b.Args,
	}, nil
}

var _ cmd.Builder = MkdirArgs{}

type Mkdir struct{}

type MkdirState = cmd.State[MkdirArgs]

// Create implements infer.CustomCreate.
func (Mkdir) Create(ctx context.Context, name string, inputs cmd.CommandArgs[MkdirArgs], preview bool) (id string, output MkdirState, err error) {
	state := MkdirState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("mkdir: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Mkdir) Update(ctx context.Context, id string, olds MkdirState, news cmd.CommandArgs[MkdirArgs], preview bool) (MkdirState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("mkdir: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Mkdir) Delete(ctx context.Context, id string, props MkdirState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	return nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[MkdirArgs], MkdirState])((*Mkdir)(nil))
var _ = (infer.CustomUpdate[cmd.CommandArgs[MkdirArgs], MkdirState])((*Mkdir)(nil))
var _ = (infer.CustomDelete[MkdirState])((*Mkdir)(nil))
