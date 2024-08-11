package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

func (s *State[T]) Delete(ctx context.Context) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	var command *pb.Command
	if len(s.CustomDelete) > 0 {
		command, err = parseCommand(s.CustomDelete)
		if err != nil {
			log.Errorf("Failed to parse custom delete: %s", err)
			return fmt.Errorf("parsing custom command: %w", err)
		}
	} else {
		log.InfoStatus("Normal delete")
	}

	prev, err := s.Operation()
	if err != nil {
		log.Errorf("Failed generating operation from state: %s", err)
		return fmt.Errorf("failed to generate operation from state: %w", err)
	}

	log.InfoStatus("Sending delete request to provisioner")
	res, err := p.Delete(ctx, &pb.DeleteRequest{
		Command:  command,
		Previous: prev,
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
