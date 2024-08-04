package cmd

import (
	"context"
	"fmt"
	"path"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type TarArgs struct {
	DefaultFileManipulator

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

var _ CommandArgs = TarArgs{}

type Tar struct{}

type TarState = CommandState[TarArgs]

// Create implements infer.CustomCreate.
func (Tar) Create(ctx context.Context, name string, inputs TarArgs, preview bool) (id string, output TarState, err error) {
	state := TarState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("tar: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Tar) Update(ctx context.Context, id string, olds TarState, news TarArgs, preview bool) (TarState, error) {
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

var _ = (infer.CustomCreate[TarArgs, TarState])((*Tar)(nil))
var _ = (infer.CustomUpdate[TarArgs, TarState])((*Tar)(nil))
var _ = (infer.CustomDelete[TarState])((*Tar)(nil))
