package services

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
	"github.com/unmango/pulumi-baremetal/tests/util"
)

const (
	defaultProtocol string = "tcp4"
)

type Provisioner struct {
	Host

	Address string
	Port    string
	Certs   *util.CertBundle
}

func NewProvisioner(
	port string,
	clientCa *tlscert.Certificate,
	logger io.Writer,
) (*Provisioner, error) {
	certs, err := util.NewCertBundle("ca", "provisioner")
	if err != nil {
		return &Provisioner{}, err
	}

	certDir := "/etc/baremetal/pki"
	clientCaPath := path.Join(certDir, "client-ca.pem")
	certPath := path.Join(certDir, "cert.pem")
	keyPath := path.Join(certDir, "key.pem")

	image := "baremetal-provisioner:test"
	if env, ok := os.LookupEnv("PROVISIONER_IMAGE"); ok {
		image = env
	}

	req := tc.GenericContainerRequest{
		Logger: util.NewLogger(logger),
		ContainerRequest: tc.ContainerRequest{
			Image: image,
			Files: []tc.ContainerFile{
				{ContainerFilePath: clientCaPath, Reader: bytes.NewReader(clientCa.Bytes)},
				{ContainerFilePath: certPath, Reader: bytes.NewReader(certs.Cert.Bytes)},
				{ContainerFilePath: keyPath, Reader: bytes.NewReader(certs.Cert.KeyBytes)},
			},
			Env: map[string]string{
				"GRPC_GO_LOG_SEVERITY_LEVEL": "info",
			},
			Cmd: []string{
				"--network", defaultProtocol,
				"--address", fmt.Sprintf("%s:%s", "0.0.0.0", port),
				"--client-ca-file", clientCaPath,
				"--cert-file", certPath,
				"--key-file", keyPath,
				"--verbose",
				"--whitelist", "perl",
			},
			ExposedPorts: []string{port},
			WaitingFor:   wait.ForExposedPort(),
			LogConsumerCfg: &tc.LogConsumerConfig{
				Consumers: []tc.LogConsumer{util.LogToWriter(logger)},
			},
		},
	}

	return &Provisioner{
		Host{req, nil},
		"",
		port,
		certs,
	}, nil
}

// CertBundle implements TestProvisioner.
func (p *Provisioner) Ca() *tlscert.Certificate {
	return p.Certs.Ca
}

// ConnectionDetails implements TestProvisioner.
func (p *Provisioner) ConnectionDetails(ctx context.Context) (string, string, error) {
	host, port, err := p.Host.ConnectionDetails(ctx, p.Port)
	if err != nil {
		return "", "", err
	}

	return host, port.Port(), nil
}
