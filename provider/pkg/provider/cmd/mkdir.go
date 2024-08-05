package cmd

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type MkdirArgs struct {
	DefaultFileManipulator
	Directory []string `pulumi:"directory"`
	Mode      string   `pulumi:"mode,optional"`
	Parents   bool     `pulumi:"parents,optional"`
	Verbose   bool     `pulumi:"verbose,optional"`
	Help      bool     `pulumi:"help,optional"`
	Version   bool     `pulumi:"version,optional"`
}

// Cmd implements CommandArgs.
func (m MkdirArgs) Cmd() *pb.Command {
	b := builder{m.Directory}
	b.opv(m.Mode, "--mode")
	b.op(m.Parents, "--parents")
	b.op(m.Verbose, "--verbose")
	b.op(m.Help, "--help")
	b.op(m.Version, "--version")

	return &pb.Command{
		Bin:  pb.Bin_BIN_MKDIR,
		Args: b.args,
	}
}

// ExpectCreated implements CommandArgs.
// Subtle: this method shadows the method (DefaultFileManipulator).ExpectCreated of MkdirArgs.DefaultFileManipulator.
func (m MkdirArgs) ExpectCreated() []string {
	return m.Directory
}

var _ CommandArgs = MkdirArgs{}

type Mkdir struct{}

type MkdirState = CommandState[MkdirArgs]

// Create implements infer.CustomCreate.
func (Mkdir) Create(ctx context.Context, name string, inputs MkdirArgs, preview bool) (id string, output MkdirState, err error) {
	state := MkdirState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("mkdir: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Mkdir) Update(ctx context.Context, id string, olds MkdirState, news MkdirArgs, preview bool) (MkdirState, error) {
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

var _ = (infer.CustomCreate[MkdirArgs, MkdirState])((*Mkdir)(nil))
var _ = (infer.CustomUpdate[MkdirArgs, MkdirState])((*Mkdir)(nil))
var _ = (infer.CustomDelete[MkdirState])((*Mkdir)(nil))
