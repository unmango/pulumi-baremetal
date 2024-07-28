package provider

import (
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"google.golang.org/grpc"
)

type Config struct {
	Address string `pulumi:"address"`
	Port    int    `pulumi:"port"`
}

func (c Config) ProvisionerClient() (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%s:%d", c.Address, c.Port)
	return grpc.NewClient(target)
}

func (c Config) provisioner() (*provisioner, error) {
	conn, err := c.ProvisionerClient()
	if err != nil {
		return nil, err
	}

	return &provisioner{
		Cmd:  pb.NewCommandServiceClient(conn),
		conn: conn,
	}, nil
}
