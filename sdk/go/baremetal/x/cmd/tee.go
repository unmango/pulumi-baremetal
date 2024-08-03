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

// TEE(1)                           User Commands                          TEE(1)
//
// NAME
//
//	tee - read from standard input and write to standard output and files
//
// SYNOPSIS
//
//	tee [OPTION]... [FILE]...
//
// DESCRIPTION
//
//	    Copy standard input to each FILE, and also to standard output.
//
//	    -a, --append
//	           append to the given FILEs, do not overwrite
//
//	    -i, --ignore-interrupts
//	           ignore interrupt signals
//
//	    -p     operate in a more appropriate MODE with pipes.
//
//	    --output-error[=MODE]
//	           set behavior on write error.  See MODE below
//
//	    --help display this help and exit
//
//	    --version
//	           output version information and exit
//
//	MODE determines behavior with write errors on the outputs:
//	    warn   diagnose errors writing to any output
//
//	    warn-nopipe
//	           diagnose errors writing to any output not a pipe
//
//	    exit   exit on error writing to any output
//
//	    exit-nopipe
//	           exit on error writing to any output not a pipe
//
//	    The  default  MODE  for  the -p option is 'warn-nopipe'.  With "nopipe"
//	    MODEs, exit immediately if all outputs become broken  pipes.   The  de‐
//	    fault  operation when --output-error is not specified, is to exit imme‐
//	    diately on error writing to a pipe, and diagnose errors writing to  non
//	    pipe outputs.
//
// AUTHOR
//
//	Written by Mike Parker, Richard M. Stallman, and David MacKenzie.
//
// REPORTING BUGS
//
//	GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
//	Report any translation bugs to <https://translationproject.org/team/>
//
// COPYRIGHT
//
//	Copyright  ©  2024  Free Software Foundation, Inc.  License GPLv3+: GNU
//	GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
//	This is free software: you are free  to  change  and  redistribute  it.
//	There is NO WARRANTY, to the extent permitted by law.
//
// SEE ALSO
//
//	Full documentation <https://www.gnu.org/software/coreutils/tee>
//	or available locally via: info '(coreutils) tee invocation'
//
// GNU coreutils 9.5                 March 2024                            TEE(1)
type Tee struct {
	pulumi.CustomResourceState

	Args         pulumix.GPtrOutput[TeeArgsType, TeeArgsTypeOutput] `pulumi:"args"`
	CreatedFiles pulumix.ArrayOutput[string]                        `pulumi:"createdFiles"`
	ExitCode     pulumix.Output[int]                                `pulumi:"exitCode"`
	Stderr       pulumix.Output[string]                             `pulumi:"stderr"`
	Stdout       pulumix.Output[string]                             `pulumi:"stdout"`
}

// NewTee registers a new resource with the given unique name, arguments, and options.
func NewTee(ctx *pulumi.Context,
	name string, args *TeeArgs, opts ...pulumi.ResourceOption) (*Tee, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.Files == nil {
		return nil, errors.New("invalid value for required argument 'Files'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Tee
	err := ctx.RegisterResource("baremetal:cmd:Tee", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTee gets an existing Tee resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTee(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeeState, opts ...pulumi.ResourceOption) (*Tee, error) {
	var resource Tee
	err := ctx.ReadResource("baremetal:cmd:Tee", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tee resources.
type teeState struct {
}

type TeeState struct {
}

func (TeeState) ElementType() reflect.Type {
	return reflect.TypeOf((*teeState)(nil)).Elem()
}

type teeArgs struct {
	Append  *bool    `pulumi:"append"`
	Content string   `pulumi:"content"`
	Files   []string `pulumi:"files"`
}

// The set of arguments for constructing a Tee resource.
type TeeArgs struct {
	Append  pulumix.Input[*bool]
	Content pulumix.Input[string]
	Files   pulumix.Input[[]string]
}

func (TeeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teeArgs)(nil)).Elem()
}

type TeeOutput struct{ *pulumi.OutputState }

func (TeeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tee)(nil)).Elem()
}

func (o TeeOutput) ToTeeOutput() TeeOutput {
	return o
}

func (o TeeOutput) ToTeeOutputWithContext(ctx context.Context) TeeOutput {
	return o
}

func (o TeeOutput) ToOutput(ctx context.Context) pulumix.Output[Tee] {
	return pulumix.Output[Tee]{
		OutputState: o.OutputState,
	}
}

func (o TeeOutput) Args() pulumix.GPtrOutput[TeeArgsType, TeeArgsTypeOutput] {
	value := pulumix.Apply[Tee](o, func(v Tee) pulumix.GPtrOutput[TeeArgsType, TeeArgsTypeOutput] { return v.Args })
	unwrapped := pulumix.Flatten[*TeeArgsType, pulumix.GPtrOutput[TeeArgsType, TeeArgsTypeOutput]](value)
	return pulumix.GPtrOutput[TeeArgsType, TeeArgsTypeOutput]{OutputState: unwrapped.OutputState}
}

func (o TeeOutput) CreatedFiles() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Tee](o, func(v Tee) pulumix.ArrayOutput[string] { return v.CreatedFiles })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

func (o TeeOutput) ExitCode() pulumix.Output[int] {
	value := pulumix.Apply[Tee](o, func(v Tee) pulumix.Output[int] { return v.ExitCode })
	return pulumix.Flatten[int, pulumix.Output[int]](value)
}

func (o TeeOutput) Stderr() pulumix.Output[string] {
	value := pulumix.Apply[Tee](o, func(v Tee) pulumix.Output[string] { return v.Stderr })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func (o TeeOutput) Stdout() pulumix.Output[string] {
	value := pulumix.Apply[Tee](o, func(v Tee) pulumix.Output[string] { return v.Stdout })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func init() {
	pulumi.RegisterOutputType(TeeOutput{})
}
