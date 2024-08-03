package tests

import (
	"github.com/mdelapenya/tlscert"
)

type CertBundle struct {
	Ca   *tlscert.Certificate
	Cert *tlscert.Certificate
}

func ServerCerts() (*CertBundle, error) {
	ca := tlscert.SelfSignedCA("test-ca")

	req := tlscert.NewRequest("test-cert")
	req.Parent = ca
	cert := tlscert.SelfSignedFromRequest(req)

	return &CertBundle{ca, cert}, nil
}
