package provider

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address string `pulumi:"address"`
	Port    string `pulumi:"port,optional"`
}

func (c Config) NewGrpcClient() (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%s:%s", c.Address, c.Port)
	return grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

func (c Config) NewProvisioner() (Provisioner, error) {
	conn, err := c.NewGrpcClient()
	if err != nil {
		return nil, err
	}

	return NewProvisioner(conn), nil
}
