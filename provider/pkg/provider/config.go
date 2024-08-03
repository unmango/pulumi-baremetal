package provider

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address string `pulumi:"address"`
	Port    string `pulumi:"port,optional"`
	CaPem   string `pulumi:"caPem,optional"`
	CertPem string `pulumi:"certPem,optional"`
	KeyPem  string `pulumi:"keyPem,optional"`
}

func (c Config) NewGrpcClient() (*grpc.ClientConn, error) {
	target := c.Address
	if c.Port != "" {
		target = target + ":" + c.Port
	}

	creds, err := c.TransportCredentials()
	if err != nil {
		return nil, err
	}

	return grpc.NewClient(target,
		grpc.WithTransportCredentials(creds),
	)
}

func (c Config) TransportCredentials() (credentials.TransportCredentials, error) {
	if c.CaPem == "" && c.CertPem == "" && c.KeyPem == "" {
		return insecure.NewCredentials(), nil
	}

	if c.CaPem != "" && c.CertPem != "" && c.KeyPem != "" {
		cert, err := tls.X509KeyPair([]byte(c.CertPem), []byte(c.KeyPem))
		if err != nil {
			return nil, fmt.Errorf("failed to parse X509 key pair: %w", err)
		}

		ca := x509.NewCertPool()
		if ok := ca.AppendCertsFromPEM([]byte(c.CaPem)); !ok {
			return nil, errors.New("failed to append ca cert")
		}

		return credentials.NewTLS(&tls.Config{
			ServerName:   "provisioner",
			Certificates: []tls.Certificate{cert},
			RootCAs:      ca,
		}), nil
	}

	return nil, errors.New("caPem, certPem, and keyPem must all be set together")
}
