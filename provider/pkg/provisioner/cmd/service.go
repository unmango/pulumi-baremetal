package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"slices"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
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
		log.Error("no command found in request")
		return nil, fmt.Errorf("no command found in request")
	}

	args := req.Command.Args
	bin, err := binPath(req.Command.Bin)
	if err != nil {
		log.Error("unable to map bin", "err", err)
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

	createdFiles := make([]string, len(req.ExpectFiles))
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout, cmd.Stderr = stdout, stderr

	log.DebugContext(ctx, "running command", "cmd", cmd.String())
	if err = cmd.Run(); err != nil {
		log.WarnContext(ctx, "command failed", "err", err)
	} else {
		for i, file := range req.ExpectFiles {
			if _, err := os.Stat(file); err != nil {
				log.ErrorContext(ctx, "expected file did not exist", "file", file, "err", err)
			} else {
				createdFiles[i] = file
			}
		}
	}

	if cmd.ProcessState == nil {
		return nil, errors.New("failed to start command")
	}

	exitCode := cmd.ProcessState.ExitCode()
	log.InfoContext(ctx, "finished executing command", "cmd", cmd.String(), "created", createdFiles)
	return &pb.CreateResponse{
		Files: createdFiles,
		Result: &pb.Result{
			ExitCode: int32(exitCode),
			Stdout:   stdout.String(),
			Stderr:   stderr.String(),
		},
	}, nil
}

func (s *service) Update(ctx context.Context, req *pb.UpdateRequest) (res *pb.UpdateResponse, err error) {
	log := s.Log.With("op", "update", "create", req.Create)

	var delete *pb.DeleteResponse
	toDelete := req.Create.Files
	if len(toDelete) > 0 {
		delete, err = s.Delete(ctx, &pb.DeleteRequest{Create: req.Create})
	} else {
		log.InfoContext(ctx, "nothing to delete")
	}
	if err != nil {
		log.ErrorContext(ctx, "failed performing delete", "err", err)
		return nil, fmt.Errorf("failed deleting: %w", err)
	}

	create, err := s.Create(ctx, &pb.CreateRequest{
		Command:     req.Command,
		ExpectFiles: req.ExpectFiles,
	})
	if err != nil {
		log.ErrorContext(ctx, "failed performing create", "err", err)
		return nil, fmt.Errorf("failed creating: %w", err)
	}

	return &pb.UpdateResponse{
		Delete: delete.Op,
		Result: create.Result,
		Files:  create.Files,
	}, nil
}

func (s *service) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log := s.Log.With("op", "delete", "create", req.Create)
	bin := pb.Bin_BIN_RM

	toDelete := req.Create.Files
	if len(toDelete) == 0 {
		log.InfoContext(ctx, "nothing to do")
		return &pb.DeleteResponse{}, nil
	}

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
		return &pb.DeleteResponse{}, nil
	}

	log.DebugContext(ctx, "describing operation")
	op := &pb.Operation{
		Command: &pb.Command{Bin: bin, Args: cmd.Args},
		Result: &pb.Result{
			ExitCode: int32(cmd.ProcessState.ExitCode()),
			Stdout:   stdout.String(),
			Stderr:   stderr.String(),
		},
	}

	log.InfoContext(ctx, "finished executing command", "cmd", cmd.String())
	return &pb.DeleteResponse{Op: op}, nil
}

func binPath(b pb.Bin) (string, error) {
	switch b {
	case pb.Bin_BIN_MV:
		return "mv", nil
	case pb.Bin_BIN_RM:
		return "rm", nil
	case pb.Bin_BIN_TEE:
		return "tee", nil
	case pb.Bin_BIN_WGET:
		return "wget", nil
	}

	return "", fmt.Errorf("unrecognized bin: %s", b)
}
