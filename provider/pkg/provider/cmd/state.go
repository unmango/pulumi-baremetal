package cmd

import (
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type State[T Builder] struct {
	CommandArgs[T]

	ExitCode     int               `pulumi:"exitCode"`
	Stderr       string            `pulumi:"stderr"`
	Stdout       string            `pulumi:"stdout"`
	CreatedFiles []string          `pulumi:"createdFiles"`
	MovedFiles   map[string]string `pulumi:"movedFiles"`
}

func (s *State[T]) Copy() State[T] {
	return State[T]{
		CommandArgs:  s.CommandArgs,
		ExitCode:     s.ExitCode,
		Stderr:       s.Stderr,
		Stdout:       s.Stdout,
		CreatedFiles: s.CreatedFiles,
		MovedFiles:   s.MovedFiles,
	}
}

func (s *State[T]) Operation() (*pb.Operation, error) {
	command, err := s.Cmd()
	if err != nil {
		return nil, fmt.Errorf("failed to build command from state; %w", err)
	}

	return &pb.Operation{
		Command:      command,
		CreatedFiles: s.CreatedFiles,
		MovedFiles:   s.MovedFiles,
		Result: &pb.Result{
			ExitCode: int32(s.ExitCode),
			Stdout:   s.Stdout,
			Stderr:   s.Stderr,
		},
	}, nil
}
