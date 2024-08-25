package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	cmd "github.com/unmango/pulumi-baremetal/provider/pkg/command"
	"github.com/unmango/pulumi-baremetal/provider/pkg/operation"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
)

func (s *State[T]) Update(ctx context.Context, inputs CommandArgs[T], preview bool) (State[T], error) {
	log := logger.FromContext(ctx)
	p, err := inputs.Provisioner(ctx)
	if err != nil {
		log.Error("Failed creating provisioner")
		return s.Copy(), fmt.Errorf("creating provisioner: %w", err)
	}

	if preview {
		if _, err = p.Ping(ctx, &pb.PingRequest{}); err != nil {
			log.WarningStatusf("Failed pinging provisioner: %s", err)
		}

		return s.Copy(), nil
	}

	var command *pb.Command
	expectCreated := []string{}
	expectMoved := map[string]string{}

	if len(inputs.CustomUpdate) > 0 {
		command, err = cmd.Parse(inputs.CustomUpdate)
		if err != nil {
			log.Errorf("Failed parsing custom update: %s", err)
			return s.Copy(), fmt.Errorf("parsing custom command: %w", err)
		}
	} else {
		command, err = inputs.Cmd()
		if err != nil {
			log.Errorf("Failed building command from inputs: %s", err)
			return s.Copy(), fmt.Errorf("failed building command from inputs: %w", err)
		}

		expectCreated = inputs.ExpectCreated()
		expectMoved = inputs.ExpectMoved()
	}

	prev, err := s.Operation()
	if err != nil {
		log.Errorf("Failed generating operation from state: %s", err)
		return s.Copy(), fmt.Errorf("failed to generate operation from state: %w", err)
	}

	log.DebugStatus("Sending update request to provisioner")
	res, err := p.Update(ctx, &pb.UpdateRequest{
		Command:       command,
		ExpectCreated: expectCreated,
		ExpectMoved:   expectMoved,
		Previous:      prev,
	})
	if err != nil {
		log.Errorf("Failed provisioning: %s", err)
		return s.Copy(), fmt.Errorf("sending update request: %w", err)
	}

	op := operation.FromUpdate(command, res)
	display := operation.Display(op)

	if res.Result.ExitCode > 0 {
		log.Error(display)
		return s.Copy(), fmt.Errorf("update failed: %s", res.Result)
	}

	// Thought it might be nice to have a brief check mark on success.
	// In practice, this will probably happen too fast to be noticed. We'll see.
	log.InfoStatusf("✅ %s", display)

	if res.CreatedFiles == nil {
		log.DebugStatusf("%#v was empty, this is probably a bug somewhere else", res.CreatedFiles)
		res.CreatedFiles = []string{}
	}

	if res.MovedFiles == nil {
		log.DebugStatusf("%#v was empty, this is probably a bug somewhere else", res.MovedFiles)
		res.MovedFiles = map[string]string{}
	}

	log.InfoStatus(display)
	return State[T]{
		CommandArgs:  inputs,
		ExitCode:     int(res.Result.ExitCode),
		Stdout:       res.Result.Stdout,
		Stderr:       res.Result.Stderr,
		CreatedFiles: res.CreatedFiles,
		MovedFiles:   res.MovedFiles,
	}, nil
}
