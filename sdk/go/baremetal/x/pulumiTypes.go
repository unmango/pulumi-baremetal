// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package baremetal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

var _ = internal.GetEnvOrDefault

type ProvisionerConnection struct {
	Address string  `pulumi:"address"`
	CaPem   *string `pulumi:"caPem"`
	CertPem *string `pulumi:"certPem"`
	KeyPem  *string `pulumi:"keyPem"`
	Port    *string `pulumi:"port"`
}

type ProvisionerConnectionArgs struct {
	Address pulumix.Input[string]  `pulumi:"address"`
	CaPem   pulumix.Input[*string] `pulumi:"caPem"`
	CertPem pulumix.Input[*string] `pulumi:"certPem"`
	KeyPem  pulumix.Input[*string] `pulumi:"keyPem"`
	Port    pulumix.Input[*string] `pulumi:"port"`
}

func (ProvisionerConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionerConnection)(nil)).Elem()
}

func (i ProvisionerConnectionArgs) ToProvisionerConnectionOutput() ProvisionerConnectionOutput {
	return i.ToProvisionerConnectionOutputWithContext(context.Background())
}

func (i ProvisionerConnectionArgs) ToProvisionerConnectionOutputWithContext(ctx context.Context) ProvisionerConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionerConnectionOutput)
}

func (i *ProvisionerConnectionArgs) ToOutput(ctx context.Context) pulumix.Output[*ProvisionerConnectionArgs] {
	return pulumix.Val(i)
}

type ProvisionerConnectionOutput struct{ *pulumi.OutputState }

func (ProvisionerConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisionerConnection)(nil)).Elem()
}

func (o ProvisionerConnectionOutput) ToProvisionerConnectionOutput() ProvisionerConnectionOutput {
	return o
}

func (o ProvisionerConnectionOutput) ToProvisionerConnectionOutputWithContext(ctx context.Context) ProvisionerConnectionOutput {
	return o
}

func (o ProvisionerConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[ProvisionerConnection] {
	return pulumix.Output[ProvisionerConnection]{
		OutputState: o.OutputState,
	}
}

func (o ProvisionerConnectionOutput) Address() pulumix.Output[string] {
	return pulumix.Apply[ProvisionerConnection](o, func(v ProvisionerConnection) string { return v.Address })
}

func (o ProvisionerConnectionOutput) CaPem() pulumix.Output[*string] {
	return pulumix.Apply[ProvisionerConnection](o, func(v ProvisionerConnection) *string { return v.CaPem })
}

func (o ProvisionerConnectionOutput) CertPem() pulumix.Output[*string] {
	return pulumix.Apply[ProvisionerConnection](o, func(v ProvisionerConnection) *string { return v.CertPem })
}

func (o ProvisionerConnectionOutput) KeyPem() pulumix.Output[*string] {
	return pulumix.Apply[ProvisionerConnection](o, func(v ProvisionerConnection) *string { return v.KeyPem })
}

func (o ProvisionerConnectionOutput) Port() pulumix.Output[*string] {
	return pulumix.Apply[ProvisionerConnection](o, func(v ProvisionerConnection) *string { return v.Port })
}

func init() {
	pulumi.RegisterOutputType(ProvisionerConnectionOutput{})
}
