package services

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

const (
	SshUserName string = "pulumi-baremetal"
	SshPassword string = "Password!123"
	SshPort     int    = 2222
)

//go:embed .versions/openssh-server
var version string

type Sshd struct{ Host }

func NewSshd(ctx context.Context, logger io.Writer) (*Sshd, error) {
	req := tc.GenericContainerRequest{
		Logger: util.NewLogger(logger),
		ContainerRequest: tc.ContainerRequest{
			Image: fmt.Sprintf(
				"lscr.io/linuxserver/openssh-server:%s",
				strings.TrimSpace(version),
			),
			AlwaysPullImage: true,
			ExposedPorts:    []string{fmt.Sprint(SshPort)},
			WaitingFor:      wait.ForExposedPort(),
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

	return &Sshd{Host{req, nil}}, nil
}

func (s *Sshd) ConnectionProps(ctx context.Context) (resource.PropertyValue, error) {
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
