package cmd

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type service struct {
	pb.UnimplementedCommandServiceServer
	log *slog.Logger
}

func NewServer() pb.CommandServiceServer {
	return &service{log: slog.Default()}
}

// Command implements baremetalv1alpha1.CommandServiceServer.
func (c *service) Command(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	switch req.Op {
	case pb.Op_OP_CREATE:
		return c.create(ctx, req)
	case pb.Op_OP_DELETE:
		return c.delete(ctx, req)
	}

	return nil, fmt.Errorf("unsupported op: %s", req.Op)
}

func (c *service) create(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	log := c.log.With("op", req.Op)
	bin, err := getBin(req.Command)
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(ctx, bin, req.Args...)
	log = log.With("bin", bin, "args", req.Args)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdin = strings.NewReader(req.Stdin)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	log.Info("executing command")
	err = cmd.Run()
	if err != nil {
		log.Error("command failed", "err", err)
	}

	return &pb.CommandResponse{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}, nil
}

func (c *service) delete(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	log := c.log.With("op", req.Op)
	bin := "rm"

	cmd := exec.CommandContext(ctx, bin, req.Args...)
	log = log.With("bin", bin, "args", req.Args)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdin = strings.NewReader(req.Stdin)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	log.Info("executing command")
	err := cmd.Run()
	if err != nil {
		log.Error("command failed", "err", err)
	}

	return &pb.CommandResponse{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}, nil
}

func getBin(cmd pb.Command) (string, error) {
	switch cmd {
	case pb.Command_COMMAND_TEE:
		return "tee", nil
	}

	return "", fmt.Errorf("unrecognized command: %s", cmd)
}
