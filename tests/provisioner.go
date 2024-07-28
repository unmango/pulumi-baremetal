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

const protocol string = "tcp"

type testProvisioner struct {
	ct   tc.Container
	port nat.Port
}

func NewTestProvisioner(ctx context.Context, logger io.Writer) (*testProvisioner, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	port, err := nat.NewPort(protocol, "6969")
	if err != nil {
		return nil, err
	}

	repoRoot := path.Clean(path.Join(cwd, ".."))
	address := fmt.Sprintf("localhost:%d", port.Int())

	container, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			FromDockerfile: tc.FromDockerfile{
				Context:    repoRoot,
				Dockerfile: path.Join("provider", "cmd", "provisioner", "Dockerfile"),
			},
			Cmd: []string{
				"--network", protocol,
				"--address", address,
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

	return &testProvisioner{ct: container, port: port}, nil
}

func (p *testProvisioner) Stop(ctx context.Context) error {
	timeout := time.Duration(10 * time.Second)
	return p.ct.Stop(ctx, &timeout)
}
