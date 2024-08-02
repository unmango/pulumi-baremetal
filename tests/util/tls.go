package util

import (
	"errors"

	"github.com/mdelapenya/tlscert"
)

type CertBundle struct {
	Ca   *tlscert.Certificate
	Cert *tlscert.Certificate
}

func NewCertBundle(caHost, certHost string) (*CertBundle, error) {
	ca := tlscert.SelfSignedCA(caHost)

	req := tlscert.NewRequest(certHost)
	req.Parent = ca

	cert := tlscert.SelfSignedFromRequest(req)

	if ca == nil || cert == nil {
		return nil, errors.New("failed to generate bundle")
	}

	return &CertBundle{ca, cert}, nil
}
