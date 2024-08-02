package provisioner

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal/opts"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provisioner/cmd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Options struct {
	logger *slog.Logger
	grpc   []grpc.ServerOption
}

type opt func(*Options) error

type Provisioner interface {
	Serve() error
}

type provisioner struct {
	*internal.State
	listener net.Listener
	server   *grpc.Server
}

// Serve implements Provisioner.
func (p *provisioner) Serve() error {
	p.registerCommandService(cmd.NewServer(p.State))

	return p.server.Serve(p.listener)
}

func New(lis net.Listener, o ...opt) Provisioner {
	options := &Options{slog.Default(), []grpc.ServerOption{}}
	err := opts.Apply(options, o...)
	if err != nil {
		panic(err) // TODO: Update the signature to return an error
	}

	return &provisioner{
		State:    options.state(),
		listener: lis,
		server:   grpc.NewServer(options.grpc...),
	}
}

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

func Serve(lis net.Listener) error {
	return New(lis).Serve()
}

func (o *Options) state() *internal.State {
	return &internal.State{Log: o.logger}
}

func (o *Options) grpcOption(opt grpc.ServerOption) {
	o.grpc = append(o.grpc, opt)
}

func (o *Options) tlsConfig(config *tls.Config) {
	creds := credentials.NewTLS(config)
	o.grpcOption(grpc.Creds(creds))
}

func (p *provisioner) registerCommandService(srv pb.CommandServiceServer) {
	pb.RegisterCommandServiceServer(p.server, srv)
}
