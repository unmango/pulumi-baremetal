package command

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"slices"
	"strings"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/command"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal/opts"
	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedCommandServiceServer
	Log       *slog.Logger
	Whitelist []string
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
	return func(s *service) error {
		s.Log = logger
		return nil
	}
}

func WithWhitelist(list []string) opt {
	return func(s *service) error {
		s.Whitelist = list
		return nil
	}
}

func (s *service) Register(server *grpc.Server) {
	pb.RegisterCommandServiceServer(server, s)
}

func (s *service) Exec(ctx context.Context, req *pb.ExecRequest) (*pb.ExecResponse, error) {
	log := s.Log.With("op", "exec")

	if len(req.Args) == 0 {
		log.ErrorContext(ctx, "no command provided")
		return nil, errors.New("no command provided")
	}

	bin := req.Args[0]
	_, err := command.ParseBin(bin)
	if err != nil && !slices.Contains(s.Whitelist, bin) {
		log.WarnContext(ctx, "refusing to execute command", "bin", bin, "err", err)
		return nil, fmt.Errorf("refusing to execute %s %#v", bin, s.Whitelist)
	}

	args := req.Args[1:]
	log = log.With("bin", bin, "args", args)
	log.DebugContext(ctx, "building command")
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stdin = stdinReader(req.Stdin)

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

	return &pb.ExecResponse{Result: &pb.Result{
		ExitCode: int32(exitCode),
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
	}}, nil
}

func (s *service) Create(ctx context.Context, req *pb.CreateRequest) (res *pb.CreateResponse, err error) {
	log := s.Log.With("op", "create")
	createdFiles := make([]string, len(req.ExpectCreated))
	movedFiles := make(map[string]string, len(req.ExpectMoved))

	result, err := execute(ctx, req.Command, log)
	if err != nil || result.ExitCode > 0 {
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

	log.InfoContext(ctx, "processed command result",
		"result", result,
		"created", createdFiles,
		"moved", movedFiles,
	)

	return &pb.CreateResponse{
		Result:       result,
		CreatedFiles: createdFiles,
		MovedFiles:   movedFiles,
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
			Command: &pb.Command{
				Bin:  pb.Bin_BIN_RM,
				Args: cmd.Args[1:],
			},
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
			Command: &pb.Command{
				Bin:  pb.Bin_BIN_MV,
				Args: cmd.Args[1:],
			},
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
