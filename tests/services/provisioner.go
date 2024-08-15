package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/docker/go-connections/nat"
	"github.com/mdelapenya/tlscert"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

const (
	defaultProtocol string = "tcp4"
)

type TestProvisioner interface {
	TestHost

	Ca() *tlscert.Certificate
	ConnectionDetails(context.Context) (address, port string, err error)
}

type provisioner struct {
	host
	port   string
	bundle *util.CertBundle
}

func NewProvisioner(
	port string,
	clientCa *tlscert.Certificate,
	logger io.Writer,
) (TestProvisioner, error) {
	certs, err := util.NewCertBundle("ca", "provisioner")
	if err != nil {
		return nil, err
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

	return &provisioner{host{req, nil}, port, certs}, nil
}

// CertBundle implements TestProvisioner.
func (p *provisioner) Ca() *tlscert.Certificate {
	return p.bundle.Ca
}

// ConnectionDetails implements TestProvisioner.
func (p *provisioner) ConnectionDetails(ctx context.Context) (address string, port string, err error) {
	ctr, err := p.Ctr(ctx)
	if err != nil {
		return
	}

	address, err = ctr.Host(ctx)
	if err != nil {
		return
	}

	np, err := ctr.MappedPort(ctx, nat.Port(p.port))
	if err != nil {
		return
	}

	port = np.Port()
	return
}
