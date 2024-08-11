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

type Tar struct {
	pulumi.CustomResourceState

	Args         TarArgsTypeOutput        `pulumi:"args"`
	CreatedFiles pulumi.StringArrayOutput `pulumi:"createdFiles"`
	ExitCode     pulumi.IntOutput         `pulumi:"exitCode"`
	MovedFiles   pulumi.StringMapOutput   `pulumi:"movedFiles"`
	Stderr       pulumi.StringOutput      `pulumi:"stderr"`
	Stdout       pulumi.StringOutput      `pulumi:"stdout"`
	Triggers     pulumi.ArrayOutput       `pulumi:"triggers"`
}

// NewTar registers a new resource with the given unique name, arguments, and options.
func NewTar(ctx *pulumi.Context,
	name string, args *TarArgs, opts ...pulumi.ResourceOption) (*Tar, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Tar
	err := ctx.RegisterResource("baremetal:cmd:Tar", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTar gets an existing Tar resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTar(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TarState, opts ...pulumi.ResourceOption) (*Tar, error) {
	var resource Tar
	err := ctx.ReadResource("baremetal:cmd:Tar", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tar resources.
type tarState struct {
}

type TarState struct {
}

func (TarState) ElementType() reflect.Type {
	return reflect.TypeOf((*tarState)(nil)).Elem()
}

type tarArgs struct {
	Args     TarArgsType   `pulumi:"args"`
	Triggers []interface{} `pulumi:"triggers"`
}

// The set of arguments for constructing a Tar resource.
type TarArgs struct {
	Args     TarArgsTypeInput
	Triggers pulumi.ArrayInput
}

func (TarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tarArgs)(nil)).Elem()
}

type TarInput interface {
	pulumi.Input

	ToTarOutput() TarOutput
	ToTarOutputWithContext(ctx context.Context) TarOutput
}

func (*Tar) ElementType() reflect.Type {
	return reflect.TypeOf((**Tar)(nil)).Elem()
}

func (i *Tar) ToTarOutput() TarOutput {
	return i.ToTarOutputWithContext(context.Background())
}

func (i *Tar) ToTarOutputWithContext(ctx context.Context) TarOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TarOutput)
}

func (i *Tar) ToOutput(ctx context.Context) pulumix.Output[*Tar] {
	return pulumix.Output[*Tar]{
		OutputState: i.ToTarOutputWithContext(ctx).OutputState,
	}
}

type TarOutput struct{ *pulumi.OutputState }

func (TarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tar)(nil)).Elem()
}

func (o TarOutput) ToTarOutput() TarOutput {
	return o
}

func (o TarOutput) ToTarOutputWithContext(ctx context.Context) TarOutput {
	return o
}

func (o TarOutput) ToOutput(ctx context.Context) pulumix.Output[*Tar] {
	return pulumix.Output[*Tar]{
		OutputState: o.OutputState,
	}
}

func (o TarOutput) Args() TarArgsTypeOutput {
	return o.ApplyT(func(v *Tar) TarArgsTypeOutput { return v.Args }).(TarArgsTypeOutput)
}

func (o TarOutput) CreatedFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Tar) pulumi.StringArrayOutput { return v.CreatedFiles }).(pulumi.StringArrayOutput)
}

func (o TarOutput) ExitCode() pulumi.IntOutput {
	return o.ApplyT(func(v *Tar) pulumi.IntOutput { return v.ExitCode }).(pulumi.IntOutput)
}

func (o TarOutput) MovedFiles() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Tar) pulumi.StringMapOutput { return v.MovedFiles }).(pulumi.StringMapOutput)
}

func (o TarOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Tar) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

func (o TarOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Tar) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

func (o TarOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Tar) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TarInput)(nil)).Elem(), &Tar{})
	pulumi.RegisterOutputType(TarOutput{})
}
