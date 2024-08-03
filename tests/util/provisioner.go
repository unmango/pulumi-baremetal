package util

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/mdelapenya/tlscert"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultProtocol string = "tcp"
	defaultAddress  string = "localhost"
)

type TestProvisioner interface {
	TestHost

	Ca() *tlscert.Certificate
	ConnectionDetails(context.Context) (address, port string, err error)
}

type provisioner struct {
	host
	port   string
	bundle *CertBundle
}

func NewProvisioner(
	port string,
	clientCa *tlscert.Certificate,
	logger io.Writer,
) (TestProvisioner, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	certs, err := NewCertBundle("ca", "provisioner")
	if err != nil {
		return nil, err
	}

	certDir := "/etc/baremetal/pki"
	clientCaPath := path.Join(certDir, "client-ca.pem")
	certPath := path.Join(certDir, "cert.pem")
	keyPath := path.Join(certDir, "key.pem")

	req := tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			FromDockerfile: tc.FromDockerfile{
				Context:    path.Clean(path.Join(cwd, "..")),
				Dockerfile: path.Join("provider", "cmd", "provisioner", "Dockerfile"),
			},
			Files: []tc.ContainerFile{
				{
					ContainerFilePath: clientCaPath,
					Reader:            bytes.NewReader(clientCa.Bytes),
				},
				{
					ContainerFilePath: certPath,
					Reader:            bytes.NewReader(certs.Cert.Bytes),
				},
				{
					ContainerFilePath: keyPath,
					Reader:            bytes.NewReader(certs.Cert.KeyBytes),
				},
			},
			Cmd: []string{
				"--network", defaultProtocol,
				"--address", fmt.Sprintf("%s:%s", "0.0.0.0", port),
				"--client-ca-file", clientCaPath,
				"--cert-file", certPath,
				"--key-file", keyPath,
				"--verbose",
			},
			ExposedPorts: []string{port},
			WaitingFor:   wait.ForExposedPort(),
			LogConsumerCfg: &tc.LogConsumerConfig{
				Consumers: []tc.LogConsumer{LogToWriter(logger)},
			},
		},
	}

	return &provisioner{host{req, nil}, port, certs}, nil
}

// CertBundle implements TestProvisioner.
func (p *provisioner) Ca() *tlscert.Certificate {
	return p.bundle.Ca
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
