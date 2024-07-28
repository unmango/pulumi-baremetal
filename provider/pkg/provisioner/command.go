package provisioner

import (
	"context"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type commandService struct {
	pb.UnimplementedCommandServiceServer
}

func NewCommandService() pb.CommandServiceServer {
	return &commandService{}
}

// Command implements baremetalv1alpha1.CommandServiceServer.
func (c *commandService) Command(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	return &pb.CommandResponse{
		Stdout: "Hi friend",
		Stderr: "No errors here",
	}, nil
}
