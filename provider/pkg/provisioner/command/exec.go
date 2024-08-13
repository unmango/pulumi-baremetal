package command

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	cmd "github.com/unmango/pulumi-baremetal/provider/pkg/command"
)

// I'm toying with extracting the core execution logic out of `Create` so that i.e. `Update` and `Delete`
// aren't simply delegating calls to `Create`. Semantically this is correct but it feels awkward
// executing a gRPC endpoint from within another endpoint on the same service.

// I'm not sure how I want to handle the filesystem expectations yet, so this function exists
// but `Update` and `Delete` will continue to delgate to `Create` for now.

func execute(ctx context.Context, command *pb.Command, log *slog.Logger) (*pb.Result, error) {
	if command == nil {
		log.ErrorContext(ctx, "no command found in request")
		return nil, fmt.Errorf("no command found in request")
	}

	bin, err := cmd.BinValue(command.Bin)
	if err != nil {
		log.ErrorContext(ctx, "unable to map bin", "err", err)
		return nil, fmt.Errorf("mapping bin: %w", err)
	}

	log = log.With("bin", bin, "args", command.Args)
	log.DebugContext(ctx, "building command")
	cmd := exec.CommandContext(ctx, bin, command.Args...)
	cmd.Stdin = stdinReader(command.Stdin)

	if cmd.Err != nil {
		log.ErrorContext(ctx, "failed building command", "err", cmd.Err)
		return nil, fmt.Errorf("failed building command: %w", cmd.Err)
	}

	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout, cmd.Stderr = stdout, stderr

	log.DebugContext(ctx, "running command", "cmd", cmd.String())
	if err = cmd.Run(); err != nil {
		log.WarnContext(ctx, "command failed", "err", err)
	}

	if cmd.ProcessState == nil {
		return nil, errors.New("failed to start command")
	}

	exitCode := cmd.ProcessState.ExitCode()
	log.InfoContext(ctx, "finished running command",
		"cmd", cmd.String(),
		"exit_code", exitCode,
	)

	return &pb.Result{
		ExitCode: int32(exitCode),
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
	}, nil
}
