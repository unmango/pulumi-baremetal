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

type TarArgs struct {
	CommandArgsBase

	Args []string `pulumi:"args,optional"`

	// Operation modes
	Append  bool `pulumi:"append,optional"`
	Create  bool `pulumi:"create,optional"`
	Delete  bool `pulumi:"delete,optional"`
	Diff    bool `pulumi:"diff,optional"`
	Extract bool `pulumi:"extract,optional"`
	List    bool `pulumi:"list,optional"`
	Update  bool `pulumi:"update,optional"`
	Version bool `pulummi:"version,optional"`

	// Operation modifiers
	NoSeek bool `pulumi:"noSeek,optional"`
	Sparse bool `pulumi:"sparse,optional"`

	// Overwrite control
	KeepOldFiles         bool `pulumi:"keepOldfiles,optional"`
	KeepNewerFiles       bool `pulumi:"keepNewerFiles,optional"`
	KeepDirectorySymlink bool `pulumi:"keepDirectorySymlink,optional"`
	NoOverwriteDir       bool `pulumi:"noOverwriteDir,optional"`
	Overwrite            bool `pulumi:"overwrite,optional"`
	OverwriteDir         bool `pulumi:"overwriteDir,optional"`
	RemoveFiles          bool `pulumi:"removeFiles,optional"`
	SkipOldFiles         bool `pulumi:"skipOldFiles,optional"`
	UnlinkFirst          bool `pulumi:"unlinkFirst,optional"`
	Verify               bool `pulumi:"verify,optional"`

	// Output stream selection
	IgnoreCommandError bool `pulumi:"ignoreCommandError,optional"`
	ToStdout           bool `pulumi:"toStdout,optional"`

	// Device selection and switching
	File string `pulumi:"file,optional"`

	// Compression options
	Bzip2 bool `pulumi:"bzip2,optional"`
	Gzip  bool `pulumi:"gzip,optional"`
	Xz    bool `pulumi:"xz,optional"`
	Lzip  bool `pulumi:"lzip,optional"`
	Lzma  bool `pulumi:"lzma,optional"`
	Lzop  bool `pulumi:"lzop,optional"`
	Zstd  bool `pulumi:"zstd,optional"`

	// Local file selection
	Directory         string `pulumi:"directory,optional"`
	Exclude           string `pulumi:"exclude,optional"`
	ExcludeVcs        bool   `pulumi:"excludeVcs,optional"`
	ExcludeVcsIgnores bool   `pulumi:"excludeVcsIgnores,optional"`
	Suffix            string `pulumi:"suffix,optional"`
	ExcludeFrom       string `pulumi:"excludeFrom,optional"`

	// File name transformations
	StripComponents int    `pulumi:"stripComponents,optional"`
	Transform       string `pulumi:"transform,optional"`

	// Informative output
	Verbose bool `pulumi:"verbose,optional"`
}

// Cmd implements CommandArgs.
func (t TarArgs) Cmd() *pb.Command {
	b := builder{t.Args}
	b.op(t.Append, "--append")
	b.op(t.Create, "--create")
	b.op(t.Delete, "--delete")
	b.op(t.Diff, "--diff")
	b.op(t.Extract, "--extract")
	b.op(t.List, "--list")

	b.opv(t.File, "--file")

	b.op(t.Bzip2, "--bzip2")
	b.op(t.Gzip, "--gzip")
	b.op(t.Lzip, "--lzip")
	b.op(t.Lzma, "--lzma")
	b.op(t.Lzop, "--lzop")
	b.op(t.Xz, "--xz")
	b.op(t.Zstd, "--zstd")

	b.op(t.ExcludeVcs, "--exclude-vcs")
	b.op(t.ExcludeVcsIgnores, "--exclude-vcs-ignores")
	b.op(t.IgnoreCommandError, "--ignore-command-error")
	b.op(t.KeepDirectorySymlink, "--keep-directory-symlink")
	b.op(t.KeepNewerFiles, "--keep-newer-files")
	b.op(t.KeepOldFiles, "--keep-old-files")
	b.op(t.NoOverwriteDir, "--no-overwrite-dir")
	b.op(t.NoSeek, "--no-seek")
	b.op(t.Overwrite, "--overwrite")
	b.op(t.OverwriteDir, "--overwrite-dir")
	b.op(t.RemoveFiles, "--remove-files")
	b.op(t.SkipOldFiles, "--skip-old-files")
	b.op(t.Sparse, "--sparse")
	b.op(t.ToStdout, "--to-stdout")
	b.op(t.UnlinkFirst, "--unlink-first")
	b.op(t.Update, "--update")
	b.op(t.Verbose, "--verbose")
	b.op(t.Verify, "--verify")
	b.op(t.Version, "--version")

	b.opv(t.Directory, "--directory")
	b.opv(t.Exclude, "--exclude")
	b.opv(t.ExcludeFrom, "--exclude-from")
	b.opv(t.Suffix, "--suffix")
	b.opv(t.Transform, "--transform")

	return &pb.Command{
		Bin:  pb.Bin_BIN_TAR,
		Args: b.args,
	}
}

