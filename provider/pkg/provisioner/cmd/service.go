package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal"
)

type service struct {
	pb.UnimplementedCommandServiceServer
	*internal.State
}

func NewServer(state *internal.State) pb.CommandServiceServer {
	return &service{State: state}
}

func (s *service) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log := s.Log.With("op", "create", "bin", req.Command.Bin, "args", req.Command.Args)
	if req.Command == nil {
		log.Error("no command found in request")
		return nil, fmt.Errorf("no command found in request")
	}

	bin, err := bin(req.Command.Bin)
	if err != nil {
		log.Error("unable to map bin", "err", err)
		return nil, fmt.Errorf("mapping bin: %w", err)
	}

	log.DebugContext(ctx, "building command")
	cmd := exec.CommandContext(ctx, bin, req.Command.Args...)
	cmd.Stdin = stdinReader(req.Command.Stdin)

	stderr, stdout := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	log.DebugContext(ctx, "running command", "cmd", cmd)
	if err = cmd.Run(); err != nil {
		log.ErrorContext(ctx, "failed to run command", "err", err)
	}

	log.InfoContext(ctx, "successfully ran command", "cmd", cmd)
	return &pb.CreateResponse{Result: &pb.Result{
		ExitCode: int32(cmd.ProcessState.ExitCode()),
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
	}}, nil
}

func (s *service) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log := s.Log.With("op", "delete", "bin", req.Create.Bin, "args", req.Create.Args)
	var cmd *exec.Cmd
	bin := pb.Bin_BIN_UNSPECIFIED

	switch req.Create.Bin {
	case pb.Bin_BIN_TEE:
		// All of the args to `tee` should be files
		toRemove := req.Create.Args
		args := prepend("-f", toRemove)

		log.DebugContext(ctx, "building command")
		cmd = exec.CommandContext(ctx, "rm", args...)
	default:
		log.ErrorContext(ctx, "unsupported bin")
		return nil, fmt.Errorf("unsupported bin: %s", req.Create.Bin)
	}

	if cmd == nil {
		log.InfoContext(ctx, "nothing to do")
		return &pb.DeleteResponse{}, nil
	}

	stderr, stdout := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	log.DebugContext(ctx, "running command", "cmd", cmd)
	if err := cmd.Run(); err != nil {
		log.ErrorContext(ctx, "failed to run command", "err", err)
		return &pb.DeleteResponse{}, nil
	}

	log.DebugContext(ctx, "describing operation")
	op := &pb.Operation{
		Command: &pb.Command{Bin: bin, Args: cmd.Args},
		Result: &pb.Result{
			ExitCode: int32(cmd.ProcessState.ExitCode()),
			Stdout:   stdout.String(),
			Stderr:   stderr.String(),
		},
	}

	log.InfoContext(ctx, "successfully ran command", "cmd", cmd)
	return &pb.DeleteResponse{Op: op}, nil
}

func bin(b pb.Bin) (string, error) {
	switch b {
	case pb.Bin_BIN_TEE:
		return "tee", nil
	}

	return "", fmt.Errorf("unrecognized bin: %s", b)
}
