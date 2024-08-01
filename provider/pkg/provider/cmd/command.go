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
}

type CommandState[T CommandOpts] struct {
	Stderr       string   `pulumi:"stderr"`
	Stdout       string   `pulumi:"stdout"`
	CreateOpts   *T       `pulumi:"createOpts"`
	CreatedFiles []string `pulumi:"createdFiles"`
}

func (s *CommandState[T]) Create(ctx context.Context, opts T) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	log.Debug("sending create request")
	res, err := p.Create(ctx, &pb.CreateRequest{Command: opts.Cmd()})
	if err != nil {
		return fmt.Errorf("sending create request: %w", err)
	}

	s.Stdout = res.Result.Stdout
	s.Stderr = res.Result.Stderr
	s.CreateOpts = &opts

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

	if s.CreateOpts == nil {
		log.Error("invalid command state")
		return errors.New("invalid state, create was nil")
	}

	log.Debug("sending delete request")
	res, err := p.Delete(ctx, &pb.DeleteRequest{
		Create: (*s.CreateOpts).Cmd(),
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
