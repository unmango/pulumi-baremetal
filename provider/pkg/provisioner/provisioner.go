package provisioner

import (
	"log/slog"
	"net"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal/opts"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provisioner/cmd"
	"google.golang.org/grpc"
)

type Options struct {
	logger *slog.Logger
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
	p.RegisterCommandServiceServer(cmd.NewServer(p.State))

	return p.server.Serve(p.listener)
}

func New(lis net.Listener, o ...opt) Provisioner {
	options := &Options{}
	opts.Apply(options, o...)

	return &provisioner{
		State:    options.State(),
		listener: lis,
		server:   grpc.NewServer(),
	}
}

func WithLogger(logger *slog.Logger) opt {
	return func(o *Options) error {
		o.logger = logger
		return nil
	}
}

func Serve(lis net.Listener) error {
	return New(lis).Serve()
}

func (o *Options) State() *internal.State {
	return &internal.State{Log: o.logger}
}

func (p *provisioner) RegisterCommandServiceServer(srv pb.CommandServiceServer) {
	pb.RegisterCommandServiceServer(p.server, srv)
}
