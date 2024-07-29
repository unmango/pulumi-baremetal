package util

import (
	"context"
	"fmt"

	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	SshUserName string = "pulumi-baremetal"
	SshPassword string = "Password!123"
)

const version string = "version-9.7_p1-r4" // TODO: Go embed?

type SshServer interface {
	TestHost
}

type server struct{ host }

var _ = (SshServer)((*server)(nil))

func NewSshServer(ctx context.Context) (SshServer, error) {
	ctr, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			Image:        fmt.Sprintf("lscr.io/linuxserver/openssh-server:%s", version),
			ExposedPorts: []string{"22"},
			WaitingFor:   wait.ForExposedPort(),
			Env: map[string]string{
				"PUID":            "1000",
				"PGID":            "1000",
				"TZ":              "America/Chicago",
				"SUDO_ACCESS":     "true",
				"PASSWORD_ACCESS": "true",
				"USER_NAME":       SshUserName,
				"USER_PASSWORD":   SshPassword,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	host := host{ctr: ctr}
	return &server{host: host}, nil
}
