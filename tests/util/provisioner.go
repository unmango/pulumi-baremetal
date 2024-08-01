package util

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultProtocol string = "tcp"
	defaultAddress  string = "localhost"
)

type TestProvisioner interface {
	TestHost

	ConfigureProvider(context.Context, integration.Server) error
}

type provisioner struct {
	host
	port string
}

func NewProvisioner(port string, logger io.Writer) (TestProvisioner, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	req := tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			FromDockerfile: tc.FromDockerfile{
				Context:    path.Clean(path.Join(cwd, "..")),
				Dockerfile: path.Join("provider", "cmd", "provisioner", "Dockerfile"),
			},
			Cmd: []string{
				"--network", defaultProtocol,
				"--address", fmt.Sprintf("%s:%s", "0.0.0.0", port),
				"--verbose",
			},
			ExposedPorts: []string{port},
			WaitingFor:   wait.ForExposedPort(),
			LogConsumerCfg: &tc.LogConsumerConfig{
				Consumers: []tc.LogConsumer{LogToWriter(logger)},
			},
		},
	}

	return &provisioner{host{req, nil}, port}, nil
}

func (prov provisioner) ConfigureProvider(ctx context.Context, server integration.Server) error {
	ip, err := prov.Ip(ctx)
	if err != nil {
		return err
	}

	if ip == "" {
		return errors.New("container returned empty ip")
	}

	port := prov.port
	return server.Configure(p.ConfigureRequest{
		Args: resource.PropertyMap{
			"address": resource.NewStringProperty(ip),
			"port":    resource.NewStringProperty(port),
		},
	})
}
