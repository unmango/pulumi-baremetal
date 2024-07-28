package provider

import (
	"context"
	_ "embed"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

//go:embed tee.man
var resourceDoc string

type Tee struct{}

func (t *Tee) Annotate(a infer.Annotator) {
	a.Describe(&t, resourceDoc)
}

type TeeArgs struct {
	Files []string `pulumi:"files"`
	Stdin string   `pulumi:"stdin"`
}

type TeeState struct {
	TeeArgs
	Stdout string `pulumi:"stdout"`
}

func (Tee) Create(ctx context.Context, name string, input TeeArgs, preview bool) (string, TeeState, error) {
	state := TeeState{}
	c := infer.GetConfig[Config](ctx)

	p, err := c.provisioner()
	if err != nil {
		return name, state, err
	}
	defer p.conn.Close()

	msg := "Hi friend"
	tState := pb.State{Pulumi: []byte(msg)}

	res, err := p.Cmd.Tee(ctx, &pb.TeeRequest{State: &tState})
	if err != nil {
		return name, state, err
	}

	state.Stdout = string(res.State.Pulumi)

	return name, state, nil
}
