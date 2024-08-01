package cmd

import (
	"io"
	"slices"
	"strings"
)

func stdinReader(stdin *string) io.Reader {
	if stdin == nil {
		return nil
	}

	return strings.NewReader(*stdin)
}

func prepend[T any](x T, xs []T) []T {
	return slices.Insert(xs, 0, x)
}

// func run(cmd *exec.Cmd, log *slog.Logger) (*pb.CommandResponse, error) {
// 	stdout := &bytes.Buffer{}
// 	stderr := &bytes.Buffer{}

// 	cmd.Stdout = stdout
// 	cmd.Stderr = stderr

// 	log.Info("executing command")
// 	if err := cmd.Run(); err != nil {
// 		log.Error("command failed", "err", err)
// 	}

// 	log.Debug("command succeeded")
// 	return &pb.CommandResponse{
// 		Stdout: stdout.String(),
// 		Stderr: stderr.String(),
// 	}, nil
// }
