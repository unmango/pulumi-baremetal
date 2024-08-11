package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

func (s *State[T]) Update(ctx context.Context, inputs CommandArgs[T], preview bool) (State[T], error) {
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
	return State[T]{
		CommandArgs:  inputs,
		ExitCode:     int(res.Result.ExitCode),
		Stdout:       res.Result.Stdout,
		Stderr:       res.Result.Stderr,
		CreatedFiles: res.CreatedFiles,
		MovedFiles:   res.MovedFiles,
	}, nil
}
