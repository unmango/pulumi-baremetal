package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type CommandArgs struct {
	Args     []string `pulumi:"args"`
	Triggers []any    `pulumi:"triggers,optional"`
}

type Command struct{}

type CommandState struct {
	CommandArgs

	ExitCode int    `pulumi:"exitCode"`
	Stdout   string `pulumi:"stdout"`
	Stderr   string `pulumi:"stderr"`
}

// Create implements infer.CustomCreate.
func (Command) Create(ctx context.Context, name string, inputs CommandArgs, preview bool) (string, CommandState, error) {
	log := logger.FromContext(ctx)
	state := CommandState{}

	p, err := provisioner.FromContext(ctx)
	if err != nil {
		log.Error("Failed creating provisioner")
		return name, state, fmt.Errorf("creating provisioner: %w", err)
	}

	if preview {
		if _, err = p.Ping(ctx, &pb.PingRequest{}); err != nil {
			log.WarningStatusf("Failed pinging provisioner: %s", err)
		}

		return name, state, nil
	}

	display := display(inputs.Args)
	log.DebugStatus("Sending exec request to provisioner")
	res, err := p.Exec(ctx, &pb.ExecRequest{
		Args: inputs.Args,
	})
	if err != nil {
		log.Errorf("command:%s %s", display, err)
		return name, state, fmt.Errorf("sending exec request: %w", err)
	}

	if res.Result.ExitCode > 0 {
		log.Error(display)
		return name, state, fmt.Errorf("exec failed: %s", res.Result)
	}

	state.CommandArgs = inputs
	state.ExitCode = int(res.Result.ExitCode)
	state.Stdout = res.Result.Stdout
	state.Stderr = res.Result.Stderr

	log.InfoStatus(display)
	return name, state, nil
}

var _ = (infer.CustomCreate[CommandArgs, CommandState])((*Command)(nil))

func display(args []string) string {
	return strings.Join(args, " ")
}