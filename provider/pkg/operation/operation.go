package operation

import (
	"fmt"
	"strings"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/command"
)

func Display(o *pb.Operation) string {
	cmd := command.Display(o.Command)

	return fmt.Sprintf(
		"exitCode: %d, command: %s",
		o.Result.ExitCode, cmd,
	)
}

func DisplayAll(ops []*pb.Operation) string {
	b := strings.Builder{}
	b.WriteRune('[')

	for _, o := range ops {
		_, _ = b.WriteString(Display(o) + ", ")
	}

	b.WriteRune(']')
	return b.String()
}

func DisplayCommand(c *pb.Command, r *pb.Result) string {
	return Display(&pb.Operation{Command: c, Result: r})
}

func FromCreate(command *pb.Command, res *pb.CreateResponse) *pb.Operation {
	return &pb.Operation{
		Command:      command,
		Result:       res.Result,
		CreatedFiles: res.CreatedFiles,
		MovedFiles:   res.MovedFiles,
	}
}

func FromUpdate(command *pb.Command, res *pb.UpdateResponse) *pb.Operation {
	return &pb.Operation{
		Command:      command,
		Result:       res.Result,
		CreatedFiles: res.CreatedFiles,
		MovedFiles:   res.MovedFiles,
	}
}
