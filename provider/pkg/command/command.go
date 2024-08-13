package command

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

var (
	BinToValue = map[pb.Bin]string{
		pb.Bin_BIN_CHMOD:  "chmod",
		pb.Bin_BIN_MKDIR:  "mkdir",
		pb.Bin_BIN_MKTEMP: "mktemp",
		pb.Bin_BIN_MV:     "mv",
		pb.Bin_BIN_RM:     "rm",
		pb.Bin_BIN_TAR:    "tar",
		pb.Bin_BIN_TEE:    "tee",
		pb.Bin_BIN_WGET:   "wget",
	}
	ValueToBin = make(map[string]pb.Bin, len(BinToValue))
)

func init() {
	for k, v := range BinToValue {
		ValueToBin[v] = k
	}
}

func Display(c *pb.Command) string {
	bin, ok := BinToValue[c.Bin]
	if !ok {
		bin = "<nil>"
	}

	parts := slices.Concat([]string{bin}, c.Args)
	return strings.Join(parts, " ")
}

func Parse(args []string) (*pb.Command, error) {
	if len(args) == 0 {
		return nil, errors.New("must have at least one argument to form a command")
	}

	bin, err := ParseBin(args[0])
	if err != nil {
		return nil, fmt.Errorf("failed parsing bin for command: %s", err)
	}

	return &pb.Command{
		Bin:   bin,
		Args:  args[1:],
		Stdin: nil, // TODO
	}, nil
}

func ParseBin(bin string) (pb.Bin, error) {
	if v, ok := ValueToBin[bin]; ok {
		return v, nil
	}

	return pb.Bin_BIN_UNSPECIFIED, fmt.Errorf("unsupported command: %s", bin)
}
