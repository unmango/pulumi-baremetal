package coreutils

import (
	"context"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type ChmodArgs struct {
	cmd.ArgsBase

	Files          []string `pulumi:"files"`
	Mode           []string `pulumi:"mode,optional"`
	OctalMode      *string  `pulumi:"octalMode,optional"`
	Changes        *bool    `pulumi:"changes,optional"`
	NoPreserveRoot *bool    `pulumi:"noPreserveRoot,optional"`
	PreserveRoot   *bool    `pulumi:"preserveRoot,optional"`
	Quiet          *bool    `pulumi:"quiet,optional"`
	Reference      *string  `pulumi:"reference,optional"`
	Recursive      *bool    `pulumi:"recursive,optional"`
	Verbose        *bool    `pulumi:"verbose,optional"`
	Help           *bool    `pulumi:"help,optional"`
	Version        *bool    `pulumi:"version,optional"`
}

// Cmd implements CommandArgs.
func (m ChmodArgs) Cmd() (*pb.Command, error) {
	b := cmd.B{}
	b.OpP(m.Changes, "--changes")
	b.OpP(m.NoPreserveRoot, "--no-preserve-root")
	b.OpP(m.PreserveRoot, "--preserve-root")
	b.OpP(m.Quiet, "--quiet")
	b.OpvP(m.Reference, "--reference")
	b.OpP(m.Recursive, "--recursive")
	b.OpP(m.Verbose, "--verbose")
	b.OpP(m.Help, "--help")
	b.OpP(m.Version, "--version")

	b.Arg(strings.Join(m.Mode, ","))
	b.ArgP(m.OctalMode)

	for _, f := range m.Files {
		b.Arg(f)
	}

	return &pb.Command{
		Bin:  pb.Bin_BIN_CHMOD,
		Args: b.Args,
	}, nil
}

var _ cmd.Builder = ChmodArgs{}

type Chmod struct{}

type ChmodState = cmd.State[ChmodArgs]

// Create implements infer.CustomCreate.
func (Chmod) Create(ctx context.Context, name string, inputs cmd.CommandArgs[ChmodArgs], preview bool) (id string, output ChmodState, err error) {
	state := ChmodState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("chmod: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Chmod) Update(ctx context.Context, id string, olds ChmodState, news cmd.CommandArgs[ChmodArgs], preview bool) (ChmodState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("chmod: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Chmod) Delete(ctx context.Context, id string, props ChmodState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}

	// TODO: Provisioner: Read current perms before modifying them
	return nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[ChmodArgs], ChmodState])((*Chmod)(nil))
var _ = (infer.CustomUpdate[cmd.CommandArgs[ChmodArgs], ChmodState])((*Chmod)(nil))
var _ = (infer.CustomDelete[MkdirState])((*Mkdir)(nil))
