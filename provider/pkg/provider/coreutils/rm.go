package coreutils

import (
	"context"
	"fmt"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type RmArgs struct {
	cmd.ArgsBase

	Dir           bool     `pulumi:"dir,optional"`
	Files         []string `pulumi:"files"`
	Force         bool     `pulumi:"force,optional"`
	Help          bool     `pulumi:"help,optional"`
	OneFileSystem bool     `pulumi:"oneFileSystem,optional"`
	Recursive     bool     `pulumi:"recursive,optional"`
	Verbose       bool     `pulumi:"verbose,optional"`
}

// Cmd implements CommandArgs.
func (r RmArgs) Cmd() (*pb.Command, error) {
	b := cmd.B{Args: r.Files}

	b.Op(r.Dir, "--dir")
	b.Op(r.Force, "--force")
	b.Op(r.Help, "--help")
	b.Op(r.OneFileSystem, "--one-file-system")
	b.Op(r.Verbose, "--verbose")

	return &pb.Command{
		Bin:  pb.Bin_BIN_RM,
		Args: b.Args,
	}, nil
}

var _ cmd.Builder = RmArgs{}

type Rm struct{}

type RmState = cmd.State[RmArgs]

// Create implements infer.CustomCreate.
func (Rm) Create(ctx context.Context, name string, inputs cmd.CommandArgs[RmArgs], preview bool) (id string, output RmState, err error) {
	state := RmState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("rm: %w", err)
	}

	return name, state, nil
}

// Diff implements infer.CustomDiff.
func (Rm) Diff(ctx context.Context, id string, olds RmState, news cmd.CommandArgs[RmArgs]) (provider.DiffResponse, error) {
	diff, err := olds.Diff(ctx, news)
	if err != nil {
		return provider.DiffResponse{}, fmt.Errorf("rm: %w", err)
	}

	if news.Args.Dir != olds.Args.Dir {
		diff["args.dir"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if !slices.Equal(news.Args.Files, olds.Args.Files) {
		diff["args.files"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Force != olds.Args.Force {
		diff["args.force"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.OneFileSystem != olds.Args.OneFileSystem {
		diff["args.oneFileSystem"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Recursive != olds.Args.Recursive {
		diff["args.recursive"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

// Update implements infer.CustomUpdate.
func (Rm) Update(ctx context.Context, id string, olds RmState, news cmd.CommandArgs[RmArgs], preview bool) (RmState, error) {
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

var _ = (infer.CustomCreate[cmd.CommandArgs[RmArgs], RmState])((*Rm)(nil))
var _ = (infer.CustomDiff[cmd.CommandArgs[RmArgs], RmState])((*Rm)(nil))
var _ = (infer.CustomUpdate[cmd.CommandArgs[RmArgs], RmState])((*Rm)(nil))
var _ = (infer.CustomDelete[RmState])((*Rm)(nil))
