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

type KubeadmArgsTypeArgs struct {
	Commands pulumix.Input[[]string] `pulumi:"commands"`
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

func (i *KubeadmArgsTypeArgs) ToOutput(ctx context.Context) pulumix.Output[*KubeadmArgsTypeArgs] {
	return pulumix.Val(i)
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

func (o KubeadmArgsTypeOutput) Commands() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[KubeadmArgsType](o, func(v KubeadmArgsType) []string { return v.Commands })
	return pulumix.ArrayOutput[string]{OutputState: value.OutputState}
}

func init() {
	pulumi.RegisterOutputType(KubeadmArgsTypeOutput{})
}