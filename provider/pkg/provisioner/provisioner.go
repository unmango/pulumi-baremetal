package provisioner

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal/opts"
	cmd "github.com/unmango/pulumi-baremetal/provider/pkg/provisioner/command"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type Options struct {
	logger     *slog.Logger
	grpc       []grpc.ServerOption
	reflection bool
}

type opt func(*Options) error

type Provisioner interface {
	Serve() error
}

type provisioner struct {
	internal.State
	listener   net.Listener
	server     *grpc.Server
	reflection bool
}

func New(lis net.Listener, o ...opt) Provisioner {
	options := &Options{
		slog.Default(),
		[]grpc.ServerOption{},
		false,
	}

	if err := opts.Apply(options, o...); err != nil {
		panic(err) // TODO: Update the signature to return an error
	}

	return &provisioner{
		State:      options.state(),
		listener:   lis,
		server:     grpc.NewServer(options.grpc...),
		reflection: options.reflection,
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

func WithReflection(enable bool) opt {
	return func(o *Options) error {
		o.reflection = enable
		return nil
	}
}

// Serve implements Provisioner.
func (p *provisioner) Serve() error {
	if p.reflection {
		p.Log.Debug("enabling reflection")
		reflection.Register(p.server)
	}

	p.Log.Debug("registering command server")
	p.registerCommandService(cmd.NewServer(p.State))

	return p.server.Serve(p.listener)
}

func Serve(lis net.Listener) error {
	return New(lis).Serve()
}

func (o *Options) state() internal.State {
	return internal.State{Log: o.logger}
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
