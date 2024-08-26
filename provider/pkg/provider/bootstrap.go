package provider

import (
	"bytes"
	"context"
	"fmt"
	"path"

	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/internal/logger"
)

const (
	repo = "https://github.com/unmango/pulumi-baremetal"
	bin  = "provisioner"
)

type Bootstrap struct{}

type BootstrapArgs struct {
	Arch       string      `pulumi:"arch,optional"`
	Connection *Connection `pulumi:"connection,optional"`
	Directory  string      `pulumi:"directory,optional"`
	Os         string      `pulumi:"os,optional"`
	Version    *string     `pulumi:"version"`
}

// Annotate implements infer.Annotated.
func (b *BootstrapArgs) Annotate(a infer.Annotator) {
	a.Describe(&b.Arch, "The CPU architecture of the remote system.")
	a.SetDefault(&b.Arch, "amd64")
	a.Describe(&b.Directory, "The directory to store the provisioner binary.")
	a.SetDefault(&b.Directory, "/usr/local/bin")
	a.Describe(&b.Os, "The OS type of the remote system.")
	a.SetDefault(&b.Os, "linux")
	a.Describe(&b.Version, "The version of the provisioner to bootstrap.")
}

var _ = (infer.Annotated)((*BootstrapArgs)(nil))

type BootstrapState struct {
	BootstrapArgs

	ArchiveName string `pulumi:"archiveName"`
	BinPath     string `pulumi:"binPath"`
	Url         string `pulumi:"url"`
}

// Annotate implements infer.Annotated.
func (b *BootstrapState) Annotate(a infer.Annotator) {
	a.Describe(&b.ArchiveName, "Name part of the provisioner release archive file.")
	a.Describe(&b.BinPath, "Provisioner binary path on the remote system.")
	a.Describe(&b.Url, "Url of the provisioner release archive.")
}

var _ = (infer.Annotated)((*BootstrapState)(nil))

// Create implements infer.CustomCreate.
func (Bootstrap) Create(ctx context.Context, name string, inputs BootstrapArgs, preview bool) (id string, output BootstrapState, err error) {
	log := logger.FromContext(ctx)
	state := BootstrapState{BootstrapArgs: inputs}

	client, err := inputs.Connection.Dial(ctx)
	if err != nil {
		return name, state, err
	}

	session, err := client.NewSession()
	if err != nil {
		return name, state, err
	}
	defer session.Close()

	version := "latest"
	if inputs.Version != nil {
		version = *inputs.Version
	}

	archive := fmt.Sprintf("pulumi-resource-baremetal-v%s-%s-%s.tar.gz", version, inputs.Os, inputs.Arch)
	url := fmt.Sprintf("%s/releases/%s/assets/%s", repo, version, archive)
	binPath := path.Join(inputs.Directory, bin)
	log.Debugf("Using URL: %s", url)

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	cmd := fmt.Sprintf(`
tmp="$(mktemp --directory)"
wget --directory-prefix $tmp %s
mkdir --parents %s
tar -C $tmp -xzvf %s --strip-components 1
mv $tmp/provisioner %s/provisioner
`, url, inputs.Directory, archive, inputs.Directory)

	err = session.Run(cmd)
	if err != nil {
		log.Info(stdout.String())
		log.Info(stderr.String())
		return name, state, err
	}

	state.ArchiveName = archive
	state.BinPath = binPath
	state.Url = url

	return name, state, nil
}

var _ = (infer.CustomCreate[BootstrapArgs, BootstrapState])((*Bootstrap)(nil))

// Construct implements infer.ComponentResource.
// func (*Bootstrap) Construct(ctx *p.Context, name string, typ string, inputs BootstrapArgs, opts p.ResourceOption) (*BootstrapState, error) {
// 	state := &BootstrapState{conn: mapConnection(inputs.Connection)}
// 	err := ctx.RegisterComponentResource("baremetal", name, state, opts)
// 	if err != nil {
// 		return nil, err
// 	}

// 	version := inputs.Version
// 	archiveName := fmt.Sprintf("pulumi-resource-baremetal-v%s-linux-arm64.tar.gz", version)
// 	url := fmt.Sprintf("https://github.com/unmango/pulumi-baremetal/releases/download/v%s/%s", version, archiveName)

// 	_ = ctx.Log.Debug("Mktemp", nil)
// 	if err := state.mktemp(ctx, p.Any(version)); err != nil {
// 		return nil, err
// 	}

// 	_ = ctx.Log.Debug("Download", nil)
// 	if err := state.download(ctx, url); err != nil {
// 		return nil, err
// 	}

