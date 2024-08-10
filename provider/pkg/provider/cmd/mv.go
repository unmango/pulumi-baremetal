package cmd

import (
	"context"
	"fmt"
	"path"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type MvArgs struct {
	CommandArgsBase

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
	b := builder{m.Source}
	b.opv(m.Backup, "--backup")
	b.op(m.Force, "--force")
	b.op(m.NoClobber, "--no-clobber")
	b.op(m.StripTrailingSlashes, "--strip-trailing-slashes")
	b.opv(m.Suffix, "--suffix")
	b.op(m.Update, "--update")
	b.op(m.Verbose, "--verbose")
	b.op(m.Version, "--version")

	b.opv(m.TargetDirectory, "--target-directory")
	b.op(m.NoTargetDirectory, "--no-target-directory")

	b.arg(m.Destination)
	b.arg(m.Directory)

	return &pb.Command{
		Bin:  pb.Bin_BIN_MV,
		Args: b.args,
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

var _ CommandBuilder = MvArgs{}

type Mv struct{}

type MvState = CommandState[MvArgs]

// Create implements infer.CustomCreate.
func (Mv) Create(ctx context.Context, name string, inputs CommandArgs[MvArgs], preview bool) (id string, output MvState, err error) {
	state := MvState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("mv: %w", err)
	}

	return name, state, nil
}

// Diff implements infer.CustomDiff.
func (Mv) Diff(ctx context.Context, id string, olds MvState, news CommandArgs[MvArgs]) (provider.DiffResponse, error) {
	diff, err := olds.Diff(ctx, news)
	if err != nil {
		return provider.DiffResponse{}, fmt.Errorf("mv: %w", err)
	}

	if news.Args.Backup != olds.Args.Backup {
		diff["args.backup"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Destination != olds.Args.Destination {
		diff["args.destination"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Directory != olds.Args.Directory {
		diff["args.directory"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Force != olds.Args.Force {
		diff["args.force"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.NoClobber != olds.Args.NoClobber {
		diff["args.noClobber"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.NoTargetDirectory != olds.Args.NoTargetDirectory {
		diff["args.noTargetDirectory"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if !slices.Equal(news.Args.Source, olds.Args.Source) {
		diff["args.source"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.StripTrailingSlashes != olds.Args.StripTrailingSlashes {
		diff["args.stripTrailingSlashes"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Suffix != olds.Args.Suffix {
		diff["args.suffix"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.TargetDirectory != olds.Args.TargetDirectory {
		diff["args.targetDirectory"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Update != olds.Args.Update {
		diff["args.update"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

// Update implements infer.CustomUpdate.
func (Mv) Update(ctx context.Context, id string, olds MvState, news CommandArgs[MvArgs], preview bool) (MvState, error) {
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

var _ = (infer.CustomCreate[CommandArgs[MvArgs], MvState])((*Mv)(nil))
var _ = (infer.CustomDiff[CommandArgs[MvArgs], MvState])((*Mv)(nil))
var _ = (infer.CustomUpdate[CommandArgs[MvArgs], MvState])((*Mv)(nil))
var _ = (infer.CustomDelete[MvState])((*Mv)(nil))
