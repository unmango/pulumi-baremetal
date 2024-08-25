package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/config"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/provisioner"
)

type CommandArgs struct {
	Create     []string                      `pulumi:"create"`
	Update     []string                      `pulumi:"update,optional"`
	Delete     []string                      `pulumi:"delete,optional"`
	Triggers   []any                         `pulumi:"triggers,optional"`
	Connection *config.ProvisionerConnection `pulumi:"connection,optional"`
}

func (a CommandArgs) Provisioner(ctx context.Context) (*provisioner.Provisioner, error) {
	if a.Connection != nil {
		return provisioner.FromConnection(*a.Connection)
	} else {
		return provisioner.FromContext(ctx)
	}
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

	p, err := inputs.Provisioner(ctx)
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

	display := display(inputs.Create)
	log.DebugStatus("Sending exec request to provisioner")
	res, err := p.Exec(ctx, &pb.ExecRequest{
		Args: inputs.Create,
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

// Update implements infer.CustomUpdate.
func (Command) Update(ctx context.Context, id string, olds CommandState, news CommandArgs, preview bool) (CommandState, error) {
	log := logger.FromContext(ctx)
	state := CommandState{
		CommandArgs: olds.CommandArgs,
		ExitCode:    olds.ExitCode,
		Stdout:      olds.Stdout,
		Stderr:      olds.Stderr,
	}

	// What to do if the connection changes?
	p, err := news.Provisioner(ctx)
	if err != nil {
		log.Error("Failed creating provisioner")
		return state, fmt.Errorf("creating provisioner: %w", err)
	}

	if preview {
		if _, err = p.Ping(ctx, &pb.PingRequest{}); err != nil {
			log.WarningStatusf("Failed pinging provisioner: %s", err)
		}

		return state, nil
	}

	display := display(news.Update)
	log.DebugStatus("Sending exec request to provisioner")
	res, err := p.Exec(ctx, &pb.ExecRequest{
		Args: news.Update,
	})
	if err != nil {
		log.Errorf("command:%s %s", display, err)
		return state, fmt.Errorf("sending exec request: %w", err)
	}

	if res.Result.ExitCode > 0 {
		log.Error(display)
		return state, fmt.Errorf("exec failed: %s", res.Result)
	}

	state.CommandArgs = news
	state.ExitCode = int(res.Result.ExitCode)
	state.Stdout = res.Result.Stdout
	state.Stderr = res.Result.Stderr

	log.InfoStatus(display)
	return state, nil
}

// Delete implements infer.CustomDelete.
func (Command) Delete(ctx context.Context, id string, props CommandState) error {
	log := logger.FromContext(ctx)
	if len(props.Delete) == 0 {
		log.DebugStatus("No custom delete provided")
		return nil
	}

	p, err := props.Provisioner(ctx)
	if err != nil {
		log.Error("Failed creating provisioner")
		return fmt.Errorf("creating provisioner: %w", err)
	}

	display := display(props.Delete)
	log.DebugStatus("Sending exec request to provisioner")
	res, err := p.Exec(ctx, &pb.ExecRequest{
		Args: props.Delete,
	})
	if err != nil {
		log.Errorf("command:%s %s", display, err)
		return fmt.Errorf("sending exec request: %w", err)
	}

	if res.Result.ExitCode > 0 {
		log.Error(display)
		return fmt.Errorf("exec failed: %s", res.Result)
	}

	log.InfoStatus(display)
	return nil
}

var _ = (infer.CustomCreate[CommandArgs, CommandState])((*Command)(nil))
var _ = (infer.CustomUpdate[CommandArgs, CommandState])((*Command)(nil))
var _ = (infer.CustomDelete[CommandState])((*Command)(nil))

func display(args []string) string {
	return strings.Join(args, " ")
}
