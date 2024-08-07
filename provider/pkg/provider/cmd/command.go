package cmd

import (
	"context"
	"fmt"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type FsManipulator interface {
	ExpectCreated() []string
	ExpectMoved() map[string]string
}

type CommandBuilder interface {
	FsManipulator
	Cmd() *pb.Command
}

type FsEntryType string

var _ = (infer.Enum[FsEntryType])((*FsEntryType)(nil))

const (
	FileFsEntry      FsEntryType = "file"
	DirectoryFsEntry FsEntryType = "directory"
)

func (*FsEntryType) Values() []infer.EnumValue[FsEntryType] {
	return []infer.EnumValue[FsEntryType]{
		{Value: FileFsEntry, Description: "A file"},
		{Value: DirectoryFsEntry, Description: "A directory"},
	}
}

type FsEntry struct {
	Path string      `pulumi:"path"`
	Type FsEntryType `pulumi:"type"`
}

type CommandArgs[T CommandBuilder] struct {
	Args     T         `pulumi:"args"`
	Triggers []any     `pulumi:"triggers,optional"`
	Expect   []FsEntry `pulumi:"expect,optional"`
}

type CommandState[T CommandBuilder] struct {
	CommandArgs[T]

	ExitCode int                `pulumi:"exitCode"`
	Stderr   string             `pulumi:"stderr"`
	Stdout   string             `pulumi:"stdout"`
	Creates  []FsEntry          `pulumi:"creates"`
	Moves    map[string]FsEntry `pulumi:"moves"`
}

func (s *CommandState[T]) Create(ctx context.Context, inputs T, preview bool) error {
	log := logger.FromContext(ctx)
	if preview {
		// Could dial the host and warn if the connection fails
		log.Debug("skipping during preview")
		return nil
	}

	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	log.Debug("sending create request")
	res, err := p.Create(ctx, &pb.CreateRequest{
		Command:       inputs.Cmd(),
		ExpectCreated: inputs.ExpectCreated(),
		ExpectMoved:   inputs.ExpectMoved(),
	})
	if err != nil {
		return fmt.Errorf("sending create request: %w", err)
	}

	if res.CreatedFiles == nil {
		log.Debugf("%#v was empty, this is probably a bug somewhere else", res.CreatedFiles)
		res.CreatedFiles = []string{}
	}

	if res.MovedFiles == nil {
		log.Debugf("%#v was empty, this is probably a bug somewhere else", res.MovedFiles)
		res.MovedFiles = map[string]string{}
	}

	s.Args = inputs
	s.ExitCode = int(res.Result.ExitCode)
	s.Stdout = res.Result.Stdout
	s.Stderr = res.Result.Stderr
	s.Creates = res.CreatedFiles
	s.Moves = res.MovedFiles

	log.Info("create success")
	return nil
}

func (s *CommandState[T]) Diff(ctx context.Context, inputs CommandArgs[T]) (provider.DiffResponse, error) {
	diff := map[string]provider.PropertyDiff{}

	if !slices.Equal(s.Triggers, inputs.Triggers) {
		diff["triggers"] = provider.PropertyDiff{Kind: provider.Update}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (s *CommandState[T]) Update(ctx context.Context, inputs CommandArgs[T], preview bool) (CommandState[T], error) {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	result := s.Copy()

	if err != nil {
		log.Error("failed creating provisioner")
		return result, fmt.Errorf("creating provisioner: %w", err)
	}

	log.Debug("sending update request")
	res, err := p.Update(ctx, &pb.UpdateRequest{
		Command:       inputs.Cmd(),
		ExpectCreated: inputs.ExpectCreated(),
		ExpectMoved:   inputs.ExpectMoved(),
		Create: &pb.Operation{
			Command:      s.Args.Cmd(),
			CreatedFiles: s.Creates,
			MovedFiles:   s.Moves,
			Result: &pb.Result{
				ExitCode: int32(s.ExitCode),
				Stdout:   s.Stdout,
				Stderr:   s.Stderr,
			},
		},
	})
	if err != nil {
		return result, fmt.Errorf("sending update request: %w", err)
	}

	if res.CreatedFiles == nil {
		log.Debugf("%#v was empty, this is probably a bug somewhere else", res.CreatedFiles)
		res.CreatedFiles = []string{}
	}

	if res.MovedFiles == nil {
		log.Debugf("%#v was empty, this is probably a bug somewhere else", res.MovedFiles)
		res.MovedFiles = map[string]string{}
	}

	result.Args = inputs
	result.ExitCode = int(res.Result.ExitCode)
	result.Stdout = res.Result.Stdout
	result.Stderr = res.Result.Stderr
	result.Creates = res.CreatedFiles
	result.Moves = res.MovedFiles

	log.Info("update success")
	return result, nil
}

func (s *CommandState[T]) Delete(ctx context.Context) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	log.Debug("sending delete request")
	res, err := p.Delete(ctx, &pb.DeleteRequest{
		Create: &pb.Operation{
			Command:      s.Args.Cmd(),
			CreatedFiles: s.Creates,
			MovedFiles:   s.Moves,
			Result: &pb.Result{
				ExitCode: int32(s.ExitCode),
				Stdout:   s.Stdout,
				Stderr:   s.Stderr,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("sending delete request: %w", err)
	}

	if len(res.Commands) == 0 {
		log.Info("provisioner did not perform any operations")
	} else {
		failed := []*pb.Operation{}
		for _, c := range res.Commands {
			exitCode := c.Result.ExitCode
			if exitCode != 0 {
				log.Errorf("provisioner returned non-zero exit code: %d", exitCode)
				failed = append(failed, c)
			}
		}

		if len(failed) > 0 {
			return fmt.Errorf("a delete operation failed: %s", failed)
		}
	}

	log.Info("delete success")
	return nil
}

func (s *CommandState[T]) Copy() CommandState[T] {
	return CommandState[T]{
		CommandArgs: s.CommandArgs,
		ExitCode:    s.ExitCode,
		Stderr:      s.Stderr,
		Stdout:      s.Stdout,
		Creates:     s.Creates,
		Moves:       s.Moves,
	}
}
