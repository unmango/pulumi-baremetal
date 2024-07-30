// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cmd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

var _ = internal.GetEnvOrDefault

type TeeOpts struct {
	Files []string `pulumi:"files"`
}

type TeeOptsArgs struct {
	Files pulumix.Input[[]string] `pulumi:"files"`
}

func (TeeOptsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TeeOpts)(nil)).Elem()
}

func (i TeeOptsArgs) ToTeeOptsOutput() TeeOptsOutput {
	return i.ToTeeOptsOutputWithContext(context.Background())
}

func (i TeeOptsArgs) ToTeeOptsOutputWithContext(ctx context.Context) TeeOptsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeeOptsOutput)
}

func (i *TeeOptsArgs) ToOutput(ctx context.Context) pulumix.Output[*TeeOptsArgs] {
	return pulumix.Val(i)
}

type TeeOptsOutput struct{ *pulumi.OutputState }

func (TeeOptsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TeeOpts)(nil)).Elem()
}

func (o TeeOptsOutput) ToTeeOptsOutput() TeeOptsOutput {
	return o
}

func (o TeeOptsOutput) ToTeeOptsOutputWithContext(ctx context.Context) TeeOptsOutput {
	return o
}

func (o TeeOptsOutput) ToOutput(ctx context.Context) pulumix.Output[TeeOpts] {
	return pulumix.Output[TeeOpts]{
		OutputState: o.OutputState,
	}
}

func (o TeeOptsOutput) Files() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[TeeOpts](o, func(v TeeOpts) []string { return v.Files })
	return pulumix.ArrayOutput[string]{OutputState: value.OutputState}
}

func init() {
	pulumi.RegisterOutputType(TeeOptsOutput{})
}
