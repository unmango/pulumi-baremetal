package provider

import (
	"fmt"
	"path"
	"strings"

	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

const binName string = "provisioner"

type BootstrapArgs struct {
	Connection *remote.Connection `pulumi:"connection,optional" provider:"type=command@1.0.1:remote:Connection"`
	Directory  string             `pulumi:"directory"`
	Version    string             `pulumi:"version"`
}

var _ = (infer.Annotated)((*BootstrapArgs)(nil))

// Annotate implements infer.Annotated.
func (b *BootstrapArgs) Annotate(a infer.Annotator) {
	a.Describe(&b.Directory, "The directory to store the provisioner binary.")
	a.SetDefault(&b.Directory, "/usr/local/bin")
	a.Describe(&b.Version, "The version of the provisioner to bootstrap")
}

type Bootstrap struct {
	pulumi.ResourceState
	conn remote.ConnectionArgs

	ArchiveName pulumix.Output[string] `pulumi:"archiveName"`
	BinPath     pulumix.Output[string] `pulumi:"binPath"`
	Download    *remote.Command        `pulumi:"download" provider:"type=command@1.0.1:remote:Command"`
	Extract     *remote.Command        `pulumi:"extract" provider:"type=command@1.0.1:remote:Command"`
	FileName    string                 `pulumi:"fileName"`
	Mkdir       *remote.Command        `pulumi:"mkdir" provider:"type=command@1.0.1:remote:Command"`
	Mktemp      *remote.Command        `pulumi:"mktemp" provider:"type=command@1.0.1:remote:Command"`
	Mv          *remote.Command        `pulumi:"mv" provider:"type=command@1.0.1:remote:Command"`
	TempDir     pulumix.Output[string] `pulumi:"tempDir"`
	Url         string                 `pulumi:"url"`
}

var _ = (infer.Annotated)((*Bootstrap)(nil))

// Annotate implements infer.Annotated.
func (b *Bootstrap) Annotate(a infer.Annotator) {
	a.Describe(&b.ArchiveName, "Name part of the provisioner release archive file.")
	a.Describe(&b.BinPath, "Provisioner binary path on the remote system.")
	a.Describe(&b.Download, "Binary download command.")
	a.Describe(&b.FileName, "Name part of the provisioner binary file.")
	a.Describe(&b.Mkdir, "Destination directory mkdir command.")
	a.Describe(&b.Mktemp, "Temp download directory mktemp command.")
	a.Describe(&b.Mv, "Command to move the binary from the temp directory to the destination.")
	a.Describe(&b.TempDir, "Temp directory path output by the mktemp command.")
	a.Describe(&b.Url, "Url of the provisioner release archive.")
}

var _ = (infer.ComponentResource[BootstrapArgs, *Bootstrap])((*Bootstrap)(nil))

// Construct implements infer.ComponentResource.
func (*Bootstrap) Construct(ctx *pulumi.Context, name string, typ string, inputs BootstrapArgs, opts pulumi.ResourceOption) (*Bootstrap, error) {
	state := &Bootstrap{conn: mapConnection(inputs.Connection)}

	version := inputs.Version
	archiveName := fmt.Sprintf("provisioner-v%s-linux-arm64.tar.gz", version)
	url := fmt.Sprintf("https://github.com/unmango/pulumi-baremetal/releases/download/v%s/%s", version, archiveName)

	ctx.Log.Debug("Mktemp", nil)
	if err := state.mktemp(ctx, pulumi.Any(version)); err != nil {
		return nil, err
	}

	ctx.Log.Debug("Download", nil)
	if err := state.download(ctx, url); err != nil {
		return nil, err
	}

	ctx.Log.Debug("Mkdir", nil)
	if err := state.mkdir(ctx, inputs.Directory); err != nil {
		return nil, err
	}

	ctx.Log.Debug("Extract", nil)
	if err := state.extract(ctx, archiveName); err != nil {
		return nil, err
	}

	ctx.Log.Debug("Mv", nil)
	if err := state.mv(ctx, inputs.Directory, pulumi.Any(version)); err != nil {
		return nil, err
	}

	return state, ctx.RegisterResourceOutputs(state,
		pulumi.ToMap(map[string]interface{}{
			"archiveName": state.ArchiveName,
			"binPath":     state.BinPath,
			"download":    state.Download,
			"extract":     state.Extract,
			"fileName":    state.FileName,
			"mkdir":       state.Mkdir,
			"mktemp":      state.Mktemp,
			"mv":          state.Mv,
			"tempDir":     state.TempDir,
			"url":         state.Url,
		}),
	)
}

func (s *Bootstrap) mktemp(ctx *pulumi.Context, triggers ...pulumix.Input[any]) error {
	cmd, err := remote.NewCommand(ctx, "mktemp", &remote.CommandArgs{
		Connection: s.conn,
		Create:     pulumi.StringPtr("mktemp --directory"),
		Triggers:   pulumix.Cast[pulumi.ArrayOutput](pulumix.All(triggers...)),
	}, pulumi.Parent(s))
	if err != nil {
		return err
	}

	tempDir, err := pulumix.ConvertTyped[string](cmd.Stdout)
	if err != nil {
		return err
	}

	s.TempDir = tempDir
	s.Mktemp = cmd

	return nil
}

func (s *Bootstrap) download(ctx *pulumi.Context, url string) error {
	create := pulumix.Apply(s.TempDir, func(dir string) string {
		return fmt.Sprintf("wget --directory-prefix %s %s", dir, url)
	})

	cmd, err := remote.NewCommand(ctx, "download", &remote.CommandArgs{
		Connection: s.conn,
		Create:     pulumix.Cast[pulumi.StringOutput](create),
	}, pulumi.Parent(s), pulumi.DependsOn([]pulumi.Resource{s.Mktemp}))

	s.Download = cmd
	s.Url = url

	return err
}

func (s *Bootstrap) mkdir(ctx *pulumi.Context, dir string) error {
	cmd, err := remote.NewCommand(ctx, "mkdir", &remote.CommandArgs{
		Connection: s.conn,
		Create:     pulumi.StringPtr(fmt.Sprintf("mkdir --parents %s", dir)),
	}, pulumi.Parent(s))

	s.Mkdir = cmd
	return err
}

func (s *Bootstrap) extract(ctx *pulumi.Context, archiveName string) error {
	create := pulumix.Apply(s.TempDir, func(dir string) string {
		return strings.Join([]string{
			"tar", "--extract", "--gzip",
			"--strip-components", "1",
			"--file", path.Join(dir, archiveName),
			"--directory", dir,
			"--verbose",
		}, " ")
	})

	cmd, err := remote.NewCommand(ctx, "extract",
		&remote.CommandArgs{
			Connection: s.conn,
			Create:     pulumix.Cast[pulumi.StringOutput](create),
		},
		pulumi.Parent(s),
		pulumi.DependsOn([]pulumi.Resource{
			s.Mktemp,
			s.Download,
		}),
	)

	s.ArchiveName = pulumix.Val(archiveName)
	s.Extract = cmd

	return err
}

func (s *Bootstrap) mv(ctx *pulumi.Context, dir string, triggers ...pulumix.Input[any]) error {
	downloadPath := pulumix.Apply(s.TempDir, func(dir string) string {
		return path.Join(dir, binName)
	})

	binPath := path.Join(dir, binName)
	create := pulumix.Apply(downloadPath, func(dp string) string {
		return fmt.Sprintf("mv %s %s", dp, binPath)
	})

	cmd, err := remote.NewCommand(ctx, "mv",
		&remote.CommandArgs{
			Connection: s.conn,
			Create:     pulumix.Cast[pulumi.StringOutput](create),
			Delete:     pulumi.StringPtr(fmt.Sprintf("rm -f %s", binPath)),
			Triggers:   pulumix.Cast[pulumi.ArrayOutput](pulumix.All(triggers...)),
		},
		pulumi.Parent(s),
		pulumi.DependsOn([]pulumi.Resource{
			s.Extract,
			s.Mkdir,
		}),
	)

	s.FileName = binName
	s.BinPath = pulumix.Val(binPath)
	s.Mv = cmd

	return err
}

func mapConnection(conn *remote.Connection) remote.ConnectionArgs {
	if conn == nil {
		panic("TODO")
	}

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
