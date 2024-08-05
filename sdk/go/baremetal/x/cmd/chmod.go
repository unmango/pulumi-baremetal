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

type Chmod struct {
	pulumi.CustomResourceState

	Args         pulumix.GPtrOutput[ChmodArgsType, ChmodArgsTypeOutput] `pulumi:"args"`
	CreatedFiles pulumix.ArrayOutput[string]                            `pulumi:"createdFiles"`
	ExitCode     pulumix.Output[int]                                    `pulumi:"exitCode"`
	MovedFiles   pulumix.MapOutput[string]                              `pulumi:"movedFiles"`
	Stderr       pulumix.Output[string]                                 `pulumi:"stderr"`
	Stdout       pulumix.Output[string]                                 `pulumi:"stdout"`
}

// NewChmod registers a new resource with the given unique name, arguments, and options.
func NewChmod(ctx *pulumi.Context,
	name string, args *ChmodArgs, opts ...pulumi.ResourceOption) (*Chmod, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Files == nil {
		return nil, errors.New("invalid value for required argument 'Files'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Chmod
	err := ctx.RegisterResource("baremetal:cmd:Chmod", name, args, &resource, opts...)
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
	err := ctx.ReadResource("baremetal:cmd:Chmod", name, id, state, &resource, opts...)
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
	Changes        *bool    `pulumi:"changes"`
	Files          []string `pulumi:"files"`
	Help           *bool    `pulumi:"help"`
	Mode           []string `pulumi:"mode"`
	NoPreserveRoot *bool    `pulumi:"noPreserveRoot"`
	OctalMode      *string  `pulumi:"octalMode"`
	PreserveRoot   *bool    `pulumi:"preserveRoot"`
	Quiet          *bool    `pulumi:"quiet"`
	Recursive      *bool    `pulumi:"recursive"`
	Reference      *string  `pulumi:"reference"`
	Verbose        *bool    `pulumi:"verbose"`
	Version        *bool    `pulumi:"version"`
}

// The set of arguments for constructing a Chmod resource.
type ChmodArgs struct {
	Changes        pulumix.Input[*bool]
	Files          pulumix.Input[[]string]
	Help           pulumix.Input[*bool]
	Mode           pulumix.Input[[]string]
	NoPreserveRoot pulumix.Input[*bool]
	OctalMode      pulumix.Input[*string]
	PreserveRoot   pulumix.Input[*bool]
	Quiet          pulumix.Input[*bool]
	Recursive      pulumix.Input[*bool]
	Reference      pulumix.Input[*string]
	Verbose        pulumix.Input[*bool]
	Version        pulumix.Input[*bool]
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

func (o ChmodOutput) CreatedFiles() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Chmod](o, func(v Chmod) pulumix.ArrayOutput[string] { return v.CreatedFiles })
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

func init() {
	pulumi.RegisterOutputType(ChmodOutput{})
}
