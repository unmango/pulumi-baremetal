// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package coreutils

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

type Cat struct {
	pulumi.CustomResourceState

	Args         pulumix.GPtrOutput[CatArgsType, CatArgsTypeOutput] `pulumi:"args"`
	CreatedFiles pulumix.ArrayOutput[string]                        `pulumi:"createdFiles"`
	CustomDelete pulumix.ArrayOutput[string]                        `pulumi:"customDelete"`
	CustomUpdate pulumix.ArrayOutput[string]                        `pulumi:"customUpdate"`
	ExitCode     pulumix.Output[int]                                `pulumi:"exitCode"`
	MovedFiles   pulumix.MapOutput[string]                          `pulumi:"movedFiles"`
	Stderr       pulumix.Output[string]                             `pulumi:"stderr"`
	Stdout       pulumix.Output[string]                             `pulumi:"stdout"`
	Triggers     pulumix.ArrayOutput[any]                           `pulumi:"triggers"`
}

// NewCat registers a new resource with the given unique name, arguments, and options.
func NewCat(ctx *pulumi.Context,
	name string, args *CatArgs, opts ...pulumi.ResourceOption) (*Cat, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cat
	err := ctx.RegisterResource("baremetal:coreutils:Cat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCat gets an existing Cat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatState, opts ...pulumi.ResourceOption) (*Cat, error) {
	var resource Cat
	err := ctx.ReadResource("baremetal:coreutils:Cat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cat resources.
type catState struct {
}

type CatState struct {
}

func (CatState) ElementType() reflect.Type {
	return reflect.TypeOf((*catState)(nil)).Elem()
}

type catArgs struct {
	Args         CatArgsType   `pulumi:"args"`
	CustomDelete []string      `pulumi:"customDelete"`
	CustomUpdate []string      `pulumi:"customUpdate"`
	Triggers     []interface{} `pulumi:"triggers"`
}

// The set of arguments for constructing a Cat resource.
type CatArgs struct {
	Args         pulumix.Input[*CatArgsTypeArgs]
	CustomDelete pulumix.Input[[]string]
	CustomUpdate pulumix.Input[[]string]
	Triggers     pulumix.Input[[]any]
}

func (CatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catArgs)(nil)).Elem()
}

type CatOutput struct{ *pulumi.OutputState }

func (CatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cat)(nil)).Elem()
}

func (o CatOutput) ToCatOutput() CatOutput {
	return o
}

func (o CatOutput) ToCatOutputWithContext(ctx context.Context) CatOutput {
	return o
}

func (o CatOutput) ToOutput(ctx context.Context) pulumix.Output[Cat] {
	return pulumix.Output[Cat]{
		OutputState: o.OutputState,
	}
}

func (o CatOutput) Args() pulumix.GPtrOutput[CatArgsType, CatArgsTypeOutput] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.GPtrOutput[CatArgsType, CatArgsTypeOutput] { return v.Args })
	unwrapped := pulumix.Flatten[*CatArgsType, pulumix.GPtrOutput[CatArgsType, CatArgsTypeOutput]](value)
	return pulumix.GPtrOutput[CatArgsType, CatArgsTypeOutput]{OutputState: unwrapped.OutputState}
}

func (o CatOutput) CreatedFiles() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.ArrayOutput[string] { return v.CreatedFiles })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o CatOutput) CustomDelete() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.ArrayOutput[string] { return v.CustomDelete })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o CatOutput) CustomUpdate() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.ArrayOutput[string] { return v.CustomUpdate })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o CatOutput) ExitCode() pulumix.Output[int] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.Output[int] { return v.ExitCode })
	return pulumix.Flatten[int, pulumix.Output[int]](value)
}

func (o CatOutput) MovedFiles() pulumix.MapOutput[string] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.MapOutput[string] { return v.MovedFiles })
	unwrapped := pulumix.Flatten[map[string]string, pulumix.MapOutput[string]](value)
	return pulumix.MapOutput[string]{OutputState: unwrapped.OutputState}
}

func (o CatOutput) Stderr() pulumix.Output[string] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.Output[string] { return v.Stderr })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o CatOutput) Stdout() pulumix.Output[string] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.Output[string] { return v.Stdout })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o CatOutput) Triggers() pulumix.ArrayOutput[any] {
	value := pulumix.Apply[Cat](o, func(v Cat) pulumix.ArrayOutput[any] { return v.Triggers })
	unwrapped := pulumix.Flatten[[]interface{}, pulumix.ArrayOutput[any]](value)
	return pulumix.ArrayOutput[any]{OutputState: unwrapped.OutputState}
}

func init() {
	pulumi.RegisterOutputType(CatOutput{})
}