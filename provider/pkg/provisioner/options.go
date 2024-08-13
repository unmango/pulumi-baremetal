package provisioner

import (
	"crypto/tls"
	"fmt"
	"log/slog"

	"github.com/unmango/pulumi-baremetal/provider/pkg/internal/opts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Options struct {
	logger     *slog.Logger
	grpc       []grpc.ServerOption
	reflection bool
	whitelist  []string
}

type opt func(*Options) error

func WithLogger(logger *slog.Logger) opt {
	return func(o *Options) error {
		o.logger = logger
		return nil
	}
}

func WithGrpcOption(opt grpc.ServerOption) opt {
	return func(o *Options) error {
		o.grpcOption(opt)
		return nil
	}
}

func WithTLS(config *tls.Config) opt {
	return func(o *Options) error {
		o.tlsConfig(config)
		return nil
	}
}

func WithOptionalCertificates(caFile, certFile, keyFile string) opt {
	missingFile := caFile == "" || certFile == "" || keyFile == ""
	return opts.If(!missingFile, func(o *Options) error {
		certs, err := LoadCertificates(caFile, certFile, keyFile)
		if err != nil {
			return fmt.Errorf("failed to load certificates: %w", err)
		}

		o.tlsConfig(certs)
		return nil
	})
}

func WithReflection(enable bool) opt {
	return func(o *Options) error {
		o.reflection = enable
		return nil
	}
}

func WithWhitelist(whitelist []string) opt {
	return opts.Safe[opt](func(o *Options) {
		o.whitelist = whitelist
	})
}

func (o *Options) grpcOption(opt grpc.ServerOption) {
	o.grpc = append(o.grpc, opt)
}

func (o *Options) tlsConfig(config *tls.Config) {
	creds := credentials.NewTLS(config)
	o.grpcOption(grpc.Creds(creds))
}
