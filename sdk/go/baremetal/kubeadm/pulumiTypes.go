// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubeadm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

var _ = internal.GetEnvOrDefault

type KubeadmArgsType struct {
	Commands []string `pulumi:"commands"`
}

// KubeadmArgsTypeInput is an input type that accepts KubeadmArgsTypeArgs and KubeadmArgsTypeOutput values.
// You can construct a concrete instance of `KubeadmArgsTypeInput` via:
//
//	KubeadmArgsTypeArgs{...}
type KubeadmArgsTypeInput interface {
	pulumi.Input

	ToKubeadmArgsTypeOutput() KubeadmArgsTypeOutput
	ToKubeadmArgsTypeOutputWithContext(context.Context) KubeadmArgsTypeOutput
}

type KubeadmArgsTypeArgs struct {
	Commands pulumi.StringArrayInput `pulumi:"commands"`
}

func (KubeadmArgsTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeadmArgsType)(nil)).Elem()
}

func (i KubeadmArgsTypeArgs) ToKubeadmArgsTypeOutput() KubeadmArgsTypeOutput {
	return i.ToKubeadmArgsTypeOutputWithContext(context.Background())
}

func (i KubeadmArgsTypeArgs) ToKubeadmArgsTypeOutputWithContext(ctx context.Context) KubeadmArgsTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeadmArgsTypeOutput)
}

func (i KubeadmArgsTypeArgs) ToOutput(ctx context.Context) pulumix.Output[KubeadmArgsType] {
	return pulumix.Output[KubeadmArgsType]{
		OutputState: i.ToKubeadmArgsTypeOutputWithContext(ctx).OutputState,
	}
}

type KubeadmArgsTypeOutput struct{ *pulumi.OutputState }

func (KubeadmArgsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeadmArgsType)(nil)).Elem()
}

func (o KubeadmArgsTypeOutput) ToKubeadmArgsTypeOutput() KubeadmArgsTypeOutput {
	return o
}

func (o KubeadmArgsTypeOutput) ToKubeadmArgsTypeOutputWithContext(ctx context.Context) KubeadmArgsTypeOutput {
	return o
}

func (o KubeadmArgsTypeOutput) ToOutput(ctx context.Context) pulumix.Output[KubeadmArgsType] {
	return pulumix.Output[KubeadmArgsType]{
		OutputState: o.OutputState,
	}
}

func (o KubeadmArgsTypeOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KubeadmArgsType) []string { return v.Commands }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeadmArgsTypeInput)(nil)).Elem(), KubeadmArgsTypeArgs{})
	pulumi.RegisterOutputType(KubeadmArgsTypeOutput{})
}
