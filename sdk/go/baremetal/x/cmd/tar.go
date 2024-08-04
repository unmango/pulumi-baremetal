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

type Tar struct {
	pulumi.CustomResourceState

	Args         pulumix.GPtrOutput[TarArgsType, TarArgsTypeOutput] `pulumi:"args"`
	CreatedFiles pulumix.ArrayOutput[string]                        `pulumi:"createdFiles"`
	ExitCode     pulumix.Output[int]                                `pulumi:"exitCode"`
	MovedFiles   pulumix.MapOutput[string]                          `pulumi:"movedFiles"`
	Stderr       pulumix.Output[string]                             `pulumi:"stderr"`
	Stdout       pulumix.Output[string]                             `pulumi:"stdout"`
}

// NewTar registers a new resource with the given unique name, arguments, and options.
func NewTar(ctx *pulumi.Context,
	name string, args *TarArgs, opts ...pulumi.ResourceOption) (*Tar, error) {
	if args == nil {
		args = &TarArgs{}
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
	Append               *bool    `pulumi:"append"`
	Args                 []string `pulumi:"args"`
	Bzip2                *bool    `pulumi:"bzip2"`
	Create               *bool    `pulumi:"create"`
	Delete               *bool    `pulumi:"delete"`
	Diff                 *bool    `pulumi:"diff"`
	Directory            *string  `pulumi:"directory"`
	Exclude              *string  `pulumi:"exclude"`
	ExcludeFrom          *string  `pulumi:"excludeFrom"`
	ExcludeVcs           *bool    `pulumi:"excludeVcs"`
	ExcludeVcsIgnores    *bool    `pulumi:"excludeVcsIgnores"`
	Extract              *bool    `pulumi:"extract"`
	File                 *string  `pulumi:"file"`
	Gzip                 *bool    `pulumi:"gzip"`
	IgnoreCommandError   *bool    `pulumi:"ignoreCommandError"`
	KeepDirectorySymlink *bool    `pulumi:"keepDirectorySymlink"`
	KeepNewerFiles       *bool    `pulumi:"keepNewerFiles"`
	KeepOldfiles         *bool    `pulumi:"keepOldfiles"`
	List                 *bool    `pulumi:"list"`
	Lzip                 *bool    `pulumi:"lzip"`
	Lzma                 *bool    `pulumi:"lzma"`
	Lzop                 *bool    `pulumi:"lzop"`
	NoOverwriteDir       *bool    `pulumi:"noOverwriteDir"`
	NoSeek               *bool    `pulumi:"noSeek"`
	Overwrite            *bool    `pulumi:"overwrite"`
	OverwriteDir         *bool    `pulumi:"overwriteDir"`
	RemoveFiles          *bool    `pulumi:"removeFiles"`
	SkipOldFiles         *bool    `pulumi:"skipOldFiles"`
	Sparse               *bool    `pulumi:"sparse"`
	StripComponents      *int     `pulumi:"stripComponents"`
	Suffix               *string  `pulumi:"suffix"`
	ToStdout             *bool    `pulumi:"toStdout"`
	Transform            *string  `pulumi:"transform"`
	UnlinkFirst          *bool    `pulumi:"unlinkFirst"`
	Update               *bool    `pulumi:"update"`
	Verbose              *bool    `pulumi:"verbose"`
	Verify               *bool    `pulumi:"verify"`
	Xz                   *bool    `pulumi:"xz"`
	Zstd                 *bool    `pulumi:"zstd"`
}

// The set of arguments for constructing a Tar resource.
type TarArgs struct {
	Append               pulumix.Input[*bool]
	Args                 pulumix.Input[[]string]
	Bzip2                pulumix.Input[*bool]
	Create               pulumix.Input[*bool]
	Delete               pulumix.Input[*bool]
	Diff                 pulumix.Input[*bool]
	Directory            pulumix.Input[*string]
	Exclude              pulumix.Input[*string]
	ExcludeFrom          pulumix.Input[*string]
	ExcludeVcs           pulumix.Input[*bool]
	ExcludeVcsIgnores    pulumix.Input[*bool]
	Extract              pulumix.Input[*bool]
	File                 pulumix.Input[*string]
	Gzip                 pulumix.Input[*bool]
	IgnoreCommandError   pulumix.Input[*bool]
	KeepDirectorySymlink pulumix.Input[*bool]
	KeepNewerFiles       pulumix.Input[*bool]
	KeepOldfiles         pulumix.Input[*bool]
	List                 pulumix.Input[*bool]
	Lzip                 pulumix.Input[*bool]
	Lzma                 pulumix.Input[*bool]
	Lzop                 pulumix.Input[*bool]
	NoOverwriteDir       pulumix.Input[*bool]
	NoSeek               pulumix.Input[*bool]
	Overwrite            pulumix.Input[*bool]
	OverwriteDir         pulumix.Input[*bool]
	RemoveFiles          pulumix.Input[*bool]
	SkipOldFiles         pulumix.Input[*bool]
	Sparse               pulumix.Input[*bool]
	StripComponents      pulumix.Input[*int]
	Suffix               pulumix.Input[*string]
	ToStdout             pulumix.Input[*bool]
	Transform            pulumix.Input[*string]
	UnlinkFirst          pulumix.Input[*bool]
	Update               pulumix.Input[*bool]
	Verbose              pulumix.Input[*bool]
	Verify               pulumix.Input[*bool]
	Xz                   pulumix.Input[*bool]
	Zstd                 pulumix.Input[*bool]
}

func (TarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tarArgs)(nil)).Elem()
}

