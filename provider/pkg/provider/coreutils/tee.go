package coreutils

import (
	"context"
	_ "embed"
	"fmt"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/asset"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type TeeArgs struct {
	cmd.ArgsBase

	Append  bool        `pulumi:"append,optional"`
	Content asset.Asset `pulumi:"content,optional"`
	Files   []string    `pulumi:"files"`
	Stdin   string      `pulumi:"stdin,optional"`
}

func (o TeeArgs) Cmd() (*pb.Command, error) {
	args := []string{}
	if o.Append {
		args = append(args, "--append")
	}

	var stdin string
	if len(o.Stdin) > 0 {
		stdin = o.Stdin
	} else {
		data, err := o.Content.Bytes()
		if err != nil {
			return nil, fmt.Errorf("failed to read asset bytes: %w", err)
		}
		stdin = string(data)
	}

	return &pb.Command{
		Bin:   pb.Bin_BIN_TEE,
		Args:  append(args, o.Files...),
		Stdin: &stdin,
	}, nil
}

// ExpectCreated implements FileManipulator.
func (o TeeArgs) ExpectCreated() []string {
	return o.Files
}

//go:embed tee.man
var teeMan string

type Tee struct{}

func (t *Tee) Annotate(a infer.Annotator) {
	a.Describe(&t, teeMan)
}

type TeeState = cmd.State[TeeArgs]

// Create implements infer.CustomCreate.
func (Tee) Create(ctx context.Context, name string, inputs cmd.CommandArgs[TeeArgs], preview bool) (string, TeeState, error) {
	state := TeeState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("tee: %w", err)
	}

	return name, state, nil
}

// Diff implements infer.CustomDiff.
func (Tee) Diff(ctx context.Context, id string, olds TeeState, news cmd.CommandArgs[TeeArgs]) (provider.DiffResponse, error) {
	diff, err := olds.Diff(ctx, news)
	if err != nil {
		return provider.DiffResponse{}, fmt.Errorf("tee: %w", err)
	}

	if news.Args.Append != olds.Args.Append {
		diff["args.append"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Content != olds.Args.Content {
		diff["args.content"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if !slices.Equal(news.Args.Files, olds.Args.Files) {
		diff["args.files"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

// Update implements infer.CustomUpdate.
func (Tee) Update(ctx context.Context, id string, olds TeeState, news cmd.CommandArgs[TeeArgs], preview bool) (TeeState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("tee: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Tee) Delete(ctx context.Context, id string, props TeeState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("tee: %w", err)
	}

	return nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[TeeArgs], TeeState])((*Tee)(nil))
var _ = (infer.CustomDiff[cmd.CommandArgs[TeeArgs], TeeState])((*Tee)(nil))
var _ = (infer.CustomUpdate[cmd.CommandArgs[TeeArgs], TeeState])((*Tee)(nil))
var _ = (infer.CustomDelete[TeeState])((*Tee)(nil))
