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

// Tee implements baremetalv1alpha1.CommandServiceServer.
func (c *commandService) Tee(ctx context.Context, req *pb.TeeRequest) (*pb.TeeResponse, error) {
	return &pb.TeeResponse{
		State: &pb.State{},
	}, nil
}
