package coreutils

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type CatArgs struct {
	cmd.ArgsBase

	Files           []string `pulumi:"files"`
	ShowAll         *bool    `pulumi:"showAll,optional"`
	NumberNonblank  *bool    `pulumi:"numberNonblank,optional"`
	E               *bool    `pulumi:"e,optional"`
	ShowEnds        *bool    `pulumi:"showEnds,optional"`
	Number          *bool    `pulumi:"number,optional"`
	SqueezeBlank    *bool    `pulumi:"squeezeBlank,optional"`
	T               *bool    `pulumi:"t,optional"`
	ShowTabs        *bool    `pulumi:"showTabs,optional"`
	ShowNonprinting *bool    `pulumi:"showNonprinting,optional"`
	Help            *bool    `pulumi:"help,optional"`
	Version         *bool    `pulumi:"version,optional"`
}

// Cmd implements CommandArgs.
func (m CatArgs) Cmd() (*pb.Command, error) {
	b := cmd.B{Args: m.Files}
	b.OpP(m.ShowAll, "--show-all")
	b.OpP(m.NumberNonblank, "--number-nonblank")
	b.OpP(m.E, "-e")
	b.OpP(m.ShowEnds, "--show-ends")
	b.OpP(m.Number, "--number")
	b.OpP(m.SqueezeBlank, "--squeeze-blank")
	b.OpP(m.T, "-t")
	b.OpP(m.ShowTabs, "--show-tabs")
	b.OpP(m.ShowNonprinting, "--show-nonprinting")
	b.OpP(m.Help, "--help")
	b.OpP(m.Version, "--version")

	return &pb.Command{
		Bin:  pb.Bin_BIN_CAT,
		Args: b.Args,
	}, nil
}

var _ cmd.Builder = CatArgs{}

type Cat struct{}

type CatState = cmd.State[CatArgs]

// Create implements infer.CustomCreate.
func (Cat) Create(ctx context.Context, name string, inputs cmd.CommandArgs[CatArgs], preview bool) (id string, output CatState, err error) {
	state := CatState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("cat: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Cat) Update(ctx context.Context, id string, olds CatState, news cmd.CommandArgs[CatArgs], preview bool) (CatState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("cat: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Cat) Delete(ctx context.Context, id string, props CatState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("cat: %w", err)
	}

	return nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[CatArgs], CatState])((*Cat)(nil))
var _ = (infer.CustomUpdate[cmd.CommandArgs[CatArgs], CatState])((*Cat)(nil))
var _ = (infer.CustomDelete[CatState])((*Cat)(nil))
