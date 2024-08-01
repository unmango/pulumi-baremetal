package cmd

import (
	"context"
	"fmt"
	"log/slog"
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

// Command implements baremetalv1alpha1.CommandServiceServer.
func (s *service) Command(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	s.Log.Debug("parsing command op")
	switch req.Op {
	case pb.Op_OP_CREATE:
		return s.create(ctx, req)
	case pb.Op_OP_DELETE:
		return s.delete(ctx, req)
	}

	s.Log.Error("unsupported op", "op", req.Op)
	return nil, fmt.Errorf("unsupported op: %s", req.Op)
}

func (s *service) create(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	log := s.logger(req)
	bin, err := getBin(req.Command)
	if err != nil {
		log.Error("getting bin from command", "err", err)
		return nil, fmt.Errorf("getting bin from command: %w", err)
	}

	log = log.With("bin", bin, "args", req.Args)
	log.Debug("creating command with context")
	cmd := exec.CommandContext(ctx, bin, req.Args...)

	return run(cmd, log)
}

func (s *service) delete(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	switch req.Command {
	case pb.Command_COMMAND_TEE:
		return s.deleteTee(ctx, req)
	}

	return nil, fmt.Errorf("unsupported command: %s", req.Command)
}

func getBin(cmd pb.Command) (string, error) {
	switch cmd {
	case pb.Command_COMMAND_TEE:
		return "tee", nil
	}

	return "", fmt.Errorf("unrecognized command: %s", cmd)
}

func (s *service) logger(req *pb.CommandRequest) *slog.Logger {
	return s.Log.With("op", req.Op, "cmd", req.Command)
}
