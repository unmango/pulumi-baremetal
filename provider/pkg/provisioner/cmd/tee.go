package cmd

import (
	"context"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

func (s *service) deleteTee(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	log := s.logger(req)

	// All args should be files that were created if the command succeeded.
	// We might want stricter semantics around this for safety in the future.
	log.Debug("attempting to remove files", "files", req.Args)
	return run(newCommand(ctx, "rm", prepend("-f", req.Args)), log)
}
