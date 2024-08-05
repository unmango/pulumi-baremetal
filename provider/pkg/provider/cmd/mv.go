package cmd

import (
	"context"
	"fmt"
	"path"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type MvArgs struct {
	DefaultFileManipulator
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

var _ CommandArgs = MvArgs{}

type Mv struct{}

type MvState = CommandState[MvArgs]

// Create implements infer.CustomCreate.
func (Mv) Create(ctx context.Context, name string, inputs MvArgs, preview bool) (id string, output MvState, err error) {
	state := MvState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("mv: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Mv) Update(ctx context.Context, id string, olds MvState, news MvArgs, preview bool) (MvState, error) {
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

var _ = (infer.CustomCreate[MvArgs, MvState])((*Mv)(nil))
var _ = (infer.CustomUpdate[MvArgs, MvState])((*Mv)(nil))
