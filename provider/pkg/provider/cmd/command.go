package cmd

import (
	"fmt"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

func parseCommand(args []string) (*pb.Command, error) {
	var bin pb.Bin
	switch args[0] {
	case "chmod":
		bin = pb.Bin_BIN_CHMOD
	case "mkdir":
		bin = pb.Bin_BIN_MKDIR
	case "mktemp":
		bin = pb.Bin_BIN_MKTEMP
	case "mv":
		bin = pb.Bin_BIN_MV
	case "rm":
		bin = pb.Bin_BIN_RM
	case "tar":
		bin = pb.Bin_BIN_TAR
	case "tee":
		bin = pb.Bin_BIN_TEE
	case "wget":
		bin = pb.Bin_BIN_WGET
	default:
		return nil, fmt.Errorf("unsupported command: %s", args[0])
	}

	return &pb.Command{
		Bin:   bin,
		Args:  args[1:],
		Stdin: nil, // TODO
	}, nil
}
