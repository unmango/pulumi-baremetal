package cmd

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type RmArgs struct {
	Dir           bool     `pulumi:"dir,optional"`
	Files         []string `pulumi:"files"`
	Force         bool     `pulumi:"force,optional"`
	Help          bool     `pulumi:"help,optional"`
	OneFileSystem bool     `pulumi:"oneFileSystem,optional"`
	Recursive     bool     `pulumi:"recursive,optional"`
	Verbose       bool     `pulumi:"verbose,optional"`
}

// Cmd implements CommandArgs.
func (r RmArgs) Cmd() *pb.Command {
	b := builder{r.Files}
	b.op(r.Dir, "--dir")
	b.op(r.Force, "--force")
	b.op(r.Help, "--help")
	b.op(r.OneFileSystem, "--one-file-system")
	b.op(r.Verbose, "--verbose")

	return &pb.Command{
		Bin:  pb.Bin_BIN_RM,
		Args: b.args,
	}
}

// ExpectedFiles implements CommandArgs.
func (r RmArgs) ExpectedFiles() []string {
	return []string{}
}

var _ CommandArgs = RmArgs{}

type Rm struct{}

type RmState = CommandState[RmArgs]

// Create implements infer.CustomCreate.
func (Rm) Create(ctx context.Context, name string, inputs RmArgs, preview bool) (id string, output RmState, err error) {
	state := RmState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("rm: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Rm) Update(ctx context.Context, id string, olds RmState, news RmArgs, preview bool) (RmState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("rm: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Rm) Delete(ctx context.Context, id string, props RmState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("rm: %w", err)
	}

	return nil
}

var _ = (infer.CustomCreate[RmArgs, RmState])((*Rm)(nil))
var _ = (infer.CustomUpdate[RmArgs, RmState])((*Rm)(nil))
var _ = (infer.CustomDelete[RmState])((*Rm)(nil))
