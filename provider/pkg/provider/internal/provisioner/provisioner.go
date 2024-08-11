package provisioner

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider"
	"google.golang.org/grpc"
)

type provisioner struct {
	pb.CommandServiceClient
	conn *grpc.ClientConn
}

// Close implements io.Closer.
func (p *provisioner) Close() error {
	return p.conn.Close()
}

func FromContext(ctx context.Context) (*provisioner, error) {
	config := infer.GetConfig[provider.Config](ctx)
	return FromConfig(config)
}

func FromConfig(config provider.Config) (*provisioner, error) {
	conn, err := config.NewGrpcClient()
	if err != nil {
		return nil, err
	}

	return new(conn), nil
}

func new(conn *grpc.ClientConn) *provisioner {
	cmd := pb.NewCommandServiceClient(conn)
	return &provisioner{cmd, conn}
}
