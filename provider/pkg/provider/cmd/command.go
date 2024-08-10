package cmd

import (
	"context"
	"fmt"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type FsManipulator interface {
	ExpectCreated() []string
	ExpectMoved() map[string]string
}

type CommandArgsBase struct{}

func (CommandArgsBase) ExpectCreated() []string {
	return []string{}
}

func (CommandArgsBase) ExpectMoved() map[string]string {
	return map[string]string{}
}

type CommandBuilder interface {
	FsManipulator
	Cmd() *pb.Command
}

type CommandArgs[T CommandBuilder] struct {
	Args         T        `pulumi:"args"`
	Triggers     []any    `pulumi:"triggers,optional"`
	CustomUpdate []string `pulumi:"customUpdate,optional"`
	CustomDelete []string `pulumi:"customDelete,optional"`
}

func (a *CommandArgs[T]) Cmd() *pb.Command {
	return a.Args.Cmd()
}

func (a *CommandArgs[T]) ExpectCreated() []string {
	return a.Args.ExpectCreated()
}

func (a *CommandArgs[T]) ExpectMoved() map[string]string {
	return a.Args.ExpectMoved()
}

type CommandState[T CommandBuilder] struct {
	CommandArgs[T]

	ExitCode     int               `pulumi:"exitCode"`
	Stderr       string            `pulumi:"stderr"`
	Stdout       string            `pulumi:"stdout"`
	CreatedFiles []string          `pulumi:"createdFiles"`
	MovedFiles   map[string]string `pulumi:"movedFiles"`
}

func (s *CommandState[T]) Create(ctx context.Context, inputs CommandArgs[T], preview bool) error {
	log := logger.FromContext(ctx)
	if preview {
		// Could dial the host and warn if the connection fails
		log.DebugStatus("skipping during preview")
		return nil
	}

	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	log.InfoStatus("Sending create request to provisioner")
	res, err := p.Create(ctx, &pb.CreateRequest{
		Command:       inputs.Cmd(),
		ExpectCreated: inputs.ExpectCreated(),
		ExpectMoved:   inputs.ExpectMoved(),
	})
	if err != nil {
		return fmt.Errorf("sending create request: %w", err)
	}

	if res.CreatedFiles == nil {
		log.DebugStatusf("%#v was empty, this is probably a bug somewhere else", res.CreatedFiles)
		res.CreatedFiles = []string{}
	}

	if res.MovedFiles == nil {
		log.DebugStatusf("%#v was empty, this is probably a bug somewhere else", res.MovedFiles)
		res.MovedFiles = map[string]string{}
	}

	s.CommandArgs = inputs
	s.ExitCode = int(res.Result.ExitCode)
	s.Stdout = res.Result.Stdout
	s.Stderr = res.Result.Stderr
	s.CreatedFiles = res.CreatedFiles
	s.MovedFiles = res.MovedFiles

	log.InfoStatus("Create success")
	return nil
}

func (s *CommandState[T]) Diff(ctx context.Context, inputs CommandArgs[T]) (map[string]provider.PropertyDiff, error) {
	diff := map[string]provider.PropertyDiff{}
	if !slices.Equal(s.Triggers, inputs.Triggers) {
		diff["triggers"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if !slices.Equal(s.CustomUpdate, inputs.CustomUpdate) {
		diff["customUpdate"] = provider.PropertyDiff{Kind: provider.Update}
	}

	if len(inputs.CustomDelete) > len(s.CustomDelete) {
		diff["customDelete"] = provider.PropertyDiff{Kind: provider.Add}
	}

	return diff, nil
}

func (s *CommandState[T]) Update(ctx context.Context, inputs CommandArgs[T], preview bool) (CommandState[T], error) {
	log := logger.FromContext(ctx)
	if preview {
		// Could dial the host and warn if the connection fails
		log.DebugStatus("skipping during preview")
		return s.Copy(), nil
	}

	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return s.Copy(), fmt.Errorf("creating provisioner: %w", err)
	}

	var command *pb.Command
	expectCreated := []string{}
	expectMoved := map[string]string{}

	if len(inputs.CustomUpdate) > 0 {
		command, err = parseCommand(s.CustomUpdate)
		command, err = parseCommand(inputs.CustomUpdate)
		if err != nil {
			log.Errorf("Failed to parse custom update: %s", err)
			return s.Copy(), fmt.Errorf("parsing custom command: %w", err)
		}
	} else {
		command = inputs.Cmd()
		expectCreated = inputs.ExpectCreated()
		expectMoved = inputs.ExpectMoved()
	}

	log.DebugStatus("Sending update request to provisioner")
	res, err := p.Update(ctx, &pb.UpdateRequest{
		Command:       command,
		ExpectCreated: expectCreated,
		ExpectMoved:   expectMoved,
		Previous: &pb.Operation{
			Command:      s.Cmd(),
			CreatedFiles: s.CreatedFiles,
			MovedFiles:   s.MovedFiles,
			Result: &pb.Result{
				ExitCode: int32(s.ExitCode),
				Stdout:   s.Stdout,
				Stderr:   s.Stderr,
			},
		},
	})
	if err != nil {
		return s.Copy(), fmt.Errorf("sending update request: %w", err)
	}

	if res.CreatedFiles == nil {
		log.DebugStatusf("%#v was empty, this is probably a bug somewhere else", res.CreatedFiles)
		res.CreatedFiles = []string{}
	}

	if res.MovedFiles == nil {
		log.DebugStatusf("%#v was empty, this is probably a bug somewhere else", res.MovedFiles)
		res.MovedFiles = map[string]string{}
	}

	log.InfoStatus("Update success")
	return CommandState[T]{
		CommandArgs:  inputs,
		ExitCode:     int(res.Result.ExitCode),
		Stdout:       res.Result.Stdout,
		Stderr:       res.Result.Stderr,
		CreatedFiles: res.CreatedFiles,
		MovedFiles:   res.MovedFiles,
	}, nil
}

func (s *CommandState[T]) Delete(ctx context.Context) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	var command *pb.Command
	if len(s.CustomDelete) > 0 {
		command, err = parseCommand(s.CustomUpdate)
		command, err = parseCommand(s.CustomDelete)
		if err != nil {
			log.Errorf("Failed to parse custom delete: %s", err)
			return fmt.Errorf("parsing custom command: %w", err)
		}
	} else {
		log.InfoStatus("Normal delete")
	}

	log.InfoStatus("Sending delete request to provisioner")
	res, err := p.Delete(ctx, &pb.DeleteRequest{
		Command: command,
		Previous: &pb.Operation{
			Command:      s.Cmd(),
			CreatedFiles: s.CreatedFiles,
			MovedFiles:   s.MovedFiles,
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
		log.InfoStatus("provisioner did not perform any operations")
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

	log.InfoStatus("Delete success")
	return nil
}

func (a CommandArgs[T]) UpdateKind() provider.DiffKind {
	if len(a.CustomUpdate) > 0 {
		return provider.Update
	} else {
		return provider.UpdateReplace
	}
}

func (a CommandArgs[T]) DeleteBeforeReplace() bool {
	return len(a.CustomUpdate) == 0
}

func (s *CommandState[T]) Copy() CommandState[T] {
	return CommandState[T]{
		CommandArgs:  s.CommandArgs,
		ExitCode:     s.ExitCode,
		Stderr:       s.Stderr,
		Stdout:       s.Stdout,
		CreatedFiles: s.CreatedFiles,
		MovedFiles:   s.MovedFiles,
	}
}

func parseCommand(args []string) (*pb.Command, error) {
	bin := pb.Bin_BIN_UNSPECIFIED
	switch args[0] {
	case "chmod":
		bin = pb.Bin_BIN_CHMOD
	case "mkdir":
		bin = pb.Bin_BIN_MKDIR
	case "mktemp":
		bin = pb.Bin_BIN_MKTEMP
	case "mv":
		bin = pb.Bin_BIN_MV
	case "rm":
		bin = pb.Bin_BIN_RM
	case "tar":
		bin = pb.Bin_BIN_TAR
	case "tee":
		bin = pb.Bin_BIN_TEE
	case "wget":
		bin = pb.Bin_BIN_WGET
	default:
		return nil, fmt.Errorf("unsupported command: %s", args[0])
	}

	return &pb.Command{
		Bin:   bin,
		Args:  args[1:],
		Stdin: nil, // TODO
	}, nil
}
