package provider

import (
	"fmt"

	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type BootstrapArgs struct {
	Connection remote.Connection `pulumi:"connection" provider:"type=command@1.0.1:remote:Connection"`
	Version    string            `pulumi:"version"`
}

type Bootstrap struct {
	pulumi.ResourceState

	Mktemp   *remote.Command `pulumi:"mktemp" provider:"type=command@1.0.1:remote:Command"`
	Download *remote.Command `pulumi:"download" provider:"type=command@1.0.1:remote:Command"`
	Url      string          `pulumi:"url"`
}

var _ infer.ComponentResource[BootstrapArgs, *Bootstrap] = &Bootstrap{}

// Construct implements infer.ComponentResource.
func (*Bootstrap) Construct(ctx *pulumi.Context, name string, typ string, inputs BootstrapArgs, opts pulumi.ResourceOption) (*Bootstrap, error) {
	state := &Bootstrap{}

	conn := mapConnection(inputs.Connection)
	version := inputs.Version
	fileName := fmt.Sprintf("provisioner-v%s-linux-arm64.tar.gz", version)
	url := fmt.Sprintf("https://github.com/unmango/pulumi-baremetal/releases/download/v%s/%s", version, fileName)

	mktemp, err := remote.NewCommand(ctx, "mktemp", &remote.CommandArgs{
		Connection: conn,
		Create:     pulumi.StringPtr("mktemp --directory"),
	}, pulumi.Parent(state))
	if err != nil {
		return nil, err
	}

	tempDir := pulumix.Apply(mktemp.Stdout, func(stdout string) string {
		return stdout
	})

	createDownload := pulumix.Apply(tempDir, func(dir string) string {
		return fmt.Sprintf("wget --directory-prefix %s %s", dir, url)
	})

	download, err := remote.NewCommand(ctx, "download", &remote.CommandArgs{
		Connection: conn,
		Create:     pulumix.Cast[pulumi.StringOutput](createDownload),
	}, pulumi.Parent(state), pulumi.DependsOn([]pulumi.Resource{mktemp}))
	if err != nil {
		return nil, err
	}

	state.Download = download

	return state, nil
}

func mapConnection(conn remote.Connection) remote.ConnectionArgs {
	return remote.ConnectionArgs{
		AgentSocketPath:    pulumi.StringPtrFromPtr(conn.AgentSocketPath),
		DialErrorLimit:     pulumi.IntPtrFromPtr(conn.DialErrorLimit),
		Host:               pulumi.String(conn.Host),
		Password:           pulumi.StringPtrFromPtr(conn.Password),
		PerDialTimeout:     pulumi.IntPtrFromPtr(conn.PerDialTimeout),
		Port:               pulumi.Float64PtrFromPtr(conn.Port),
		PrivateKey:         pulumi.StringPtrFromPtr(conn.PrivateKey),
		PrivateKeyPassword: pulumi.StringPtrFromPtr(conn.Proxy.PrivateKeyPassword),
	}
}
