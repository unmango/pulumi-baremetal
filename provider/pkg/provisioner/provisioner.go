package provisioner

import (
	"crypto/tls"
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
	_ = opts.Apply(options, o...)

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

func WithTLS(config *tls.Config) opt {
	return func(o *Options) error {
		creds := credentials.NewTLS(config)
		o.grpc = append(o.grpc, grpc.Creds(creds))
		return nil
	}
}

func Serve(lis net.Listener) error {
	return New(lis).Serve()
}

func (o *Options) state() *internal.State {
	return &internal.State{Log: o.logger}
}

func (p *provisioner) registerCommandService(srv pb.CommandServiceServer) {
	pb.RegisterCommandServiceServer(p.server, srv)
}
