package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	cmd "github.com/unmango/pulumi-baremetal/provider/pkg/command"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal"
)

type service struct {
	pb.UnimplementedCommandServiceServer
	internal.State
}

func NewServer(state internal.State) pb.CommandServiceServer {
	log := state.Log.With("service", "command")
	return &service{State: state.WithLogger(log)}
}

func (s *service) Create(ctx context.Context, req *pb.CreateRequest) (res *pb.CreateResponse, err error) {
	log := s.Log.With("op", "create")
	if req.Command == nil {
		log.ErrorContext(ctx, "no command found in request")
		return nil, fmt.Errorf("no command found in request")
	}

	args := req.Command.Args
	bin, err := cmd.BinValue(req.Command.Bin)
	if err != nil {
		log.ErrorContext(ctx, "unable to map bin", "err", err)
		return nil, fmt.Errorf("mapping bin: %w", err)
	}

	log = log.With("bin", bin, "args", args)
	log.DebugContext(ctx, "building command")
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stdin = stdinReader(req.Command.Stdin)

	if cmd.Err != nil {
		log.ErrorContext(ctx, "failed building command", "err", cmd.Err)
		return nil, fmt.Errorf("failed building command: %w", cmd.Err)
	}

	createdFiles := make([]string, len(req.ExpectCreated))
	movedFiles := make(map[string]string, len(req.ExpectMoved))
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout, cmd.Stderr = stdout, stderr

	log.DebugContext(ctx, "running command", "cmd", cmd.String())
	if err = cmd.Run(); err != nil {
		log.WarnContext(ctx, "command failed", "err", err)
	} else {
		for i, file := range req.ExpectCreated {
			if _, err := os.Stat(file); err != nil {
				log.ErrorContext(ctx, "expected file did not exist", "file", file, "err", err)
			} else {
				createdFiles[i] = file
			}
		}
		for src, dest := range req.ExpectMoved {
			srcExists, destExists := false, true
			if _, err = os.Stat(src); !errors.Is(err, os.ErrNotExist) {
				log.ErrorContext(ctx, "expected file not to exist", "file", src, "err", err)
				srcExists = true
			}
			if _, err = os.Stat(dest); err != nil {
				log.ErrorContext(ctx, "expected file did not exist", "file", dest, "err", err)
				destExists = false
			}
			if !srcExists && destExists {
				movedFiles[src] = dest
			}
		}
	}

	if cmd.ProcessState == nil {
		return nil, errors.New("failed to start command")
	}

	exitCode := cmd.ProcessState.ExitCode()
	log.InfoContext(ctx, "finished executing command",
		"cmd", cmd.String(),
		"exit_code", exitCode,
		"created", createdFiles,
		"moved", movedFiles,
	)

	return &pb.CreateResponse{
		CreatedFiles: createdFiles,
		MovedFiles:   movedFiles,
		Result: &pb.Result{
			ExitCode: int32(exitCode),
			Stdout:   stdout.String(),
			Stderr:   stderr.String(),
		},
	}, nil
}

func (s *service) Update(ctx context.Context, req *pb.UpdateRequest) (res *pb.UpdateResponse, err error) {
	log := s.Log.With("op", "update", "prev", req.Previous)

	create, err := s.Create(ctx, &pb.CreateRequest{
		Command:       req.Command,
		ExpectCreated: req.ExpectCreated,
		ExpectMoved:   req.ExpectMoved,
	})
	if err != nil {
		log.ErrorContext(ctx, "failed performing create", "err", err)
		return nil, fmt.Errorf("failed creating: %w", err)
	}

	return &pb.UpdateResponse{
		Result:       create.Result,
		CreatedFiles: create.CreatedFiles,
		MovedFiles:   create.MovedFiles,
	}, nil
}

func (s *service) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log := s.Log.With("op", "delete", "prev", req.Previous)

	if req.Command != nil {
		log.Info("executing custom delete")
		create, err := s.Create(ctx, &pb.CreateRequest{
			Command:       req.Command,
			ExpectCreated: []string{},
			ExpectMoved:   map[string]string{},
		})
		if err != nil {
			log.ErrorContext(ctx, "failed performing create", "err", err)
			return nil, fmt.Errorf("failed creating: %w", err)
		}

		return &pb.DeleteResponse{Commands: []*pb.Operation{{
			Result:       create.Result,
			Command:      req.Command,
			CreatedFiles: create.CreatedFiles,
			MovedFiles:   create.MovedFiles,
		}}}, nil
	}

	commands := []*pb.Operation{}
	toDelete := req.Previous.CreatedFiles
	toMove := req.Previous.MovedFiles
	if len(toDelete) == 0 && len(toMove) == 0 {
		log.InfoContext(ctx, "nothing to do")
		return &pb.DeleteResponse{Commands: commands}, nil
	}

	if len(toDelete) > 0 {
		// I think `rm` handles this these days but you can never be too sure
		if slices.Contains(toDelete, "/") {
			log.ErrorContext(ctx, "refusing to remove '/'", "remark", "nice try hackers")
			return nil, errors.New("attempted to remove root")
		}

		log.DebugContext(ctx, "building command")
		cmd := exec.CommandContext(ctx, "rm", prepend("-f", toDelete)...)
		stderr, stdout := &bytes.Buffer{}, &bytes.Buffer{}
		cmd.Stdout, cmd.Stderr = stdout, stderr

		log.DebugContext(ctx, "running command", "cmd", cmd.String())
		if err := cmd.Run(); err != nil {
			log.ErrorContext(ctx, "command failed", "err", err)
		}

		log.InfoContext(ctx, "finished executing command", "cmd", cmd.String())
		commands = append(commands, &pb.Operation{
			Command: &pb.Command{Bin: pb.Bin_BIN_RM, Args: cmd.Args},
			Result: &pb.Result{
				ExitCode: int32(cmd.ProcessState.ExitCode()),
				Stdout:   stdout.String(),
				Stderr:   stderr.String(),
			},
		})
	}

	for src, dest := range req.Previous.MovedFiles {
		log.DebugContext(ctx, "building command")
		cmd := exec.CommandContext(ctx, "mv", dest, src)
		stderr, stdout := &bytes.Buffer{}, &bytes.Buffer{}
		cmd.Stdout, cmd.Stderr = stdout, stderr

		log.DebugContext(ctx, "running command", "cmd", cmd.String())
		if err := cmd.Run(); err != nil {
			log.ErrorContext(ctx, "command failed",
				"err", err,
				"stdout", stdout.String(),
				"stderr", stderr.String(),
			)
		}

		log.InfoContext(ctx, "finished executing command", "cmd", cmd.String())
		commands = append(commands, &pb.Operation{
			Command: &pb.Command{Bin: pb.Bin_BIN_MV, Args: cmd.Args},
			Result: &pb.Result{
				ExitCode: int32(cmd.ProcessState.ExitCode()),
				Stdout:   stdout.String(),
				Stderr:   stderr.String(),
			},
			CreatedFiles: []string{},
			MovedFiles:   map[string]string{dest: src},
		})
	}

	return &pb.DeleteResponse{Commands: commands}, nil
}

func stdinReader(stdin *string) io.Reader {
	if stdin == nil {
		return nil
	}

	return strings.NewReader(*stdin)
}

func prepend[T any](x T, xs []T) []T {
	return slices.Insert(xs, 0, x)
}
