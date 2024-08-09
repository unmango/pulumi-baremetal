package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type ChmodArgs struct {
	CommandArgsBase

	Files          []string `pulumi:"files"`
	Mode           []string `pulumi:"mode,optional"`
	OctalMode      string   `pulumi:"octalMode,optional"`
	Changes        bool     `pulumi:"changes,optional"`
	NoPreserveRoot bool     `pulumi:"noPreserveRoot,optional"`
	PreserveRoot   bool     `pulumi:"preserveRoot,optional"`
	Quiet          bool     `pulumi:"quiet,optional"`
	Reference      string   `pulumi:"reference,optional"`
	Recursive      bool     `pulumi:"recursive,optional"`
	Verbose        bool     `pulumi:"verbose,optional"`
	Help           bool     `pulumi:"help,optional"`
	Version        bool     `pulumi:"version,optional"`
}

// Cmd implements CommandArgs.
func (m ChmodArgs) Cmd() *pb.Command {
	b := builder{}
	b.op(m.Changes, "--changes")
	b.op(m.NoPreserveRoot, "--no-preserve-root")
	b.op(m.PreserveRoot, "--preserve-root")
	b.op(m.Quiet, "--quiet")
	b.opv(m.Reference, "--reference")
	b.op(m.Recursive, "--recursive")
	b.op(m.Verbose, "--verbose")
	b.op(m.Help, "--help")
	b.op(m.Version, "--version")

	b.arg(strings.Join(m.Mode, ","))
	b.arg(m.OctalMode)

	for _, f := range m.Files {
		b.arg(f)
	}

	return &pb.Command{
		Bin:  pb.Bin_BIN_CHMOD,
		Args: b.args,
	}
}

var _ CommandBuilder = ChmodArgs{}

type Chmod struct{}

type ChmodState = CommandState[ChmodArgs]

// Create implements infer.CustomCreate.
func (Chmod) Create(ctx context.Context, name string, inputs CommandArgs[ChmodArgs], preview bool) (id string, output ChmodState, err error) {
	state := ChmodState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("chmod: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Chmod) Update(ctx context.Context, id string, olds ChmodState, news CommandArgs[ChmodArgs], preview bool) (ChmodState, error) {
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

	return nil
}

var _ = (infer.CustomCreate[CommandArgs[ChmodArgs], ChmodState])((*Chmod)(nil))
var _ = (infer.CustomUpdate[CommandArgs[ChmodArgs], ChmodState])((*Chmod)(nil))
