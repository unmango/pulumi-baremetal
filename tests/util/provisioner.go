package util

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/docker/go-connections/nat"
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
	Ctr() tc.Container
}

type provisioner struct {
	host
	port nat.Port
}

func NewTestProvisioner(ctx context.Context, logger io.Writer) (TestProvisioner, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	port, err := nat.NewPort(defaultProtocol, "6969")
	if err != nil {
		return nil, err
	}

	container, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			FromDockerfile: tc.FromDockerfile{
				Context:    path.Clean(path.Join(cwd, "..")),
				Dockerfile: path.Join("provider", "cmd", "provisioner", "Dockerfile"),
			},
			Cmd: []string{
				"--network", defaultProtocol,
				"--address", fmt.Sprintf("%s:%d", "0.0.0.0", port.Int()),
				"--verbose",
			},
			ExposedPorts: []string{port.Port()},
			WaitingFor:   wait.ForListeningPort(port),
			LogConsumerCfg: &tc.LogConsumerConfig{
				Consumers: []tc.LogConsumer{LogToWriter(logger)},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &provisioner{
		host: host{ctr: container},
		port: port,
	}, nil
}

func (p provisioner) Ctr() tc.Container {
	return p.ctr
}

func (p provisioner) Start(ctx context.Context) error {
	return p.ctr.Start(ctx)
}

func (p provisioner) Stop(ctx context.Context) error {
	timeout := time.Duration(10 * time.Second)
	return p.ctr.Stop(ctx, &timeout)
}

func (prov provisioner) ConfigureProvider(ctx context.Context, server integration.Server) error {
	ip, err := prov.ctr.ContainerIP(ctx)
	if err != nil {
		return err
	}

	if ip == "" {
		return errors.New("container returned empty ip")
	}

	port := prov.port.Int()

	return server.Configure(p.ConfigureRequest{
		Args: resource.PropertyMap{
			"address": resource.NewStringProperty(ip),
			"port":    resource.NewNumberProperty(float64(port)),
		},
	})
}
