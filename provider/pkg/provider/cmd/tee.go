package cmd

import (
	"context"
	_ "embed"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider"
)

//go:embed tee.man
var resourceDoc string

type Tee struct{}

func (t *Tee) Annotate(a infer.Annotator) {
	a.Describe(&t, resourceDoc)
}

type TeeOpts struct {
	Files []string `pulumi:"files"`
}

type TeeArgs struct {
	Stdin string `pulumi:"stdin"`

	Create *TeeOpts `pulumi:"create,optional"`
	Update *TeeOpts `pulumi:"create,optional"`
	Delete *TeeOpts `pulumi:"create,optional"`
}

type TeeState struct {
	TeeArgs

	CreatedFiles []string `pulumi:"created_files"`
	Stderr       string   `pulumi:"stderr"`
	Stdout       string   `pulumi:"stdout"`
}

var _ infer.CustomCreate[TeeArgs, TeeState] = Tee{}
var _ infer.CustomDelete[TeeState] = Tee{}

// Create implements infer.CustomCreate.
func (Tee) Create(ctx context.Context, name string, inputs TeeArgs, preview bool) (string, TeeState, error) {
	state := TeeState{}
	if err := state.create(ctx, inputs); err != nil {
		return name, state, err
	}

	return name, state, nil
}

// Delete implements infer.CustomDelete.
func (Tee) Delete(ctx context.Context, id string, props TeeState) error {
	return props.delete(ctx)
}

func (state *TeeState) create(ctx context.Context, input TeeArgs) error {
	p, err := provisioner(ctx)
	if err != nil {
		return err
	}

	res, err := p.Cmd().Command(ctx, &pb.CommandRequest{
		Op:      pb.Op_OP_CREATE,
		Command: pb.Command_COMMAND_TEE,
		Args:    input.Create.Files,
		Flags:   map[string]*pb.Flag{},
		Stdin:   input.Stdin,
	})
	if err != nil {
		return err
	}

	state.CreatedFiles = input.Create.Files
	state.Stderr = res.Stderr
	state.Stdout = res.Stdout

	return nil
}

func (state *TeeState) delete(ctx context.Context) error {
	p, err := provisioner(ctx)
	if err != nil {
		return err
	}

	res, err := p.Cmd().Command(ctx, &pb.CommandRequest{
		Op:      pb.Op_OP_DELETE,
		Command: pb.Command_COMMAND_TEE,
		Args:    state.CreatedFiles,
		Flags:   map[string]*pb.Flag{},
	})
	if err != nil {
		return err
	}

	state.CreatedFiles = []string{}
	state.Stderr = res.Stderr
	state.Stdout = res.Stdout

	return nil
}

func provisioner(ctx context.Context) (provider.Provisioner, error) {
	c := infer.GetConfig[provider.Config](ctx)
	return c.NewProvisioner()
}
