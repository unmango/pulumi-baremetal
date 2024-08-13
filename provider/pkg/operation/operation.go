package operation

import (
	"fmt"

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
