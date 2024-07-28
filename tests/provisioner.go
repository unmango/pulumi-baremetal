package tests

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultProtocol string = "tcp"
	defaultAddress  string = "localhost"
)

type testProvisioner struct {
	ct   tc.Container
	port nat.Port
}

func NewTestProvisioner(ctx context.Context, logger io.Writer) (*testProvisioner, error) {
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

	return &testProvisioner{
		ct:   container,
		port: port,
	}, nil
}

func (p *testProvisioner) Start(ctx context.Context) error {
	return p.ct.Start(ctx)
}

func (p *testProvisioner) Stop(ctx context.Context) error {
	timeout := time.Duration(10 * time.Second)
	return p.ct.Stop(ctx, &timeout)
}
