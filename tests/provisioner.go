package tests

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

type testProvisioner struct {
	ct   tc.Container
	port nat.Port
}

func NewTestProvisioner(ctx context.Context, logger io.Writer) (testProvisioner, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return testProvisioner{}, err
	}

	port, err := nat.NewPort(defaultProtocol, "6969")
	if err != nil {
		return testProvisioner{}, err
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
		return testProvisioner{}, err
	}

	return testProvisioner{
		ct:   container,
		port: port,
	}, nil
}

func (p testProvisioner) Start(ctx context.Context) error {
	return p.ct.Start(ctx)
}

func (p testProvisioner) Stop(ctx context.Context) error {
	timeout := time.Duration(10 * time.Second)
	return p.ct.Stop(ctx, &timeout)
}

func (p testProvisioner) Exec(ctx context.Context, args ...string) error {
	code, output, err := p.ct.Exec(ctx, args)
	if err != nil {
		return err
	}

	out, err := io.ReadAll(output)
	if err != nil {
		return err
	}

	if code != 0 {
		return fmt.Errorf("unexpected return code: %d, output: %s", code, out)
	}

	return nil
}

func (prov testProvisioner) ConfigureProvider(ctx context.Context, server integration.Server) error {
	ip, err := prov.ct.ContainerIP(ctx)
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
