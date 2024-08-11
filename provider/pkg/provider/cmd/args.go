package cmd

import (
	provider "github.com/pulumi/pulumi-go-provider"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type CommandArgs[T Builder] struct {
	Args         T        `pulumi:"args"`
	Triggers     []any    `pulumi:"triggers,optional"`
	CustomUpdate []string `pulumi:"customUpdate,optional"`
	CustomDelete []string `pulumi:"customDelete,optional"`
}

func (a *CommandArgs[T]) Cmd() *pb.Command {
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
