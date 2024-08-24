package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	cmd "github.com/unmango/pulumi-baremetal/provider/pkg/command"
	"github.com/unmango/pulumi-baremetal/provider/pkg/operation"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

func (s *State[T]) Delete(ctx context.Context) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("Failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	var command *pb.Command
	if len(s.CustomDelete) > 0 {
		command, err = cmd.Parse(s.CustomDelete)
		if err != nil {
			log.Errorf("Failed to parse custom delete: %s", err)
			return fmt.Errorf("parsing custom command: %w", err)
		}
	} else {
		log.DebugStatus("Normal delete")
	}

	prev, err := s.Operation()
	if err != nil {
		log.Errorf("Failed generating operation from state: %s", err)
		return fmt.Errorf("failed to generate operation from state: %w", err)
	}

	log.DebugStatus("Sending delete request to provisioner")
	res, err := p.Delete(ctx, &pb.DeleteRequest{
		Command:  command,
		Previous: prev,
	})
	if err != nil {
		log.Errorf("Failed provisioning: %s", err)
		return fmt.Errorf("sending delete request: %w", err)
	}

	if len(res.Commands) == 0 {
		log.DebugStatus("Provisioner did not perform any operations")
	} else {
		failed := []*pb.Operation{}
		for _, o := range res.Commands {
			exitCode := o.Result.ExitCode
			display := operation.Display(o)

			if exitCode != 0 {
				log.Errorf(display)
				failed = append(failed, o)
			} else {
				log.InfoStatusf("✅ %s", display)
			}
		}

		if len(failed) > 0 {
			log.WarningStatusf("%d delete operation(s) failed", len(failed))
		}
	}

	log.InfoStatus(operation.DisplayAll(res.Commands))
	return nil
}
