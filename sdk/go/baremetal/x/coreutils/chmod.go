// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package coreutils

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

type Chmod struct {
	pulumi.CustomResourceState

	Args         pulumix.GPtrOutput[ChmodArgsType, ChmodArgsTypeOutput]                                     `pulumi:"args"`
	Connection   pulumix.GPtrOutput[baremetal.ProvisionerConnection, baremetal.ProvisionerConnectionOutput] `pulumi:"connection"`
	CreatedFiles pulumix.ArrayOutput[string]                                                                `pulumi:"createdFiles"`
	CustomDelete pulumix.ArrayOutput[string]                                                                `pulumi:"customDelete"`
	CustomUpdate pulumix.ArrayOutput[string]                                                                `pulumi:"customUpdate"`
	ExitCode     pulumix.Output[int]                                                                        `pulumi:"exitCode"`
	MovedFiles   pulumix.MapOutput[string]                                                                  `pulumi:"movedFiles"`
	Stderr       pulumix.Output[string]                                                                     `pulumi:"stderr"`
	Stdout       pulumix.Output[string]                                                                     `pulumi:"stdout"`
	Triggers     pulumix.ArrayOutput[any]                                                                   `pulumi:"triggers"`
}

// NewChmod registers a new resource with the given unique name, arguments, and options.
func NewChmod(ctx *pulumi.Context,
	name string, args *ChmodArgs, opts ...pulumi.ResourceOption) (*Chmod, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Chmod
	err := ctx.RegisterResource("baremetal:coreutils:Chmod", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChmod gets an existing Chmod resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChmod(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChmodState, opts ...pulumi.ResourceOption) (*Chmod, error) {
	var resource Chmod
	err := ctx.ReadResource("baremetal:coreutils:Chmod", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Chmod resources.
type chmodState struct {
}

type ChmodState struct {
}

func (ChmodState) ElementType() reflect.Type {
	return reflect.TypeOf((*chmodState)(nil)).Elem()
}

type chmodArgs struct {
	Args         ChmodArgsType                    `pulumi:"args"`
	Connection   *baremetal.ProvisionerConnection `pulumi:"connection"`
	CustomDelete []string                         `pulumi:"customDelete"`
	CustomUpdate []string                         `pulumi:"customUpdate"`
	Triggers     []interface{}                    `pulumi:"triggers"`
}

// The set of arguments for constructing a Chmod resource.
type ChmodArgs struct {
	Args         pulumix.Input[*ChmodArgsTypeArgs]
	Connection   pulumix.Input[*baremetal.ProvisionerConnectionArgs]
	CustomDelete pulumix.Input[[]string]
	CustomUpdate pulumix.Input[[]string]
	Triggers     pulumix.Input[[]any]
}

func (ChmodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chmodArgs)(nil)).Elem()
}

type ChmodOutput struct{ *pulumi.OutputState }

func (ChmodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Chmod)(nil)).Elem()
}

func (o ChmodOutput) ToChmodOutput() ChmodOutput {
	return o
}

func (o ChmodOutput) ToChmodOutputWithContext(ctx context.Context) ChmodOutput {
	return o
}

func (o ChmodOutput) ToOutput(ctx context.Context) pulumix.Output[Chmod] {
	return pulumix.Output[Chmod]{
		OutputState: o.OutputState,
	}
}

func (o ChmodOutput) Args() pulumix.GPtrOutput[ChmodArgsType, ChmodArgsTypeOutput] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.GPtrOutput[ChmodArgsType, ChmodArgsTypeOutput] { return v.Args })
	unwrapped := pulumix.Flatten[*ChmodArgsType, pulumix.GPtrOutput[ChmodArgsType, ChmodArgsTypeOutput]](value)
	return pulumix.GPtrOutput[ChmodArgsType, ChmodArgsTypeOutput]{OutputState: unwrapped.OutputState}
}

func (o ChmodOutput) Connection() pulumix.GPtrOutput[baremetal.ProvisionerConnection, baremetal.ProvisionerConnectionOutput] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.GPtrOutput[baremetal.ProvisionerConnection, baremetal.ProvisionerConnectionOutput] {
		return v.Connection
	})
	unwrapped := pulumix.Flatten[*baremetal.ProvisionerConnection, pulumix.GPtrOutput[baremetal.ProvisionerConnection, baremetal.ProvisionerConnectionOutput]](value)
	return pulumix.GPtrOutput[baremetal.ProvisionerConnection, baremetal.ProvisionerConnectionOutput]{OutputState: unwrapped.OutputState}
}

func (o ChmodOutput) CreatedFiles() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.ArrayOutput[string] { return v.CreatedFiles })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o ChmodOutput) CustomDelete() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.ArrayOutput[string] { return v.CustomDelete })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o ChmodOutput) CustomUpdate() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.ArrayOutput[string] { return v.CustomUpdate })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o ChmodOutput) ExitCode() pulumix.Output[int] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.Output[int] { return v.ExitCode })
	return pulumix.Flatten[int, pulumix.Output[int]](value)
}

func (o ChmodOutput) MovedFiles() pulumix.MapOutput[string] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.MapOutput[string] { return v.MovedFiles })
	unwrapped := pulumix.Flatten[map[string]string, pulumix.MapOutput[string]](value)
	return pulumix.MapOutput[string]{OutputState: unwrapped.OutputState}
}

func (o ChmodOutput) Stderr() pulumix.Output[string] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.Output[string] { return v.Stderr })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o ChmodOutput) Stdout() pulumix.Output[string] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.Output[string] { return v.Stdout })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o ChmodOutput) Triggers() pulumix.ArrayOutput[any] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.ArrayOutput[any] { return v.Triggers })
	unwrapped := pulumix.Flatten[[]interface{}, pulumix.ArrayOutput[any]](value)
	return pulumix.ArrayOutput[any]{OutputState: unwrapped.OutputState}
}

func init() {
	pulumi.RegisterOutputType(ChmodOutput{})
}