type TarOutput struct{ *pulumi.OutputState }

func (TarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tar)(nil)).Elem()
}

func (o TarOutput) ToTarOutput() TarOutput {
	return o
}

func (o TarOutput) ToTarOutputWithContext(ctx context.Context) TarOutput {
	return o
}

func (o TarOutput) ToOutput(ctx context.Context) pulumix.Output[Tar] {
	return pulumix.Output[Tar]{
		OutputState: o.OutputState,
	}
}

func (o TarOutput) Args() pulumix.GPtrOutput[TarArgsType, TarArgsTypeOutput] {
	value := pulumix.Apply[Tar](o, func(v Tar) pulumix.GPtrOutput[TarArgsType, TarArgsTypeOutput] { return v.Args })
	unwrapped := pulumix.Flatten[*TarArgsType, pulumix.GPtrOutput[TarArgsType, TarArgsTypeOutput]](value)
	return pulumix.GPtrOutput[TarArgsType, TarArgsTypeOutput]{OutputState: unwrapped.OutputState}
}

func (o TarOutput) CreatedFiles() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Tar](o, func(v Tar) pulumix.ArrayOutput[string] { return v.CreatedFiles })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o TarOutput) ExitCode() pulumix.Output[int] {
	value := pulumix.Apply[Tar](o, func(v Tar) pulumix.Output[int] { return v.ExitCode })
	return pulumix.Flatten[int, pulumix.Output[int]](value)
}

func (o TarOutput) MovedFiles() pulumix.MapOutput[string] {
	value := pulumix.Apply[Tar](o, func(v Tar) pulumix.MapOutput[string] { return v.MovedFiles })
	unwrapped := pulumix.Flatten[map[string]string, pulumix.MapOutput[string]](value)
	return pulumix.MapOutput[string]{OutputState: unwrapped.OutputState}
}

func (o TarOutput) Stderr() pulumix.Output[string] {
	value := pulumix.Apply[Tar](o, func(v Tar) pulumix.Output[string] { return v.Stderr })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o TarOutput) Stdout() pulumix.Output[string] {
	value := pulumix.Apply[Tar](o, func(v Tar) pulumix.Output[string] { return v.Stdout })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func init() {
	pulumi.RegisterOutputType(TarOutput{})
}