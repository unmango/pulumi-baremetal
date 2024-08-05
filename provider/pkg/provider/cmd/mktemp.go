package cmd

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type MktempArgs struct {
	DefaultFileManipulator
	Template  string `pulumi:"template"`
	Directory bool   `pulumi:"directory,optional"`
	DryRun    bool   `pulumi:"dryRun,optional"`
	Quiet     bool   `pulumi:"quiet,optional"`
	Suffix    string `pulumi:"suffix,optional"`
	P         string `pulumi:"p,optional"`
	TmpDir    bool   `pulumi:"tmpdir,optional"`
	T         bool   `pulumi:"t,optional"`
	Help      bool   `pulumi:"help,optional"`
	Version   bool   `pulumi:"version,optional"`
}

// Cmd implements CommandArgs.
func (m MktempArgs) Cmd() *pb.Command {
	b := builder{}
	b.op(m.Directory, "--directory")
	b.op(m.DryRun, "--dry-run")
	b.op(m.Quiet, "--quiet")
	b.opv(m.Suffix, "--suffix")
	b.opv(m.P, "-p")
	b.op(m.TmpDir, "--tmpdir")
	b.op(m.T, "-t")
	b.op(m.Help, "--help")
	b.op(m.Version, "--version")

	if m.Template != "" {
		b.arg(m.Template)
	}

	return &pb.Command{
		Bin:  pb.Bin_BIN_MKDIR,
		Args: b.args,
	}
}

var _ CommandArgs = MktempArgs{}

type Mktemp struct{}

type MktempState = CommandState[MktempArgs]

// Create implements infer.CustomCreate.
func (Mktemp) Create(ctx context.Context, name string, inputs MktempArgs, preview bool) (id string, output MktempState, err error) {
	state := MktempState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("mktemp: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Mktemp) Update(ctx context.Context, id string, olds MktempState, news MktempArgs, preview bool) (MktempState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("mktemp: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Mktemp) Delete(ctx context.Context, id string, props MktempState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("mktemp: %w", err)
	}

	return nil
}

var _ = (infer.CustomCreate[MktempArgs, MktempState])((*Mktemp)(nil))
var _ = (infer.CustomUpdate[MktempArgs, MktempState])((*Mktemp)(nil))
var _ = (infer.CustomDelete[MktempState])((*Mktemp)(nil))
