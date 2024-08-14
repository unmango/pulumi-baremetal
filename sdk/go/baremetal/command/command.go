// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package command

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

type Command struct {
	pulumi.CustomResourceState

	Args     pulumi.StringArrayOutput `pulumi:"args"`
	ExitCode pulumi.IntOutput         `pulumi:"exitCode"`
	Stderr   pulumi.StringOutput      `pulumi:"stderr"`
	Stdout   pulumi.StringOutput      `pulumi:"stdout"`
	Triggers pulumi.ArrayOutput       `pulumi:"triggers"`
}

// NewCommand registers a new resource with the given unique name, arguments, and options.
func NewCommand(ctx *pulumi.Context,
	name string, args *CommandArgs, opts ...pulumi.ResourceOption) (*Command, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Command
	err := ctx.RegisterResource("baremetal:command:Command", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommand gets an existing Command resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommandState, opts ...pulumi.ResourceOption) (*Command, error) {
	var resource Command
	err := ctx.ReadResource("baremetal:command:Command", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Command resources.
type commandState struct {
}

type CommandState struct {
}

func (CommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*commandState)(nil)).Elem()
}

type commandArgs struct {
	Args     []string      `pulumi:"args"`
	Triggers []interface{} `pulumi:"triggers"`
}

// The set of arguments for constructing a Command resource.
type CommandArgs struct {
	Args     pulumi.StringArrayInput
	Triggers pulumi.ArrayInput
}

func (CommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commandArgs)(nil)).Elem()
}

type CommandInput interface {
	pulumi.Input

	ToCommandOutput() CommandOutput
	ToCommandOutputWithContext(ctx context.Context) CommandOutput
}

func (*Command) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil)).Elem()
}

func (i *Command) ToCommandOutput() CommandOutput {
	return i.ToCommandOutputWithContext(context.Background())
}

func (i *Command) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandOutput)
}

func (i *Command) ToOutput(ctx context.Context) pulumix.Output[*Command] {
	return pulumix.Output[*Command]{
		OutputState: i.ToCommandOutputWithContext(ctx).OutputState,
	}
}

type CommandOutput struct{ *pulumi.OutputState }

func (CommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil)).Elem()
}

func (o CommandOutput) ToCommandOutput() CommandOutput {
	return o
}

func (o CommandOutput) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return o
}

func (o CommandOutput) ToOutput(ctx context.Context) pulumix.Output[*Command] {
	return pulumix.Output[*Command]{
		OutputState: o.OutputState,
	}
}

func (o CommandOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Command) pulumi.StringArrayOutput { return v.Args }).(pulumi.StringArrayOutput)
}

func (o CommandOutput) ExitCode() pulumi.IntOutput {
	return o.ApplyT(func(v *Command) pulumi.IntOutput { return v.ExitCode }).(pulumi.IntOutput)
}

func (o CommandOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

func (o CommandOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

func (o CommandOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Command) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommandInput)(nil)).Elem(), &Command{})
	pulumi.RegisterOutputType(CommandOutput{})
}