// ExpectCreated implements CommandArgs.
// Subtle: this method shadows the method (DefaultFileManipulator).ExpectCreated of TarArgs.DefaultFileManipulator.
func (t TarArgs) ExpectCreated() []string {
	if t.Create && t.File != "" {
		return []string{t.File}
	}

	if t.Extract && len(t.Args) > 0 {
		files := make([]string, len(t.Args))
		for i, m := range t.Args {
			// `path.Join` will ignore empty elements, so if `t.Directory` wasn't provided
			// this should just return `m`. In that case, `m` should have the same value
			// as the relative path we want to remove when deleting, and we can return it here.
			files[i] = path.Join(t.Directory, m)
		}
		return files
	}

	return []string{}
}

var _ CommandBuilder = TarArgs{}

type Tar struct{}

type TarState = CommandState[TarArgs]

// Create implements infer.CustomCreate.
func (Tar) Create(ctx context.Context, name string, inputs CommandArgs[TarArgs], preview bool) (id string, output TarState, err error) {
	state := TarState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("tar: %w", err)
	}

	return name, state, nil
}

// Diff implements infer.CustomDiff.
func (Tar) Diff(ctx context.Context, id string, olds TarState, news CommandArgs[TarArgs]) (provider.DiffResponse, error) {
	diff, err := olds.Diff(ctx, news)
	if err != nil {
		return provider.DiffResponse{}, fmt.Errorf("rm: %w", err)
	}

	if news.Args.Append != olds.Args.Append {
		diff["args.append"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if !slices.Equal(news.Args.Args, olds.Args.Args) { // lol
		diff["args.args"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Bzip2 != olds.Args.Bzip2 {
		diff["args.bzip2"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Create != olds.Args.Create {
		diff["args.create"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Delete != olds.Args.Delete {
		diff["args.delete"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Diff != olds.Args.Diff {
		diff["args.diff"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.ExcludeVcs != olds.Args.ExcludeVcs {
		diff["args.excludeVcs"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.ExcludeVcsIgnores != olds.Args.ExcludeVcsIgnores {
		diff["args.excludeVcsIgnoes"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Extract != olds.Args.Extract {
		diff["args.extract"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Gzip != olds.Args.Gzip {
		diff["args.gzip"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.IgnoreCommandError != olds.Args.IgnoreCommandError {
		diff["args.ignoreCommandError"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.NoOverwriteDir != olds.Args.OverwriteDir {
		diff["args.noOverwriteDir"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.RemoveFiles != olds.Args.RemoveFiles {
		diff["args.removeFiles"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Directory != olds.Args.Directory {
		diff["args.directory"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Exclude != olds.Args.Exclude {
		diff["args.exclude"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.File != olds.Args.File {
		diff["args.file"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.StripComponents != olds.Args.StripComponents {
		diff["args.stripComponents"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Transform != olds.Args.Transform {
		diff["args.transform"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

// Update implements infer.CustomUpdate.
func (Tar) Update(ctx context.Context, id string, olds TarState, news CommandArgs[TarArgs], preview bool) (TarState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("tar: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Tar) Delete(ctx context.Context, id string, props TarState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("tar: %w", err)
	}

	return nil
}

var _ = (infer.CustomCreate[CommandArgs[TarArgs], TarState])((*Tar)(nil))
var _ = (infer.CustomDiff[CommandArgs[TarArgs], TarState])((*Tar)(nil))
var _ = (infer.CustomUpdate[CommandArgs[TarArgs], TarState])((*Tar)(nil))
var _ = (infer.CustomDelete[TarState])((*Tar)(nil))
