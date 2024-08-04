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

type Mv struct {
	pulumi.CustomResourceState

	Args         MvArgsTypeOutput         `pulumi:"args"`
	CreatedFiles pulumi.StringArrayOutput `pulumi:"createdFiles"`
	ExitCode     pulumi.IntOutput         `pulumi:"exitCode"`
	MovedFiles   pulumi.StringMapOutput   `pulumi:"movedFiles"`
	Stderr       pulumi.StringOutput      `pulumi:"stderr"`
	Stdout       pulumi.StringOutput      `pulumi:"stdout"`
}

// NewMv registers a new resource with the given unique name, arguments, and options.
func NewMv(ctx *pulumi.Context,
	name string, args *MvArgs, opts ...pulumi.ResourceOption) (*Mv, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Mv
	err := ctx.RegisterResource("baremetal:cmd:Mv", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMv gets an existing Mv resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMv(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MvState, opts ...pulumi.ResourceOption) (*Mv, error) {
	var resource Mv
	err := ctx.ReadResource("baremetal:cmd:Mv", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mv resources.
type mvState struct {
}

type MvState struct {
}

func (MvState) ElementType() reflect.Type {
	return reflect.TypeOf((*mvState)(nil)).Elem()
}

type mvArgs struct {
	Backup               *string  `pulumi:"backup"`
	Destination          *string  `pulumi:"destination"`
	Directory            *string  `pulumi:"directory"`
	Force                *bool    `pulumi:"force"`
	Help                 *bool    `pulumi:"help"`
	NoClobber            *bool    `pulumi:"noClobber"`
	NoTargetDirectory    *bool    `pulumi:"noTargetDirectory"`
	Source               []string `pulumi:"source"`
	StripTrailingSlashes *bool    `pulumi:"stripTrailingSlashes"`
	Suffix               *string  `pulumi:"suffix"`
	TargetDirectory      *string  `pulumi:"targetDirectory"`
	Update               *bool    `pulumi:"update"`
	Verbose              *bool    `pulumi:"verbose"`
	Version              *bool    `pulumi:"version"`
}

// The set of arguments for constructing a Mv resource.
type MvArgs struct {
	Backup               pulumi.StringPtrInput
	Destination          pulumi.StringPtrInput
	Directory            pulumi.StringPtrInput
	Force                pulumi.BoolPtrInput
	Help                 pulumi.BoolPtrInput
	NoClobber            pulumi.BoolPtrInput
	NoTargetDirectory    pulumi.BoolPtrInput
	Source               pulumi.StringArrayInput
	StripTrailingSlashes pulumi.BoolPtrInput
	Suffix               pulumi.StringPtrInput
	TargetDirectory      pulumi.StringPtrInput
	Update               pulumi.BoolPtrInput
	Verbose              pulumi.BoolPtrInput
	Version              pulumi.BoolPtrInput
}

func (MvArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mvArgs)(nil)).Elem()
}

type MvInput interface {
	pulumi.Input

	ToMvOutput() MvOutput
	ToMvOutputWithContext(ctx context.Context) MvOutput
}

func (*Mv) ElementType() reflect.Type {
	return reflect.TypeOf((**Mv)(nil)).Elem()
}

func (i *Mv) ToMvOutput() MvOutput {
	return i.ToMvOutputWithContext(context.Background())
}

func (i *Mv) ToMvOutputWithContext(ctx context.Context) MvOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MvOutput)
}

func (i *Mv) ToOutput(ctx context.Context) pulumix.Output[*Mv] {
	return pulumix.Output[*Mv]{
		OutputState: i.ToMvOutputWithContext(ctx).OutputState,
	}
}

type MvOutput struct{ *pulumi.OutputState }

func (MvOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mv)(nil)).Elem()
}

func (o MvOutput) ToMvOutput() MvOutput {
	return o
}

func (o MvOutput) ToMvOutputWithContext(ctx context.Context) MvOutput {
	return o
}

func (o MvOutput) ToOutput(ctx context.Context) pulumix.Output[*Mv] {
	return pulumix.Output[*Mv]{
		OutputState: o.OutputState,
	}
}

func (o MvOutput) Args() MvArgsTypeOutput {
	return o.ApplyT(func(v *Mv) MvArgsTypeOutput { return v.Args }).(MvArgsTypeOutput)
}

func (o MvOutput) CreatedFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringArrayOutput { return v.CreatedFiles }).(pulumi.StringArrayOutput)
}

func (o MvOutput) ExitCode() pulumi.IntOutput {
	return o.ApplyT(func(v *Mv) pulumi.IntOutput { return v.ExitCode }).(pulumi.IntOutput)
}

func (o MvOutput) MovedFiles() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringMapOutput { return v.MovedFiles }).(pulumi.StringMapOutput)
}

func (o MvOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

func (o MvOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MvInput)(nil)).Elem(), &Mv{})
	pulumi.RegisterOutputType(MvOutput{})
}
