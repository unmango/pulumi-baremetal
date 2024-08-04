package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type CommandArgs interface {
	Cmd() *pb.Command
	ExpectedFiles() []string
}

type CommandState[T CommandArgs] struct {
	Args         T        `pulumi:"args"`
	ExitCode     int      `pulumi:"exitCode"`
	Stderr       string   `pulumi:"stderr"`
	Stdout       string   `pulumi:"stdout"`
	CreatedFiles []string `pulumi:"createdFiles"`
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
		ExpectCreated: inputs.ExpectedFiles(),
	})
	if err != nil {
		return fmt.Errorf("sending create request: %w", err)
	}

	if res.CreatedFiles == nil {
		log.Debug("files was empty, this is probably a bug somewhere else")
		res.CreatedFiles = []string{}
	}

	s.Args = inputs
	s.ExitCode = int(res.Result.ExitCode)
	s.Stdout = res.Result.Stdout
	s.Stderr = res.Result.Stderr
	s.CreatedFiles = res.CreatedFiles

	log.Info("create success")
	return nil
}

func (s *CommandState[T]) Update(ctx context.Context, inputs T, preview bool) (CommandState[T], error) {
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
		ExpectCreated: inputs.ExpectedFiles(),
		Create: &pb.Operation{
			CreatedFiles: s.CreatedFiles,
			Command:      s.Args.Cmd(),
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

	result.Args = inputs
	result.ExitCode = int(res.Result.ExitCode)
	result.Stdout = res.Result.Stdout
	result.Stderr = res.Result.Stderr
	result.CreatedFiles = res.CreatedFiles

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
			CreatedFiles: s.CreatedFiles,
			Command:      s.Args.Cmd(),
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
		Args:         s.Args,
		ExitCode:     s.ExitCode,
		Stderr:       s.Stderr,
		Stdout:       s.Stdout,
		CreatedFiles: s.CreatedFiles,
	}
}
