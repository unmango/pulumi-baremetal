package provisioner

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func LoadCertificates(caPath, certPath, keyPath string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed loading keypair: %w", err)
	}

	ca := x509.NewCertPool()
	caData, err := os.ReadFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("failed reading ca file: %w", err)
	}
	if ok := ca.AppendCertsFromPEM(caData); ok {
		return nil, fmt.Errorf("unable to append ca data from file %s", caPath)
	}

	return &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    ca,
	}, nil
}
