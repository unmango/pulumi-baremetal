package provider

import (
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address string `pulumi:"address"`
	Port    string `pulumi:"port,optional"`
}

func (c Config) NewGrpcClient() (*grpc.ClientConn, error) {
	parts := []string{}
	if c.Address != "" {
		parts = append(parts, c.Address)
	}
	if c.Port != "" {
		parts = append(parts, c.Port)
	}

	// Why must I over-engineer things
	target := strings.Join(parts, ":")
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
