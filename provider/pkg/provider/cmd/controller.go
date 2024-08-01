package cmd

import (
	"context"
	"fmt"
	"maps"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal/opts"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type controller struct {
	op  pb.Op
	cmd pb.Command
}

type options struct {
	args  *[]string
	flags *map[string]*pb.Flag
	stdin *string
}

type opt = func(*options) error

func (c *controller) run(ctx context.Context, o ...opt) (*pb.CommandResponse, error) {
	req := &options{}
	if err := opts.Apply(req, o...); err != nil {
		return nil, fmt.Errorf("applying options: %w", err)
	}

	log := logger.FromContext(ctx)
	p, err := provisioner.FromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating provisioner: %w", err)
	}

	args := []string{}
	if req.args != nil {
		args = *req.args
	}

	flags := map[string]*pb.Flag{}
	if req.flags != nil {
		flags = *req.flags
	}

	log.Debug("sending command request to provisioner")
	res, err := p.Command(ctx, &pb.CommandRequest{
		Op:      c.op,
		Command: c.cmd,
		Args:    args,
		Flags:   flags,
		Stdin:   req.stdin,
	})
	if err != nil {
		return nil, fmt.Errorf("command request: %w", err)
	}

	return res, nil
}

func withArgs(args []string) opt {
	return func(o *options) error {
		o.args = &args
		return nil
	}
}

func withFlags(flags map[string]*pb.Flag) opt {
	return func(o *options) error {
		if o.flags == nil {
			o.flags = &map[string]*pb.Flag{}
		}

		maps.Copy(*o.flags, flags)

		return nil
	}
}

func withFlag(name string, present bool) opt {
	return withFlags(map[string]*pb.Flag{
		name: {Present: present},
	})
}

func withStdin(stdin string) opt {
	return func(o *options) error {
		o.stdin = &stdin
		return nil
	}
}
