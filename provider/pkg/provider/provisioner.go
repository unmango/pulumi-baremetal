package provider

import (
	"io"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"google.golang.org/grpc"
)

type provisioner struct {
	Cmd  pb.CommandServiceClient
	conn *grpc.ClientConn
}

var _ io.Closer = &provisioner{}

func NewProvisioner(conn *grpc.ClientConn) *provisioner {
	cmd := pb.NewCommandServiceClient(conn)

	return &provisioner{Cmd: cmd, conn: conn}
}

// Close implements io.Closer.
func (p *provisioner) Close() error {
	return p.conn.Close()
}
