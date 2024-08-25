package cmd

import (
	"context"

	provider "github.com/pulumi/pulumi-go-provider"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/config"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type CommandArgs[T Builder] struct {
	Args         T                             `pulumi:"args"`
	Triggers     []any                         `pulumi:"triggers,optional"`
	CustomUpdate []string                      `pulumi:"customUpdate,optional"`
	CustomDelete []string                      `pulumi:"customDelete,optional"`
	Connection   *config.ProvisionerConnection `pulumi:"connection,optional"`
}

func (a *CommandArgs[T]) Provisioner(ctx context.Context) (*provisioner.Provisioner, error) {
	if a.Connection != nil {
		return provisioner.FromConnection(*a.Connection)
	} else {
		return provisioner.FromContext(ctx)
	}
}

func (a *CommandArgs[T]) Cmd() (*pb.Command, error) {
	return a.Args.Cmd()
}

func (a *CommandArgs[T]) ExpectCreated() []string {
	return a.Args.ExpectCreated()
}

func (a *CommandArgs[T]) ExpectMoved() map[string]string {
	return a.Args.ExpectMoved()
}

func (a CommandArgs[T]) UpdateKind() provider.DiffKind {
	if len(a.CustomUpdate) > 0 {
		return provider.Update
	} else {
		return provider.UpdateReplace
	}
}

func (a CommandArgs[T]) DeleteBeforeReplace() bool {
	return len(a.CustomUpdate) == 0
}

type ArgsBase struct{}

func (ArgsBase) ExpectCreated() []string {
	return []string{}
}

func (ArgsBase) ExpectMoved() map[string]string {
	return map[string]string{}
}
