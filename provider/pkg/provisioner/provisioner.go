package provisioner

import (
	"net"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provisioner/cmd"
	"google.golang.org/grpc"
)

func Serve(lis net.Listener) error {
	server := grpc.NewServer()
	command := NewCommandService()

	pb.RegisterCommandServiceServer(server, command)

	return server.Serve(lis)
}

func NewCommandService() pb.CommandServiceServer {
	return cmd.NewServer()
}
