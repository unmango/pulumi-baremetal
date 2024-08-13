package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

func (s *State[T]) Create(ctx context.Context, inputs CommandArgs[T], preview bool) error {
	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("Failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	if preview {
		if _, err = p.Ping(ctx, &pb.PingRequest{}); err != nil {
			log.WarningStatusf("Failed pinging provisioner: %s", err)
		}

		return nil
	}

	command, err := inputs.Cmd()
	if err != nil {
		log.Errorf("Failed constructing command: %s", err)
		return err
	}

	log.DebugStatus("Sending create request to provisioner")
	res, err := p.Create(ctx, &pb.CreateRequest{
		Command:       command,
		ExpectCreated: inputs.ExpectCreated(),
		ExpectMoved:   inputs.ExpectMoved(),
	})
	if err != nil {
		log.Errorf("command:%s %s", command, err)
		return fmt.Errorf("sending create request: %w", err)
	}

	if res.Result.ExitCode > 0 {
		log.Errorf("%s %s", command, res.Result)
		return fmt.Errorf("sending create request: %s", res.Result)
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

	log.InfoStatusf("✅ command:%v, %v", command, res.Result)
	return nil
}
