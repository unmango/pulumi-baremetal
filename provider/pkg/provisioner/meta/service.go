package meta

import (
	"context"
	"log/slog"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal/opts"
	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedMetaServiceServer
	Log *slog.Logger
}

type opt func(*service) error

func NewServer(options ...opt) *service {
	s := &service{Log: slog.Default()}
	if err := opts.Apply(s, options...); err != nil {
		panic(err) // TODO
	}

	return s
}

func WithLogger(logger *slog.Logger) opt {
	return opts.Safe[opt](func(s *service) {
		s.Log = logger
	})
}

func (s *service) Register(server *grpc.Server) {
	pb.RegisterMetaServiceServer(server, s)
}

func (s *service) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	s.Log.DebugContext(ctx, "ping")
	return &pb.PingResponse{Message: "pong"}, nil
}

func (s *service) Version(ctx context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {
	s.Log.DebugContext(ctx, "version")
	return &pb.VersionResponse{Version: provider.Version}, nil
}
