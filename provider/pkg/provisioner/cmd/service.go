package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type service struct {
	pb.UnimplementedCommandServiceServer
}

func NewServer() pb.CommandServiceServer {
	return &service{}
}

// Command implements baremetalv1alpha1.CommandServiceServer.
func (c *service) Command(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	switch req.Command {
	case pb.Command_COMMAND_TEE:
		return tee(ctx, req)
	}

	return nil, fmt.Errorf("unrecognized command: %s", req.Command)
}
