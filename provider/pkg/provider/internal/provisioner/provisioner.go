package provisioner

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/config"
	"google.golang.org/grpc"
)

type Provisioner struct {
	pb.CommandServiceClient
	pb.MetaServiceClient
	conn *grpc.ClientConn
}

// Close implements io.Closer.
func (p *Provisioner) Close() error {
	return p.conn.Close()
}

func FromContext(ctx context.Context) (*Provisioner, error) {
	config := infer.GetConfig[config.Config](ctx)
	return FromConnection(config.ProvisionerConnection)
}

func FromConnection(conn config.ProvisionerConnection) (*Provisioner, error) {
	client, err := conn.NewGrpcClient()
	if err != nil {
		return nil, err
	}

	return new(client), nil
}

func new(conn *grpc.ClientConn) *Provisioner {
	cmd := pb.NewCommandServiceClient(conn)
	meta := pb.NewMetaServiceClient(conn)
	return &Provisioner{cmd, meta, conn}
}
