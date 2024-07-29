package provider

import (
	"io"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"google.golang.org/grpc"
)

type Provisioner interface {
	io.Closer
	Cmd() pb.CommandServiceClient
}

type provisioner struct {
	cmd  pb.CommandServiceClient
	conn *grpc.ClientConn
}

var _ Provisioner = &provisioner{}

func NewProvisioner(conn *grpc.ClientConn) *provisioner {
	cmd := pb.NewCommandServiceClient(conn)

	return &provisioner{cmd: cmd, conn: conn}
}

// Cmd implements Provisioner
func (p *provisioner) Cmd() pb.CommandServiceClient {
	return p.cmd
}

// Close implements io.Closer.
func (p *provisioner) Close() error {
	return p.conn.Close()
}
