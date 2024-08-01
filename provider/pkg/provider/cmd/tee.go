package cmd

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
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
	log := logger.FromContext(ctx)

	if inputs.Create == nil {
		log.Info("nothing to do")
		return name, state, nil
	}

	if preview {
		// Could dial the host and warn if the connection fails
		log.Debug("skipping during preview")
		return name, state, nil
	}

	err := state.Create(ctx, *inputs.Create)
	if err != nil {
		log.Error("failed creating")
		return name, state, fmt.Errorf("create: %w", err)
	}

	return name, state, nil
}

// Delete implements infer.CustomDelete.
func (Tee) Delete(ctx context.Context, id string, props TeeState) error {
	log := logger.FromContext(ctx)
	if err := props.Delete(ctx); err != nil {
		log.Error("failed deleting")
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
