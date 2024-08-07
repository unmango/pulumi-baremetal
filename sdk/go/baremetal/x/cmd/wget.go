// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cmd

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

type Wget struct {
	pulumi.CustomResourceState

	Args         pulumix.GPtrOutput[WgetArgsType, WgetArgsTypeOutput] `pulumi:"args"`
	CreatedFiles pulumix.ArrayOutput[string]                          `pulumi:"createdFiles"`
	ExitCode     pulumix.Output[int]                                  `pulumi:"exitCode"`
	MovedFiles   pulumix.MapOutput[string]                            `pulumi:"movedFiles"`
	Stderr       pulumix.Output[string]                               `pulumi:"stderr"`
	Stdout       pulumix.Output[string]                               `pulumi:"stdout"`
	Triggers     pulumix.ArrayOutput[any]                             `pulumi:"triggers"`
}

// NewWget registers a new resource with the given unique name, arguments, and options.
func NewWget(ctx *pulumi.Context,
	name string, args *WgetArgs, opts ...pulumi.ResourceOption) (*Wget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Wget
	err := ctx.RegisterResource("baremetal:cmd:Wget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWget gets an existing Wget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WgetState, opts ...pulumi.ResourceOption) (*Wget, error) {
	var resource Wget
	err := ctx.ReadResource("baremetal:cmd:Wget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wget resources.
type wgetState struct {
}

type WgetState struct {
}

func (WgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*wgetState)(nil)).Elem()
}

type wgetArgs struct {
	Args     WgetArgsType  `pulumi:"args"`
	Triggers []interface{} `pulumi:"triggers"`
}

// The set of arguments for constructing a Wget resource.
type WgetArgs struct {
	Args     pulumix.Input[*WgetArgsTypeArgs]
	Triggers pulumix.Input[[]any]
}

func (WgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wgetArgs)(nil)).Elem()
}

type WgetOutput struct{ *pulumi.OutputState }

func (WgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Wget)(nil)).Elem()
}

func (o WgetOutput) ToWgetOutput() WgetOutput {
	return o
}

func (o WgetOutput) ToWgetOutputWithContext(ctx context.Context) WgetOutput {
	return o
}

func (o WgetOutput) ToOutput(ctx context.Context) pulumix.Output[Wget] {
	return pulumix.Output[Wget]{
		OutputState: o.OutputState,
	}
}

func (o WgetOutput) Args() pulumix.GPtrOutput[WgetArgsType, WgetArgsTypeOutput] {
	value := pulumix.Apply[Wget](o, func(v Wget) pulumix.GPtrOutput[WgetArgsType, WgetArgsTypeOutput] { return v.Args })
	unwrapped := pulumix.Flatten[*WgetArgsType, pulumix.GPtrOutput[WgetArgsType, WgetArgsTypeOutput]](value)
	return pulumix.GPtrOutput[WgetArgsType, WgetArgsTypeOutput]{OutputState: unwrapped.OutputState}
}

func (o WgetOutput) CreatedFiles() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Wget](o, func(v Wget) pulumix.ArrayOutput[string] { return v.CreatedFiles })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o WgetOutput) ExitCode() pulumix.Output[int] {
	value := pulumix.Apply[Wget](o, func(v Wget) pulumix.Output[int] { return v.ExitCode })
	return pulumix.Flatten[int, pulumix.Output[int]](value)
}

func (o WgetOutput) MovedFiles() pulumix.MapOutput[string] {
	value := pulumix.Apply[Wget](o, func(v Wget) pulumix.MapOutput[string] { return v.MovedFiles })
	unwrapped := pulumix.Flatten[map[string]string, pulumix.MapOutput[string]](value)
	return pulumix.MapOutput[string]{OutputState: unwrapped.OutputState}
}

func (o WgetOutput) Stderr() pulumix.Output[string] {
	value := pulumix.Apply[Wget](o, func(v Wget) pulumix.Output[string] { return v.Stderr })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o WgetOutput) Stdout() pulumix.Output[string] {
	value := pulumix.Apply[Wget](o, func(v Wget) pulumix.Output[string] { return v.Stdout })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o WgetOutput) Triggers() pulumix.ArrayOutput[any] {
	value := pulumix.Apply[Wget](o, func(v Wget) pulumix.ArrayOutput[any] { return v.Triggers })
	unwrapped := pulumix.Flatten[[]interface{}, pulumix.ArrayOutput[any]](value)
	return pulumix.ArrayOutput[any]{OutputState: unwrapped.OutputState}
}

func init() {
	pulumi.RegisterOutputType(WgetOutput{})
}
