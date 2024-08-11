package coreutils

import (
	"context"
	"fmt"
	"path"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"
)

type MvArgs struct {
	cmd.ArgsBase

	Backup               string   `pulumi:"backup,optional"`
	Destination          string   `pulumi:"destination,optional"`
	Directory            string   `pulumi:"directory,optional"`
	Force                bool     `pulumi:"force,optional"`
	Help                 bool     `pulumi:"help,optional"`
	NoClobber            bool     `pulumi:"noClobber,optional"`
	NoTargetDirectory    bool     `pulumi:"noTargetDirectory,optional"`
	Source               []string `pulumi:"source"`
	StripTrailingSlashes bool     `pulumi:"stripTrailingSlashes,optional"`
	Suffix               string   `pulumi:"suffix,optional"`
	TargetDirectory      string   `pulumi:"targetDirectory,optional"`
	Update               bool     `pulumi:"update,optional"`
	Verbose              bool     `pulumi:"verbose,optional"`
	Version              bool     `pulumi:"version,optional"`
}

// Cmd implements CommandArgs.
func (m MvArgs) Cmd() *pb.Command {
	b := cmd.B{Args: m.Source}

	b.Opv(m.Backup, "--backup")
	b.Op(m.Force, "--force")
	b.Op(m.NoClobber, "--no-clobber")
	b.Op(m.StripTrailingSlashes, "--strip-trailing-slashes")
	b.Opv(m.Suffix, "--suffix")
	b.Op(m.Update, "--update")
	b.Op(m.Verbose, "--verbose")
	b.Op(m.Version, "--version")

	b.Opv(m.TargetDirectory, "--target-directory")
	b.Op(m.NoTargetDirectory, "--no-target-directory")

	b.Arg(m.Destination)
	b.Arg(m.Directory)

	return &pb.Command{
		Bin:  pb.Bin_BIN_MV,
		Args: b.Args,
	}
}

// ExpectMoved implements FileManipulator.
func (m MvArgs) ExpectMoved() map[string]string {
	files := map[string]string{}
	mvSrc := func(d string) {
		for _, f := range m.Source {
			files[f] = path.Join(d, path.Base(f))
		}
	}

	if m.Destination != "" && len(m.Source) == 1 {
		files[m.Source[0]] = m.Destination
	} else if m.Directory != "" {
		mvSrc(m.Directory)
	} else if m.TargetDirectory != "" {
		mvSrc(m.TargetDirectory)
	}

	return files
}

var _ cmd.Builder = MvArgs{}

type Mv struct{}

type MvState = cmd.State[MvArgs]

// Create implements infer.CustomCreate.
func (Mv) Create(ctx context.Context, name string, inputs cmd.CommandArgs[MvArgs], preview bool) (id string, output MvState, err error) {
	state := MvState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("mv: %w", err)
	}

	return name, state, nil
}

// Diff implements infer.CustomDiff.
func (Mv) Diff(ctx context.Context, id string, olds MvState, news cmd.CommandArgs[MvArgs]) (provider.DiffResponse, error) {
	diff, err := olds.Diff(ctx, news)
	if err != nil {
		return provider.DiffResponse{}, fmt.Errorf("mv: %w", err)
	}

	defaultKind := news.UpdateKind()

	if news.Args.Backup != olds.Args.Backup {
		diff["args.backup"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.Destination != olds.Args.Destination {
		diff["args.destination"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.Directory != olds.Args.Directory {
		diff["args.directory"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.Force != olds.Args.Force {
		diff["args.force"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.NoClobber != olds.Args.NoClobber {
		diff["args.noClobber"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.NoTargetDirectory != olds.Args.NoTargetDirectory {
		diff["args.noTargetDirectory"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if !slices.Equal(news.Args.Source, olds.Args.Source) {
		diff["args.source"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.StripTrailingSlashes != olds.Args.StripTrailingSlashes {
		diff["args.stripTrailingSlashes"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.Suffix != olds.Args.Suffix {
		diff["args.suffix"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.TargetDirectory != olds.Args.TargetDirectory {
		diff["args.targetDirectory"] = provider.PropertyDiff{Kind: defaultKind}
	}

	if news.Args.Update != olds.Args.Update {
		diff["args.update"] = provider.PropertyDiff{Kind: defaultKind}
	}

	changes := false
	for k, v := range diff {
		if k == "customDelete" {
			continue
		}

		switch v.Kind {
		case provider.Update:
			fallthrough
		case provider.UpdateReplace:
			changes = true
		}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          changes,
		DetailedDiff:        diff,
	}, nil
}

// Update implements infer.CustomUpdate.
func (Mv) Update(ctx context.Context, id string, olds MvState, news cmd.CommandArgs[MvArgs], preview bool) (MvState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("mv: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Mv) Delete(ctx context.Context, id string, props MvState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("mv: %w", err)
	}

	return nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[MvArgs], MvState])((*Mv)(nil))
var _ = (infer.CustomDiff[cmd.CommandArgs[MvArgs], MvState])((*Mv)(nil))
var _ = (infer.CustomUpdate[cmd.CommandArgs[MvArgs], MvState])((*Mv)(nil))
var _ = (infer.CustomDelete[MvState])((*Mv)(nil))
