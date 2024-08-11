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

type Mkdir struct {
	pulumi.CustomResourceState

	Args         MkdirArgsTypeOutput      `pulumi:"args"`
	CreatedFiles pulumi.StringArrayOutput `pulumi:"createdFiles"`
	CustomDelete pulumi.StringArrayOutput `pulumi:"customDelete"`
	CustomUpdate pulumi.StringArrayOutput `pulumi:"customUpdate"`
	ExitCode     pulumi.IntOutput         `pulumi:"exitCode"`
	MovedFiles   pulumi.StringMapOutput   `pulumi:"movedFiles"`
	Stderr       pulumi.StringOutput      `pulumi:"stderr"`
	Stdout       pulumi.StringOutput      `pulumi:"stdout"`
	Triggers     pulumi.ArrayOutput       `pulumi:"triggers"`
}

// NewMkdir registers a new resource with the given unique name, arguments, and options.
func NewMkdir(ctx *pulumi.Context,
	name string, args *MkdirArgs, opts ...pulumi.ResourceOption) (*Mkdir, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Mkdir
	err := ctx.RegisterResource("baremetal:cmd:Mkdir", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMkdir gets an existing Mkdir resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMkdir(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MkdirState, opts ...pulumi.ResourceOption) (*Mkdir, error) {
	var resource Mkdir
	err := ctx.ReadResource("baremetal:cmd:Mkdir", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mkdir resources.
type mkdirState struct {
}

type MkdirState struct {
}

func (MkdirState) ElementType() reflect.Type {
	return reflect.TypeOf((*mkdirState)(nil)).Elem()
}

type mkdirArgs struct {
	Args         MkdirArgsType `pulumi:"args"`
	CustomDelete []string      `pulumi:"customDelete"`
	CustomUpdate []string      `pulumi:"customUpdate"`
	Triggers     []interface{} `pulumi:"triggers"`
}

// The set of arguments for constructing a Mkdir resource.
type MkdirArgs struct {
	Args         MkdirArgsTypeInput
	CustomDelete pulumi.StringArrayInput
	CustomUpdate pulumi.StringArrayInput
	Triggers     pulumi.ArrayInput
}

func (MkdirArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mkdirArgs)(nil)).Elem()
}

type MkdirInput interface {
	pulumi.Input

	ToMkdirOutput() MkdirOutput
	ToMkdirOutputWithContext(ctx context.Context) MkdirOutput
}

func (*Mkdir) ElementType() reflect.Type {
	return reflect.TypeOf((**Mkdir)(nil)).Elem()
}

func (i *Mkdir) ToMkdirOutput() MkdirOutput {
	return i.ToMkdirOutputWithContext(context.Background())
}

func (i *Mkdir) ToMkdirOutputWithContext(ctx context.Context) MkdirOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MkdirOutput)
}

func (i *Mkdir) ToOutput(ctx context.Context) pulumix.Output[*Mkdir] {
	return pulumix.Output[*Mkdir]{
		OutputState: i.ToMkdirOutputWithContext(ctx).OutputState,
	}
}

type MkdirOutput struct{ *pulumi.OutputState }

func (MkdirOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mkdir)(nil)).Elem()
}

func (o MkdirOutput) ToMkdirOutput() MkdirOutput {
	return o
}

func (o MkdirOutput) ToMkdirOutputWithContext(ctx context.Context) MkdirOutput {
	return o
}

func (o MkdirOutput) ToOutput(ctx context.Context) pulumix.Output[*Mkdir] {
	return pulumix.Output[*Mkdir]{
		OutputState: o.OutputState,
	}
}

func (o MkdirOutput) Args() MkdirArgsTypeOutput {
	return o.ApplyT(func(v *Mkdir) MkdirArgsTypeOutput { return v.Args }).(MkdirArgsTypeOutput)
}

func (o MkdirOutput) CreatedFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Mkdir) pulumi.StringArrayOutput { return v.CreatedFiles }).(pulumi.StringArrayOutput)
}

func (o MkdirOutput) CustomDelete() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Mkdir) pulumi.StringArrayOutput { return v.CustomDelete }).(pulumi.StringArrayOutput)
}

func (o MkdirOutput) CustomUpdate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Mkdir) pulumi.StringArrayOutput { return v.CustomUpdate }).(pulumi.StringArrayOutput)
}

func (o MkdirOutput) ExitCode() pulumi.IntOutput {
	return o.ApplyT(func(v *Mkdir) pulumi.IntOutput { return v.ExitCode }).(pulumi.IntOutput)
}

func (o MkdirOutput) MovedFiles() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Mkdir) pulumi.StringMapOutput { return v.MovedFiles }).(pulumi.StringMapOutput)
}

func (o MkdirOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Mkdir) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

func (o MkdirOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Mkdir) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

func (o MkdirOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Mkdir) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MkdirInput)(nil)).Elem(), &Mkdir{})
	pulumi.RegisterOutputType(MkdirOutput{})
}