// 	_ = ctx.Log.Debug("Mkdir", nil)
// 	if err := state.mkdir(ctx, inputs.Directory); err != nil {
// 		return nil, err
// 	}

// 	_ = ctx.Log.Debug("Extract", nil)
// 	if err := state.extract(ctx, archiveName); err != nil {
// 		return nil, err
// 	}

// 	_ = ctx.Log.Debug("Mv", nil)
// 	if err := state.mv(ctx, inputs.Directory, p.Any(version)); err != nil {
// 		return nil, err
// 	}

// 	return state, nil
// }

// func (s *BootstrapState) mktemp(ctx *p.Context, triggers ...px.Input[any]) error {
// 	cmd, err := remote.NewCommand(ctx, "mktemp", &remote.CommandArgs{
// 		Connection: s.conn,
// 		Create:     p.StringPtr("mktemp --directory"),
// 		Triggers:   px.Cast[p.ArrayOutput](px.All(triggers...)),
// 	}, p.Parent(s))
// 	if err != nil {
// 		return err
// 	}

// 	tempDir, err := px.ConvertTyped[string](cmd.Stdout)
// 	if err != nil {
// 		return err
// 	}

// 	s.TempDir = tempDir
// 	s.Mktemp = cmd

// 	return nil
// }

// func (s *BootstrapState) download(ctx *p.Context, url string) error {
// 	create := fx.Sprintf("wget --directory-prefix %s %s",
// 		s.TempDir.AsAny(),
// 		px.Val(url).AsAny())

// 	cmd, err := remote.NewCommand(ctx, "download", &remote.CommandArgs{
// 		Connection: s.conn,
// 		Create:     px.Cast[p.StringOutput](create),
// 	}, p.Parent(s), p.DependsOn([]p.Resource{s.Mktemp}))

// 	s.Download = cmd
// 	s.Url = url

// 	return err
// }

// func (s *BootstrapState) mkdir(ctx *p.Context, dir p.StringInput) error {
// 	create := px.Apply(dir.ToStringOutput(), func(dir string) string {
// 		return fmt.Sprintf("mkdir --parents %s", dir)
// 	})

// 	cmd, err := remote.NewCommand(ctx, "mkdir", &remote.CommandArgs{
// 		Connection: s.conn,
// 		Create:     px.Cast[p.StringOutput](create),
// 	}, p.Parent(s))

// 	s.Mkdir = cmd
// 	return err
// }

// func (s *BootstrapState) extract(ctx *p.Context, archiveName string) error {
// 	create := px.Apply(s.TempDir, func(dir string) string {
// 		return strings.Join([]string{
// 			"tar", "--extract", "--gzip",
// 			"--strip-components", "1",
// 			"--file", path.Join(dir, archiveName),
// 			"--directory", dir,
// 			"--verbose",
// 		}, " ")
// 	})

// 	cmd, err := remote.NewCommand(ctx, "extract",
// 		&remote.CommandArgs{
// 			Connection: s.conn,
// 			Create:     px.Cast[p.StringOutput](create),
// 		},
// 		p.Parent(s),
// 		p.DependsOn([]p.Resource{
// 			s.Mktemp,
// 			s.Download,
// 		}),
// 	)

// 	s.ArchiveName = px.Val(archiveName)
// 	s.Extract = cmd

// 	return err
// }

// func (s *BootstrapState) mv(ctx *p.Context, dir p.StringInput, triggers ...px.Input[any]) error {
// 	downloadPath := px.Apply(s.TempDir, func(tmp string) string {
// 		return path.Join(tmp, binName)
// 	})

// 	binPath := px.Apply(dir.ToStringOutput(), func(dir string) string {
// 		return path.Join(dir, binName)
// 	})

// 	create := px.Apply2(downloadPath, binPath, func(dp, bp string) string {
// 		return fmt.Sprintf("mv %s %s", dp, bp)
// 	})

// 	delete := px.Apply(binPath, func(bp string) string {
// 		return fmt.Sprintf("rm -f %s", bp)
// 	})

// 	cmd, err := remote.NewCommand(ctx, "mv",
// 		&remote.CommandArgs{
// 			Connection: s.conn,
// 			Create:     px.Cast[p.StringOutput](create),
// 			Delete:     px.Cast[p.StringOutput](delete),
// 			Triggers:   px.Cast[p.ArrayOutput](px.All(triggers...)),
// 		},
// 		p.Parent(s),
// 		p.DependsOn([]p.Resource{
// 			s.Extract,
// 			s.Mkdir,
// 		}),
// 	)

// 	s.FileName = binName
// 	s.BinPath = binPath
// 	s.Mv = cmd

// 	return err
// }
