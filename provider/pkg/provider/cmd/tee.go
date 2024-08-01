package cmd

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

//go:embed tee.man
var resourceDoc string

type Tee struct{}

func (t *Tee) Annotate(a infer.Annotator) {
	a.Describe(&t, resourceDoc)
}

type TeeOpts struct {
	Append  bool     `pulumi:"append,optional"`
	Content string   `pulumi:"content"`
	Files   []string `pulumi:"files"`
}

type TeeArgs struct {
	Create *TeeOpts `pulumi:"create,optional"`
}

type TeeState = CommandState[TeeOpts]

var _ = (infer.CustomCreate[TeeArgs, TeeState])((*Tee)(nil))
var _ = (infer.CustomDelete[TeeState])((*Tee)(nil))

// Create implements infer.CustomCreate.
func (Tee) Create(ctx context.Context, name string, inputs TeeArgs, preview bool) (string, TeeState, error) {
	state := TeeState{}
	if err := state.Create(ctx, inputs.Create, preview); err != nil {
		return name, state, fmt.Errorf("create: %w", err)
	}

	return name, state, nil
}

// Delete implements infer.CustomDelete.
func (Tee) Delete(ctx context.Context, id string, props TeeState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (o TeeOpts) Cmd() *pb.Command {
	args := []string{}
	if o.Append {
		args = append(args, "--append")
	}

	return &pb.Command{
		Bin:  pb.Bin_BIN_TEE,
		Args: append(args, o.Files...),
	}
}
