package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type CommandState struct {
	Stderr       string      `pulumi:"stderr"`
	Stdout       string      `pulumi:"stdout"`
	Create       *pb.Command `pulumi:"create"`
	CreatedFiles []string    `pulumi:"createdFiles"`
}

func (s *CommandState) create(ctx context.Context, bin pb.Bin, args []string, stdin *string) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	log.Debug("sending create request")
	cmd := &pb.Command{Bin: bin, Args: args, Stdin: stdin}
	res, err := p.Create(ctx, &pb.CreateRequest{Command: cmd})
	if err != nil {
		return fmt.Errorf("sending create request: %w", err)
	}

	s.Stdout = res.Result.Stdout
	s.Stderr = res.Result.Stderr
	s.Create = cmd

	log.Info("create success")
	return nil
}

func (s *CommandState) delete(ctx context.Context) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	log.Debug("sending create request")
	res, err := p.Delete(ctx, &pb.DeleteRequest{Create: s.Create})
	if err != nil {
		return fmt.Errorf("sending create request: %w", err)
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
