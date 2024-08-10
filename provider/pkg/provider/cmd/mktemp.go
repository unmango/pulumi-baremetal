package cmd

import (
	"context"
	"fmt"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type MktempArgs struct {
	CommandArgsBase

	Template  string `pulumi:"template,optional"`
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
		Bin:  pb.Bin_BIN_MKTEMP,
		Args: b.args,
	}
}

var _ CommandBuilder = MktempArgs{}

type Mktemp struct{}

type MktempState = CommandState[MktempArgs]

// Create implements infer.CustomCreate.
func (Mktemp) Create(ctx context.Context, name string, inputs CommandArgs[MktempArgs], preview bool) (id string, output MktempState, err error) {
	state := MktempState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("mktemp: %w", err)
	}

	return name, state, nil
}

// Diff implements infer.CustomDiff.
func (Mktemp) Diff(ctx context.Context, id string, olds MktempState, news CommandArgs[MktempArgs]) (provider.DiffResponse, error) {
	diff, err := olds.Diff(ctx, news)
	if err != nil {
		return provider.DiffResponse{}, fmt.Errorf("mv: %w", err)
	}

	if news.Args.Directory != olds.Args.Directory {
		diff["args.directory"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.DryRun != olds.Args.DryRun {
		diff["args.destination"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Quiet != olds.Args.Quiet {
		diff["args.quiet"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Suffix != olds.Args.Suffix {
		diff["args.suffix"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.T != olds.Args.T {
		diff["args.t"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Template != olds.Args.Template {
		diff["args.template"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.TmpDir != olds.Args.TmpDir {
		diff["args.tmpDir"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Directory != olds.Args.Directory {
		diff["args.suffix"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

// Update implements infer.CustomUpdate.
func (Mktemp) Update(ctx context.Context, id string, olds MktempState, news CommandArgs[MktempArgs], preview bool) (MktempState, error) {
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

var _ = (infer.CustomCreate[CommandArgs[MktempArgs], MktempState])((*Mktemp)(nil))
var _ = (infer.CustomDiff[CommandArgs[MktempArgs], MktempState])((*Mktemp)(nil))
var _ = (infer.CustomUpdate[CommandArgs[MktempArgs], MktempState])((*Mktemp)(nil))
var _ = (infer.CustomDelete[MktempState])((*Mktemp)(nil))
