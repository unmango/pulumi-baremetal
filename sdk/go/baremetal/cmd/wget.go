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

type Wget struct {
	pulumi.CustomResourceState

	Args         WgetArgsTypeOutput       `pulumi:"args"`
	CreatedFiles pulumi.StringArrayOutput `pulumi:"createdFiles"`
	ExitCode     pulumi.IntOutput         `pulumi:"exitCode"`
	Stderr       pulumi.StringOutput      `pulumi:"stderr"`
	Stdout       pulumi.StringOutput      `pulumi:"stdout"`
}

// NewWget registers a new resource with the given unique name, arguments, and options.
func NewWget(ctx *pulumi.Context,
	name string, args *WgetArgs, opts ...pulumi.ResourceOption) (*Wget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Urls == nil {
		return nil, errors.New("invalid value for required argument 'Urls'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrInput)
	}
	if args.PrivateKeyType != nil {
		args.PrivateKeyType = pulumi.ToSecret(args.PrivateKeyType).(pulumi.StringPtrInput)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Wget
	err := ctx.RegisterResource("baremetal:cmd:Wget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWget gets an existing Wget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WgetState, opts ...pulumi.ResourceOption) (*Wget, error) {
	var resource Wget
	err := ctx.ReadResource("baremetal:cmd:Wget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wget resources.
type wgetState struct {
}

type WgetState struct {
}

func (WgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*wgetState)(nil)).Elem()
}

type wgetArgs struct {
	AppendOutput       *string  `pulumi:"appendOutput"`
	Background         *bool    `pulumi:"background"`
	Base               *string  `pulumi:"base"`
	CaCertificateFile  *string  `pulumi:"caCertificateFile"`
	CaDirectory        *string  `pulumi:"caDirectory"`
	Certificate        *string  `pulumi:"certificate"`
	CertificateType    *string  `pulumi:"certificateType"`
	Config             *string  `pulumi:"config"`
	Continue           *bool    `pulumi:"continue"`
	CrlFile            *string  `pulumi:"crlFile"`
	CutDirs            *int     `pulumi:"cutDirs"`
	Debug              *bool    `pulumi:"debug"`
	DirectoryPrefix    *string  `pulumi:"directoryPrefix"`
	Execute            []string `pulumi:"execute"`
	ForceDirectories   *bool    `pulumi:"forceDirectories"`
	ForceHtml          *bool    `pulumi:"forceHtml"`
	Help               *bool    `pulumi:"help"`
	HttpsOnly          *bool    `pulumi:"httpsOnly"`
	Inet4Only          *bool    `pulumi:"inet4Only"`
	InputFile          *string  `pulumi:"inputFile"`
	KeepSessionCookies *bool    `pulumi:"keepSessionCookies"`
	NoClobber          *bool    `pulumi:"noClobber"`
	NoDirectories      *bool    `pulumi:"noDirectories"`
	NoDnsCache         *bool    `pulumi:"noDnsCache"`
	NoVerbose          *bool    `pulumi:"noVerbose"`
	OutputDocument     *string  `pulumi:"outputDocument"`
	OutputFile         *string  `pulumi:"outputFile"`
	Password           *string  `pulumi:"password"`
	PrivateKey         *string  `pulumi:"privateKey"`
	PrivateKeyType     *string  `pulumi:"privateKeyType"`
	Progress           *string  `pulumi:"progress"`
	Quiet              *bool    `pulumi:"quiet"`
	RandomWait         *bool    `pulumi:"randomWait"`
	ReportSpeed        *string  `pulumi:"reportSpeed"`
	SaveCookies        *string  `pulumi:"saveCookies"`
	ShowProgress       *bool    `pulumi:"showProgress"`
	StartPos           *string  `pulumi:"startPos"`
	Timeout            *string  `pulumi:"timeout"`
	Timestamping       *bool    `pulumi:"timestamping"`
	Tries              *int     `pulumi:"tries"`
	Urls               []string `pulumi:"urls"`
	User               *string  `pulumi:"user"`
	UserAgent          *string  `pulumi:"userAgent"`
	Verbose            *bool    `pulumi:"verbose"`
	Version            *string  `pulumi:"version"`
	Wait               *string  `pulumi:"wait"`
}

// The set of arguments for constructing a Wget resource.
type WgetArgs struct {
	AppendOutput       pulumi.StringPtrInput
	Background         pulumi.BoolPtrInput
	Base               pulumi.StringPtrInput
	CaCertificateFile  pulumi.StringPtrInput
	CaDirectory        pulumi.StringPtrInput
	Certificate        pulumi.StringPtrInput
	CertificateType    pulumi.StringPtrInput
	Config             pulumi.StringPtrInput
	Continue           pulumi.BoolPtrInput
	CrlFile            pulumi.StringPtrInput
	CutDirs            pulumi.IntPtrInput
	Debug              pulumi.BoolPtrInput
	DirectoryPrefix    pulumi.StringPtrInput
	Execute            pulumi.StringArrayInput
	ForceDirectories   pulumi.BoolPtrInput
	ForceHtml          pulumi.BoolPtrInput
	Help               pulumi.BoolPtrInput
	HttpsOnly          pulumi.BoolPtrInput
	Inet4Only          pulumi.BoolPtrInput
	InputFile          pulumi.StringPtrInput
	KeepSessionCookies pulumi.BoolPtrInput
	NoClobber          pulumi.BoolPtrInput
	NoDirectories      pulumi.BoolPtrInput
	NoDnsCache         pulumi.BoolPtrInput
	NoVerbose          pulumi.BoolPtrInput
	OutputDocument     pulumi.StringPtrInput
	OutputFile         pulumi.StringPtrInput
	Password           pulumi.StringPtrInput
	PrivateKey         pulumi.StringPtrInput
	PrivateKeyType     pulumi.StringPtrInput
	Progress           pulumi.StringPtrInput
	Quiet              pulumi.BoolPtrInput
	RandomWait         pulumi.BoolPtrInput
	ReportSpeed        pulumi.StringPtrInput
	SaveCookies        pulumi.StringPtrInput
	ShowProgress       pulumi.BoolPtrInput
	StartPos           pulumi.StringPtrInput
	Timeout            pulumi.StringPtrInput
	Timestamping       pulumi.BoolPtrInput
	Tries              pulumi.IntPtrInput
	Urls               pulumi.StringArrayInput
	User               pulumi.StringPtrInput
	UserAgent          pulumi.StringPtrInput
	Verbose            pulumi.BoolPtrInput
	Version            pulumi.StringPtrInput
	Wait               pulumi.StringPtrInput
}

func (WgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wgetArgs)(nil)).Elem()
}

type WgetInput interface {
	pulumi.Input

	ToWgetOutput() WgetOutput
	ToWgetOutputWithContext(ctx context.Context) WgetOutput
}

func (*Wget) ElementType() reflect.Type {
	return reflect.TypeOf((**Wget)(nil)).Elem()
}

func (i *Wget) ToWgetOutput() WgetOutput {
	return i.ToWgetOutputWithContext(context.Background())
}

func (i *Wget) ToWgetOutputWithContext(ctx context.Context) WgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WgetOutput)
}

