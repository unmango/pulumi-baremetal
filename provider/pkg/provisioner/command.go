package provisioner

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider"
)

type commandService struct {
	pb.UnimplementedCommandServiceServer
}

func NewCommandService() pb.CommandServiceServer {
	return &commandService{}
}

// Command implements baremetalv1alpha1.CommandServiceServer.
func (c *commandService) Command(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	reader := bytes.NewReader(req.PulumiRaw)
	dec := gob.NewDecoder(reader)

	var input provider.TeeArgs
	if err := dec.Decode(&input); err != nil {
		return nil, err
	}

	stderr, err := json.Marshal(input.Create.Files)
	if err != nil {
		return nil, err
	}

	return &pb.CommandResponse{
		Stdout: input.Stdin,
		Stderr: string(stderr),
	}, nil
}
