package provider

import (
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"google.golang.org/grpc"
)

type provisioner struct {
	Cmd  pb.CommandServiceClient
	conn *grpc.ClientConn
}
