package cmd

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
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

	CreatedFiles []string `pulumi:"createdFiles"`
	Stderr       string   `pulumi:"stderr"`
	Stdout       string   `pulumi:"stdout"`
}

var _ = (infer.CustomCreate[TeeArgs, TeeState])((*Tee)(nil))
var _ = (infer.CustomUpdate[TeeArgs, TeeState])((*Tee)(nil))
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

	if err := state.create(ctx, inputs); err != nil {
		log.Error("failed creating")
		return name, state, fmt.Errorf("create: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (t *Tee) Update(ctx context.Context, id string, olds TeeState, news TeeArgs, preview bool) (TeeState, error) {
	panic("unimplemented")
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

func (state *TeeState) create(ctx context.Context, input TeeArgs) error {
	c := controller{pb.Op_OP_CREATE, pb.Command_COMMAND_TEE}
	res, err := c.run(ctx,
		withArgs(input.Create.Files),
		withFlag("--append", input.Create.Append),
		withStdin(input.Stdin),
	)

	if err != nil {
		return fmt.Errorf("running command controller: %w", err)
	}

	state.CreatedFiles = input.Create.Files
	state.Stderr = res.Stderr
	state.Stdout = res.Stdout

	return nil
}

func (state *TeeState) delete(ctx context.Context) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	log.Debug("sending command request to provisioner")
	res, err := p.Command(ctx, &pb.CommandRequest{
		Op:      pb.Op_OP_DELETE,
		Command: pb.Command_COMMAND_TEE,
		Args:    state.CreatedFiles,
		Flags:   map[string]*pb.Flag{},
	})
	if err != nil {
		log.Error("failed sending command request")
		return fmt.Errorf("command request: %w", err)
	}

	log.Debug("assigning outputs")
	state.CreatedFiles = []string{}
	state.Stderr = res.Stderr
	state.Stdout = res.Stdout

	log.Debug("finished delete")
	return nil
}
