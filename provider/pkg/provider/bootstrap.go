package provider

import (
	"fmt"
	"path"
	"strings"

	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi-go-provider/infer"
	p "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	px "github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-baremetal/provider/pkg/fx"
)

const binName string = "provisioner"

type Bootstrap struct{}

type BootstrapArgs struct {
	Connection *remote.Connection `pulumi:"connection,optional" provider:"type=command@1.0.1:remote:Connection"`
	Directory  p.StringInput      `pulumi:"directory,optional"`
	Version    p.StringInput      `pulumi:"version"`
}

var _ = (infer.Annotated)((*BootstrapArgs)(nil))

// Annotate implements infer.Annotated.
func (b *BootstrapArgs) Annotate(a infer.Annotator) {
	a.Describe(&b.Directory, "The directory to store the provisioner binary.")
	a.SetDefault(&b.Directory, "/usr/local/bin")
	a.Describe(&b.Version, "The version of the provisioner to bootstrap")
}

type BootstrapState struct {
	p.ResourceState
	conn *remote.ConnectionArgs

	ArchiveName px.Output[string] `pulumi:"archiveName"`
	BinPath     px.Output[string] `pulumi:"binPath"`
	Download    *remote.Command   `pulumi:"download" provider:"type=command@1.0.1:remote:Command"`
	Extract     *remote.Command   `pulumi:"extract" provider:"type=command@1.0.1:remote:Command"`
	FileName    string            `pulumi:"fileName"`
	Mkdir       *remote.Command   `pulumi:"mkdir" provider:"type=command@1.0.1:remote:Command"`
	Mktemp      *remote.Command   `pulumi:"mktemp" provider:"type=command@1.0.1:remote:Command"`
	Mv          *remote.Command   `pulumi:"mv" provider:"type=command@1.0.1:remote:Command"`
	TempDir     px.Output[string] `pulumi:"tempDir"`
	Url         string            `pulumi:"url"`
}

var _ = (infer.Annotated)((*BootstrapState)(nil))

// Annotate implements infer.Annotated.
func (b *BootstrapState) Annotate(a infer.Annotator) {
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

var _ = (infer.ComponentResource[BootstrapArgs, *BootstrapState])((*Bootstrap)(nil))

// Construct implements infer.ComponentResource.
func (*Bootstrap) Construct(ctx *p.Context, name string, typ string, inputs BootstrapArgs, opts p.ResourceOption) (*BootstrapState, error) {
	state := &BootstrapState{conn: mapConnection(inputs.Connection)}
	err := ctx.RegisterComponentResource("baremetal", name, state, opts)
	if err != nil {
		return nil, err
	}

	version := inputs.Version
	archiveName := fmt.Sprintf("provisioner-v%s-linux-arm64.tar.gz", version)
	url := fmt.Sprintf("https://github.com/unmango/pulumi-baremetal/releases/download/v%s/%s", version, archiveName)

	ctx.Log.Debug("Mktemp", nil)
	if err := state.mktemp(ctx, p.Any(version)); err != nil {
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
	if err := state.mv(ctx, inputs.Directory, p.Any(version)); err != nil {
		return nil, err
	}

	return state, nil
}

func (s *BootstrapState) mktemp(ctx *p.Context, triggers ...px.Input[any]) error {
	cmd, err := remote.NewCommand(ctx, "mktemp", &remote.CommandArgs{
		Connection: s.conn,
		Create:     p.StringPtr("mktemp --directory"),
		Triggers:   px.Cast[p.ArrayOutput](px.All(triggers...)),
	}, p.Parent(s))
	if err != nil {
		return err
	}

	tempDir, err := px.ConvertTyped[string](cmd.Stdout)
	if err != nil {
		return err
	}

	s.TempDir = tempDir
	s.Mktemp = cmd

	return nil
}

func (s *BootstrapState) download(ctx *p.Context, url string) error {
	create := fx.Sprintf("wget --directory-prefix %s %s",
		s.TempDir.AsAny(),
		px.Val(url).AsAny())

	cmd, err := remote.NewCommand(ctx, "download", &remote.CommandArgs{
		Connection: s.conn,
		Create:     px.Cast[p.StringOutput](create),
	}, p.Parent(s), p.DependsOn([]p.Resource{s.Mktemp}))

	s.Download = cmd
	s.Url = url

	return err
}

func (s *BootstrapState) mkdir(ctx *p.Context, dir p.StringInput) error {
	create := px.Apply(dir.ToStringOutput(), func(dir string) string {
		return fmt.Sprintf("mkdir --parents %s", dir)
	})

	cmd, err := remote.NewCommand(ctx, "mkdir", &remote.CommandArgs{
		Connection: s.conn,
		Create:     px.Cast[p.StringOutput](create),
	}, p.Parent(s))

	s.Mkdir = cmd
	return err
}

func (s *BootstrapState) extract(ctx *p.Context, archiveName string) error {
	create := px.Apply(s.TempDir, func(dir string) string {
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
			Create:     px.Cast[p.StringOutput](create),
		},
		p.Parent(s),
		p.DependsOn([]p.Resource{
			s.Mktemp,
			s.Download,
		}),
	)

	s.ArchiveName = px.Val(archiveName)
	s.Extract = cmd

	return err
}

func (s *BootstrapState) mv(ctx *p.Context, dir p.StringInput, triggers ...px.Input[any]) error {
	downloadPath := px.Apply(s.TempDir, func(tmp string) string {
		return path.Join(tmp, binName)
	})

	binPath := px.Apply(dir.ToStringOutput(), func(dir string) string {
		return path.Join(dir, binName)
	})

	create := px.Apply2(downloadPath, binPath, func(dp, bp string) string {
		return fmt.Sprintf("mv %s %s", dp, bp)
	})

	delete := px.Apply(binPath, func(bp string) string {
		return fmt.Sprintf("rm -f %s", bp)
	})

	cmd, err := remote.NewCommand(ctx, "mv",
		&remote.CommandArgs{
			Connection: s.conn,
			Create:     px.Cast[p.StringOutput](create),
			Delete:     px.Cast[p.StringOutput](delete),
			Triggers:   px.Cast[p.ArrayOutput](px.All(triggers...)),
		},
		p.Parent(s),
		p.DependsOn([]p.Resource{
			s.Extract,
			s.Mkdir,
		}),
	)

	s.FileName = binName
	s.BinPath = binPath
	s.Mv = cmd

	return err
}

func mapConnection(conn *remote.Connection) *remote.ConnectionArgs {
	if conn == nil {
		return nil
	}

	return &remote.ConnectionArgs{
		AgentSocketPath:    p.StringPtrFromPtr(conn.AgentSocketPath),
		DialErrorLimit:     p.IntPtrFromPtr(conn.DialErrorLimit),
		Host:               p.String(conn.Host),
		Password:           p.StringPtrFromPtr(conn.Password),
		PerDialTimeout:     p.IntPtrFromPtr(conn.PerDialTimeout),
		Port:               p.Float64PtrFromPtr(conn.Port),
		PrivateKey:         p.StringPtrFromPtr(conn.PrivateKey),
		PrivateKeyPassword: p.StringPtrFromPtr(conn.Proxy.PrivateKeyPassword),
	}
}
