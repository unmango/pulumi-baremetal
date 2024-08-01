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
	Append bool     `pulumi:"append,optional"`
	Files  []string `pulumi:"files"`
}

type TeeArgs struct {
	Stdin  string   `pulumi:"stdin"`
	Create *TeeOpts `pulumi:"create,optional"`
}

type TeeState struct {
	TeeArgs
	CommandState
}

var _ = (infer.CustomCreate[TeeArgs, TeeState])((*Tee)(nil))
var _ = (infer.CustomDelete[TeeState])((*Tee)(nil))

// Create implements infer.CustomCreate.
func (Tee) Create(ctx context.Context, name string, inputs TeeArgs, preview bool) (string, TeeState, error) {
	state := TeeState{}
	log := logger.FromContext(ctx)

	if preview {
		// Could dial the host and warn if the connection fails
		log.Debug("skipping during preview")
		return name, state, nil
	}

	err := state.create(ctx,
		pb.Bin_BIN_TEE,
		squash(*inputs.Create),
		&inputs.Stdin,
	)
	if err != nil {
		log.Error("failed creating")
		return name, state, fmt.Errorf("create: %w", err)
	}

	return name, state, nil
}

// Delete implements infer.CustomDelete.
func (Tee) Delete(ctx context.Context, id string, props TeeState) error {
	log := logger.FromContext(ctx)
	if err := props.delete(ctx); err != nil {
		log.Error("failed deleting")
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func squash(inputs TeeOpts) []string {
	args := []string{}
	if inputs.Append {
		args = append(args, "--append")
	}

	return append(args, inputs.Files...)
}
