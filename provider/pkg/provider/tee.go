package provider

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/gob"

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

	Stderr string `pulumi:"stderr"`
	Stdout string `pulumi:"stdout"`
}

func (Tee) Create(ctx context.Context, name string, input TeeArgs, preview bool) (string, TeeState, error) {
	state := TeeState{}
	if err := state.create(ctx, input); err != nil {
		return name, state, err
	}

	return name, state, nil
}

func (state *TeeState) create(ctx context.Context, input TeeArgs) error {
	c := infer.GetConfig[Config](ctx)

	p, err := c.provisioner()
	if err != nil {
		return err
	}
	defer p.conn.Close()

	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err = enc.Encode(input); err != nil {
		return err
	}

	res, err := p.Cmd.Command(ctx, &pb.CommandRequest{
		PulumiRaw: buf.Bytes(),
	})
	if err != nil {
		return err
	}

	state.Stderr = res.Stderr
	state.Stdout = res.Stdout

	return nil
}
