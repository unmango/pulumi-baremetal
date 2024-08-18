package provisioner

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/config"
	"google.golang.org/grpc"
)

type provisioner struct {
	pb.CommandServiceClient
	pb.MetaServiceClient
	conn *grpc.ClientConn
}

// Close implements io.Closer.
func (p *provisioner) Close() error {
	return p.conn.Close()
}

func FromContext(ctx context.Context) (*provisioner, error) {
	config := infer.GetConfig[config.Config](ctx)
	return FromConfig(config)
}

func FromConfig(config config.Config) (*provisioner, error) {
	conn, err := config.NewGrpcClient()
	if err != nil {
		return nil, err
	}

	return new(conn), nil
}

func new(conn *grpc.ClientConn) *provisioner {
	cmd := pb.NewCommandServiceClient(conn)
	meta := pb.NewMetaServiceClient(conn)
	return &provisioner{cmd, meta, conn}
}
