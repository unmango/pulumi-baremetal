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

type Rm struct {
	pulumi.CustomResourceState

	Args         pulumix.GPtrOutput[RmArgsType, RmArgsTypeOutput] `pulumi:"args"`
	CreatedFiles pulumix.ArrayOutput[string]                      `pulumi:"createdFiles"`
	ExitCode     pulumix.Output[int]                              `pulumi:"exitCode"`
	MovedFiles   pulumix.MapOutput[string]                        `pulumi:"movedFiles"`
	Stderr       pulumix.Output[string]                           `pulumi:"stderr"`
	Stdout       pulumix.Output[string]                           `pulumi:"stdout"`
}

// NewRm registers a new resource with the given unique name, arguments, and options.
func NewRm(ctx *pulumi.Context,
	name string, args *RmArgs, opts ...pulumi.ResourceOption) (*Rm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Files == nil {
		return nil, errors.New("invalid value for required argument 'Files'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rm
	err := ctx.RegisterResource("baremetal:cmd:Rm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRm gets an existing Rm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RmState, opts ...pulumi.ResourceOption) (*Rm, error) {
	var resource Rm
	err := ctx.ReadResource("baremetal:cmd:Rm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rm resources.
type rmState struct {
}

type RmState struct {
}

func (RmState) ElementType() reflect.Type {
	return reflect.TypeOf((*rmState)(nil)).Elem()
}

type rmArgs struct {
	Dir           *bool    `pulumi:"dir"`
	Files         []string `pulumi:"files"`
	Force         *bool    `pulumi:"force"`
	Help          *bool    `pulumi:"help"`
	OneFileSystem *bool    `pulumi:"oneFileSystem"`
	Recursive     *bool    `pulumi:"recursive"`
	Verbose       *bool    `pulumi:"verbose"`
}

// The set of arguments for constructing a Rm resource.
type RmArgs struct {
	Dir           pulumix.Input[*bool]
	Files         pulumix.Input[[]string]
	Force         pulumix.Input[*bool]
	Help          pulumix.Input[*bool]
	OneFileSystem pulumix.Input[*bool]
	Recursive     pulumix.Input[*bool]
	Verbose       pulumix.Input[*bool]
}

func (RmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rmArgs)(nil)).Elem()
}

type RmOutput struct{ *pulumi.OutputState }

func (RmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rm)(nil)).Elem()
}

func (o RmOutput) ToRmOutput() RmOutput {
	return o
}

func (o RmOutput) ToRmOutputWithContext(ctx context.Context) RmOutput {
	return o
}

func (o RmOutput) ToOutput(ctx context.Context) pulumix.Output[Rm] {
	return pulumix.Output[Rm]{
		OutputState: o.OutputState,
	}
}

func (o RmOutput) Args() pulumix.GPtrOutput[RmArgsType, RmArgsTypeOutput] {
	value := pulumix.Apply[Rm](o, func(v Rm) pulumix.GPtrOutput[RmArgsType, RmArgsTypeOutput] { return v.Args })
	unwrapped := pulumix.Flatten[*RmArgsType, pulumix.GPtrOutput[RmArgsType, RmArgsTypeOutput]](value)
	return pulumix.GPtrOutput[RmArgsType, RmArgsTypeOutput]{OutputState: unwrapped.OutputState}
}

func (o RmOutput) CreatedFiles() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Rm](o, func(v Rm) pulumix.ArrayOutput[string] { return v.CreatedFiles })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o RmOutput) ExitCode() pulumix.Output[int] {
	value := pulumix.Apply[Rm](o, func(v Rm) pulumix.Output[int] { return v.ExitCode })
	return pulumix.Flatten[int, pulumix.Output[int]](value)
}

func (o RmOutput) MovedFiles() pulumix.MapOutput[string] {
	value := pulumix.Apply[Rm](o, func(v Rm) pulumix.MapOutput[string] { return v.MovedFiles })
	unwrapped := pulumix.Flatten[map[string]string, pulumix.MapOutput[string]](value)
	return pulumix.MapOutput[string]{OutputState: unwrapped.OutputState}
}

func (o RmOutput) Stderr() pulumix.Output[string] {
	value := pulumix.Apply[Rm](o, func(v Rm) pulumix.Output[string] { return v.Stderr })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o RmOutput) Stdout() pulumix.Output[string] {
	value := pulumix.Apply[Rm](o, func(v Rm) pulumix.Output[string] { return v.Stdout })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func init() {
	pulumi.RegisterOutputType(RmOutput{})
}
