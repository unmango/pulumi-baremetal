// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package coreutils

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/config"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

type Mktemp struct {
	pulumi.CustomResourceState

	Args         pulumix.GPtrOutput[MktempArgsType, MktempArgsTypeOutput]                             `pulumi:"args"`
	Connection   pulumix.GPtrOutput[config.ProvisionerConnection, config.ProvisionerConnectionOutput] `pulumi:"connection"`
	CreatedFiles pulumix.ArrayOutput[string]                                                          `pulumi:"createdFiles"`
	CustomDelete pulumix.ArrayOutput[string]                                                          `pulumi:"customDelete"`
	CustomUpdate pulumix.ArrayOutput[string]                                                          `pulumi:"customUpdate"`
	ExitCode     pulumix.Output[int]                                                                  `pulumi:"exitCode"`
	MovedFiles   pulumix.MapOutput[string]                                                            `pulumi:"movedFiles"`
	Stderr       pulumix.Output[string]                                                               `pulumi:"stderr"`
	Stdout       pulumix.Output[string]                                                               `pulumi:"stdout"`
	Triggers     pulumix.ArrayOutput[any]                                                             `pulumi:"triggers"`
}

// NewMktemp registers a new resource with the given unique name, arguments, and options.
func NewMktemp(ctx *pulumi.Context,
	name string, args *MktempArgs, opts ...pulumi.ResourceOption) (*Mktemp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Mktemp
	err := ctx.RegisterResource("baremetal:coreutils:Mktemp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMktemp gets an existing Mktemp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMktemp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MktempState, opts ...pulumi.ResourceOption) (*Mktemp, error) {
	var resource Mktemp
	err := ctx.ReadResource("baremetal:coreutils:Mktemp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mktemp resources.
type mktempState struct {
}

type MktempState struct {
}

func (MktempState) ElementType() reflect.Type {
	return reflect.TypeOf((*mktempState)(nil)).Elem()
}

type mktempArgs struct {
	Args         MktempArgsType                `pulumi:"args"`
	Connection   *config.ProvisionerConnection `pulumi:"connection"`
	CustomDelete []string                      `pulumi:"customDelete"`
	CustomUpdate []string                      `pulumi:"customUpdate"`
	Triggers     []interface{}                 `pulumi:"triggers"`
}

// The set of arguments for constructing a Mktemp resource.
type MktempArgs struct {
	Args         pulumix.Input[*MktempArgsTypeArgs]
	Connection   pulumix.Input[*config.ProvisionerConnectionArgs]
	CustomDelete pulumix.Input[[]string]
	CustomUpdate pulumix.Input[[]string]
	Triggers     pulumix.Input[[]any]
}

func (MktempArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mktempArgs)(nil)).Elem()
}

type MktempOutput struct{ *pulumi.OutputState }

func (MktempOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Mktemp)(nil)).Elem()
}

func (o MktempOutput) ToMktempOutput() MktempOutput {
	return o
}

func (o MktempOutput) ToMktempOutputWithContext(ctx context.Context) MktempOutput {
	return o
}

func (o MktempOutput) ToOutput(ctx context.Context) pulumix.Output[Mktemp] {
	return pulumix.Output[Mktemp]{
		OutputState: o.OutputState,
	}
}

func (o MktempOutput) Args() pulumix.GPtrOutput[MktempArgsType, MktempArgsTypeOutput] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.GPtrOutput[MktempArgsType, MktempArgsTypeOutput] { return v.Args })
	unwrapped := pulumix.Flatten[*MktempArgsType, pulumix.GPtrOutput[MktempArgsType, MktempArgsTypeOutput]](value)
	return pulumix.GPtrOutput[MktempArgsType, MktempArgsTypeOutput]{OutputState: unwrapped.OutputState}
}

func (o MktempOutput) Connection() pulumix.GPtrOutput[config.ProvisionerConnection, config.ProvisionerConnectionOutput] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.GPtrOutput[config.ProvisionerConnection, config.ProvisionerConnectionOutput] {
		return v.Connection
	})
	unwrapped := pulumix.Flatten[*config.ProvisionerConnection, pulumix.GPtrOutput[config.ProvisionerConnection, config.ProvisionerConnectionOutput]](value)
	return pulumix.GPtrOutput[config.ProvisionerConnection, config.ProvisionerConnectionOutput]{OutputState: unwrapped.OutputState}
}

func (o MktempOutput) CreatedFiles() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.ArrayOutput[string] { return v.CreatedFiles })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o MktempOutput) CustomDelete() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.ArrayOutput[string] { return v.CustomDelete })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o MktempOutput) CustomUpdate() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.ArrayOutput[string] { return v.CustomUpdate })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o MktempOutput) ExitCode() pulumix.Output[int] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.Output[int] { return v.ExitCode })
	return pulumix.Flatten[int, pulumix.Output[int]](value)
}

func (o MktempOutput) MovedFiles() pulumix.MapOutput[string] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.MapOutput[string] { return v.MovedFiles })
	unwrapped := pulumix.Flatten[map[string]string, pulumix.MapOutput[string]](value)
	return pulumix.MapOutput[string]{OutputState: unwrapped.OutputState}
}

func (o MktempOutput) Stderr() pulumix.Output[string] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.Output[string] { return v.Stderr })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o MktempOutput) Stdout() pulumix.Output[string] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.Output[string] { return v.Stdout })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o MktempOutput) Triggers() pulumix.ArrayOutput[any] {
	value := pulumix.Apply[Mktemp](o, func(v Mktemp) pulumix.ArrayOutput[any] { return v.Triggers })
	unwrapped := pulumix.Flatten[[]interface{}, pulumix.ArrayOutput[any]](value)
	return pulumix.ArrayOutput[any]{OutputState: unwrapped.OutputState}
}

func init() {
	pulumi.RegisterOutputType(MktempOutput{})
}