func (i *Wget) ToOutput(ctx context.Context) pulumix.Output[*Wget] {
	return pulumix.Output[*Wget]{
		OutputState: i.ToWgetOutputWithContext(ctx).OutputState,
	}
}

type WgetOutput struct{ *pulumi.OutputState }

func (WgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wget)(nil)).Elem()
}

func (o WgetOutput) ToWgetOutput() WgetOutput {
	return o
}

func (o WgetOutput) ToWgetOutputWithContext(ctx context.Context) WgetOutput {
	return o
}

func (o WgetOutput) ToOutput(ctx context.Context) pulumix.Output[*Wget] {
	return pulumix.Output[*Wget]{
		OutputState: o.OutputState,
	}
}

func (o WgetOutput) Args() WgetArgsTypeOutput {
	return o.ApplyT(func(v *Wget) WgetArgsTypeOutput { return v.Args }).(WgetArgsTypeOutput)
}

func (o WgetOutput) CreatedFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringArrayOutput { return v.CreatedFiles }).(pulumi.StringArrayOutput)
}

func (o WgetOutput) ExitCode() pulumi.IntOutput {
	return o.ApplyT(func(v *Wget) pulumi.IntOutput { return v.ExitCode }).(pulumi.IntOutput)
}

func (o WgetOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

func (o WgetOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WgetInput)(nil)).Elem(), &Wget{})
	pulumi.RegisterOutputType(WgetOutput{})
}
