package util

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"

	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultProtocol string = "tcp"
	defaultAddress  string = "localhost"
)

type TestProvisioner interface {
	TestHost

	ConnectionDetails(context.Context) (address, port string, err error)
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

// ConnectionDetails implements TestProvisioner.
func (p *provisioner) ConnectionDetails(ctx context.Context) (address string, port string, err error) {
	address, err = p.Ip(ctx)
	if err != nil {
		return
	}

	port = p.port
	return
}
