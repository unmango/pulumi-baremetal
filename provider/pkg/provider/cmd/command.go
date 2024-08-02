package cmd

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type CommandOpts interface {
	Cmd() *pb.Command
	ExpectedFiles() []string
}

type CommandState[T CommandOpts] struct {
	ExitCode     int      `pulumi:"exitCode"`
	Stderr       string   `pulumi:"stderr"`
	Stdout       string   `pulumi:"stdout"`
	Args         *T       `pulumi:"args,optional"`
	CreatedFiles []string `pulumi:"createdFiles"`
}

func (s *CommandState[T]) Create(ctx context.Context, opts *T, preview bool) error {
	log := logger.FromContext(ctx)
	if opts == nil {
		log.Info("nothing to do")
		return nil
	}

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
		Command:     (*opts).Cmd(),
		ExpectFiles: (*opts).ExpectedFiles(),
	})
	if err != nil {
		return fmt.Errorf("sending create request: %w", err)
	}

	s.ExitCode = int(res.Result.ExitCode)
	s.Stdout = res.Result.Stdout
	s.Stderr = res.Result.Stderr
	s.Args = opts
	s.CreatedFiles = res.Files

	log.Info("create success")
	return nil
}

func (s *CommandState[T]) Delete(ctx context.Context) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	if s.Args == nil {
		log.Error("invalid command state")
		return errors.New("invalid state, create was nil")
	}

	log.Debug("sending delete request")
	res, err := p.Delete(ctx, &pb.DeleteRequest{
		Create: &pb.Operation{
			Files:   s.CreatedFiles,
			Command: (*s.Args).Cmd(),
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

	if res.Op == nil {
		log.Info("provisioner did not perform any operations")
		return nil
	}

	if res.Op.Result.ExitCode > 0 {
		log.Error("provisioner returned a non-success status code")
		return fmt.Errorf("delete operation failed exit code: %d", res.Op.Result.ExitCode)
	}

	log.Info("delete success")
	return nil
}
