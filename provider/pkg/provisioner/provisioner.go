package provisioner

import (
	"log/slog"
	"net"

	"github.com/unmango/pulumi-baremetal/provider/pkg/internal/opts"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provisioner/command"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provisioner/meta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Provisioner interface {
	Serve() error
}

func New(lis net.Listener, o ...opt) Provisioner {
	options := &Options{
		logger:     slog.Default(),
		grpc:       []grpc.ServerOption{},
		reflection: false,
		whitelist:  []string{},
	}

	if err := opts.Apply(options, o...); err != nil {
		panic(err) // TODO: Update the signature to return an error
	}

	server := grpc.NewServer(options.grpc...)

	return &provisioner{
		Options:  *options,
		listener: lis,
		server:   server,
	}
}

func Serve(lis net.Listener) error {
	return New(lis).Serve()
}

type provisioner struct {
	Options

	listener net.Listener
	server   *grpc.Server
}

// Serve implements Provisioner.
func (p *provisioner) Serve() error {
	log := p.Options.logger
	if p.reflection {
		log.Debug("enabling reflection")
		reflection.Register(p.server)
	}

	log.Debug("registering services")

	command.NewServer(
		command.WithLogger(p.logger),
		command.WithWhitelist(p.whitelist),
	).Register(p.server)

	meta.NewServer(
		meta.WithLogger(p.logger),
	).Register(p.server)

	log.Debug("serving")
	return p.server.Serve(p.listener)
}
