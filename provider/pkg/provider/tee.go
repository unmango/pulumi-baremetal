package provider

import (
	"context"
	_ "embed"
	"errors"

	"github.com/pulumi/pulumi-go-provider/infer"
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

	return "", state, errors.New("TODO")
}
