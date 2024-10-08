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

// ProvisionerConnectionInput is an input type that accepts ProvisionerConnectionArgs and ProvisionerConnectionOutput values.
// You can construct a concrete instance of `ProvisionerConnectionInput` via:
//
//	ProvisionerConnectionArgs{...}
type ProvisionerConnectionInput interface {
	pulumi.Input

	ToProvisionerConnectionOutput() ProvisionerConnectionOutput
	ToProvisionerConnectionOutputWithContext(context.Context) ProvisionerConnectionOutput
}

type ProvisionerConnectionArgs struct {
	Address pulumi.StringInput    `pulumi:"address"`
	CaPem   pulumi.StringPtrInput `pulumi:"caPem"`
	CertPem pulumi.StringPtrInput `pulumi:"certPem"`
	KeyPem  pulumi.StringPtrInput `pulumi:"keyPem"`
	Port    pulumi.StringPtrInput `pulumi:"port"`
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

func (i ProvisionerConnectionArgs) ToOutput(ctx context.Context) pulumix.Output[ProvisionerConnection] {
	return pulumix.Output[ProvisionerConnection]{
		OutputState: i.ToProvisionerConnectionOutputWithContext(ctx).OutputState,
	}
}

func (i ProvisionerConnectionArgs) ToProvisionerConnectionPtrOutput() ProvisionerConnectionPtrOutput {
	return i.ToProvisionerConnectionPtrOutputWithContext(context.Background())
}

func (i ProvisionerConnectionArgs) ToProvisionerConnectionPtrOutputWithContext(ctx context.Context) ProvisionerConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionerConnectionOutput).ToProvisionerConnectionPtrOutputWithContext(ctx)
}

// ProvisionerConnectionPtrInput is an input type that accepts ProvisionerConnectionArgs, ProvisionerConnectionPtr and ProvisionerConnectionPtrOutput values.
// You can construct a concrete instance of `ProvisionerConnectionPtrInput` via:
//
//	        ProvisionerConnectionArgs{...}
//
//	or:
//
//	        nil
type ProvisionerConnectionPtrInput interface {
	pulumi.Input

	ToProvisionerConnectionPtrOutput() ProvisionerConnectionPtrOutput
	ToProvisionerConnectionPtrOutputWithContext(context.Context) ProvisionerConnectionPtrOutput
}

type provisionerConnectionPtrType ProvisionerConnectionArgs

func ProvisionerConnectionPtr(v *ProvisionerConnectionArgs) ProvisionerConnectionPtrInput {
	return (*provisionerConnectionPtrType)(v)
}

func (*provisionerConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionerConnection)(nil)).Elem()
}

func (i *provisionerConnectionPtrType) ToProvisionerConnectionPtrOutput() ProvisionerConnectionPtrOutput {
	return i.ToProvisionerConnectionPtrOutputWithContext(context.Background())
}

func (i *provisionerConnectionPtrType) ToProvisionerConnectionPtrOutputWithContext(ctx context.Context) ProvisionerConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionerConnectionPtrOutput)
}

func (i *provisionerConnectionPtrType) ToOutput(ctx context.Context) pulumix.Output[*ProvisionerConnection] {
	return pulumix.Output[*ProvisionerConnection]{
		OutputState: i.ToProvisionerConnectionPtrOutputWithContext(ctx).OutputState,
	}
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

func (o ProvisionerConnectionOutput) ToProvisionerConnectionPtrOutput() ProvisionerConnectionPtrOutput {
	return o.ToProvisionerConnectionPtrOutputWithContext(context.Background())
}

func (o ProvisionerConnectionOutput) ToProvisionerConnectionPtrOutputWithContext(ctx context.Context) ProvisionerConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisionerConnection) *ProvisionerConnection {
		return &v
	}).(ProvisionerConnectionPtrOutput)
}

func (o ProvisionerConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[ProvisionerConnection] {
	return pulumix.Output[ProvisionerConnection]{
		OutputState: o.OutputState,
	}
}

func (o ProvisionerConnectionOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v ProvisionerConnection) string { return v.Address }).(pulumi.StringOutput)
}

func (o ProvisionerConnectionOutput) CaPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionerConnection) *string { return v.CaPem }).(pulumi.StringPtrOutput)
}

func (o ProvisionerConnectionOutput) CertPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionerConnection) *string { return v.CertPem }).(pulumi.StringPtrOutput)
}

func (o ProvisionerConnectionOutput) KeyPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionerConnection) *string { return v.KeyPem }).(pulumi.StringPtrOutput)
}

func (o ProvisionerConnectionOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisionerConnection) *string { return v.Port }).(pulumi.StringPtrOutput)
}

type ProvisionerConnectionPtrOutput struct{ *pulumi.OutputState }

func (ProvisionerConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisionerConnection)(nil)).Elem()
}

func (o ProvisionerConnectionPtrOutput) ToProvisionerConnectionPtrOutput() ProvisionerConnectionPtrOutput {
	return o
}

func (o ProvisionerConnectionPtrOutput) ToProvisionerConnectionPtrOutputWithContext(ctx context.Context) ProvisionerConnectionPtrOutput {
	return o
}

func (o ProvisionerConnectionPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ProvisionerConnection] {
	return pulumix.Output[*ProvisionerConnection]{
		OutputState: o.OutputState,
	}
}

func (o ProvisionerConnectionPtrOutput) Elem() ProvisionerConnectionOutput {
	return o.ApplyT(func(v *ProvisionerConnection) ProvisionerConnection {
		if v != nil {
			return *v
		}
		var ret ProvisionerConnection
		return ret
	}).(ProvisionerConnectionOutput)
}

func (o ProvisionerConnectionPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionerConnection) *string {
		if v == nil {
			return nil
		}
		return &v.Address
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionerConnectionPtrOutput) CaPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionerConnection) *string {
		if v == nil {
			return nil
		}
		return v.CaPem
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionerConnectionPtrOutput) CertPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionerConnection) *string {
		if v == nil {
			return nil
		}
		return v.CertPem
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionerConnectionPtrOutput) KeyPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionerConnection) *string {
		if v == nil {
			return nil
		}
		return v.KeyPem
	}).(pulumi.StringPtrOutput)
}

func (o ProvisionerConnectionPtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisionerConnection) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisionerConnectionInput)(nil)).Elem(), ProvisionerConnectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisionerConnectionPtrInput)(nil)).Elem(), ProvisionerConnectionArgs{})
	pulumi.RegisterOutputType(ProvisionerConnectionOutput{})
	pulumi.RegisterOutputType(ProvisionerConnectionPtrOutput{})
}
