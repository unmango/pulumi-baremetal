package operation

import (
	"fmt"
	"strings"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/command"
)

func Display(o *pb.Operation) string {
	cmd := command.Display(o.Command)
	if o.Result == nil {
		return cmd
	}

	return fmt.Sprintf(
		"exitCode: %d, command: %s",
		o.Result.ExitCode, cmd,
	)
}

func DisplayAll(ops []*pb.Operation) string {
	b := strings.Builder{}
	b.WriteRune('[')

	d := make([]string, len(ops))
	for i, o := range ops {
		d[i] = Display(o)
	}

	b.WriteString(strings.Join(d, ", "))
	b.WriteRune(']')
	return b.String()
}

func DisplayCommand(c *pb.Command, r *pb.Result) string {
	return Display(&pb.Operation{Command: c, Result: r})
}

func FromCreate(command *pb.Command, res *pb.CreateResponse) *pb.Operation {
	if res == nil {
		return &pb.Operation{
			Command:      command,
			CreatedFiles: []string{},
			MovedFiles:   map[string]string{},
		}
	}

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
