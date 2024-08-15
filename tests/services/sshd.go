package services

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	SshUserName string = "pulumi-baremetal"
	SshPassword string = "Password!123"
	SshPort     int    = 2222
)

const version string = "version-9.7_p1-r4" // TODO: Go embed?

type Sshd interface {
	TestHost

	ConnectionProps(context.Context) (resource.PropertyValue, error)
}

type server struct{ host }

var _ = (Sshd)((*server)(nil))

func NewSshd(ctx context.Context) (Sshd, error) {
	req := tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			Image:        fmt.Sprintf("lscr.io/linuxserver/openssh-server:%s", version),
			ExposedPorts: []string{fmt.Sprint(SshPort)},
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
	}

	return &server{host{req, nil}}, nil
}

func (s *server) ConnectionProps(ctx context.Context) (resource.PropertyValue, error) {
	ip, err := s.Ip(ctx)
	if err != nil {
		return resource.PropertyValue{}, err
	}

	props := resource.NewObjectProperty(resource.PropertyMap{
		"host":     resource.NewStringProperty(ip),
		"port":     resource.NewPropertyValue(SshPort),
		"user":     resource.NewStringProperty(SshUserName),
		"password": resource.NewStringProperty(SshPassword),
	})

	return props, nil
}
