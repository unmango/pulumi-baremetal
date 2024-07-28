package cmd

import (
	"context"
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

func tee(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	stdout := fmt.Sprintf("op: %s, cmd: %s, args: %#v, flags: %#v",
		req.Op, req.Command, req.Args, req.Flags,
	)

	return &pb.CommandResponse{
		Stdout: stdout,
	}, nil
}
