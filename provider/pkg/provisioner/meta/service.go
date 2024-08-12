package cmd

import (
	"context"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal"
)

type service struct {
	pb.UnimplementedMetaServiceServer
	internal.State
}

func NewServer(state internal.State) pb.MetaServiceServer {
	log := state.Log.With("service", "command")
	return &service{State: state.WithLogger(log)}
}

func (s *service) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	s.Log.DebugContext(ctx, "ping")
	return &pb.PingResponse{}, nil
}