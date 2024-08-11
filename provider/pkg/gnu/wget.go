package gnu

import (
	"context"
	"fmt"
	"path"
	"slices"

	"github.com/unmango/pulumi-baremetal/provider/pkg/provider/cmd"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type WgetArgs struct {
	cmd.ArgsBase

	AppendOutput       string   `pulumi:"appendOutput,optional"`
	Background         bool     `pulumi:"background,optional"`
	Base               string   `pulumi:"base,optional"`
	CaCertificateFile  string   `pulumi:"caCertificateFile,optional"`
	CaDirectory        string   `pulumi:"caDirectory,optional"`
	Certificate        string   `pulumi:"certificate,optional"`
	CertificateType    string   `pulumi:"certificateType,optional"`
	Config             string   `pulumi:"config,optional"`
	Continue           bool     `pulumi:"continue,optional"`
	CrlFile            string   `pulumi:"crlFile,optional"`
	CutDirs            int      `pulumi:"cutDirs,optional"`
	Debug              bool     `pulumi:"debug,optional"`
	DirectoryPrefix    string   `pulumi:"directoryPrefix,optional"`
	Execute            []string `pulumi:"execute,optional"`
	ForceDirectories   bool     `pulumi:"forceDirectories,optional"`
	ForceHtml          bool     `pulumi:"forceHtml,optional"`
	Help               bool     `pulumi:"help,optional"`
	HttpsOnly          bool     `pulumi:"httpsOnly,optional"`
	Inet4Only          bool     `pulumi:"inet4Only,optional"`
	InputFile          string   `pulumi:"inputFile,optional"`
	KeepSessionCookies bool     `pulumi:"keepSessionCookies,optional"`
	NoClobber          bool     `pulumi:"noClobber,optional"`
	NoDirectories      bool     `pulumi:"noDirectories,optional"`
	NoDnsCache         bool     `pulumi:"noDnsCache,optional"`
	NoVerbose          bool     `pulumi:"noVerbose,optional"`
	OutputDocument     string   `pulumi:"outputDocument,optional"`
	OutputFile         string   `pulumi:"outputFile,optional"`
	Password           string   `pulumi:"password,optional" provider:"secret"`
	PrivateKey         string   `pulumi:"privateKey,optional" provider:"secret"`
	PrivateKeyType     string   `pulumi:"privateKeyType,optional" provider:"secret"`
	Progress           string   `pulumi:"progress,optional"`
	Quiet              bool     `pulumi:"quiet,optional"`
	RandomWait         bool     `pulumi:"randomWait,optional"`
	ReportSpeed        string   `pulumi:"reportSpeed,optional"`
	SaveCookies        string   `pulumi:"saveCookies,optional"`
	ShowProgress       bool     `pulumi:"showProgress,optional"`
	StartPos           string   `pulumi:"startPos,optional"`
	Timestamping       bool     `pulumi:"timestamping,optional"`
	Timeout            string   `pulumi:"timeout,optional"`
	Tries              int      `pulumi:"tries,optional"`
	User               string   `pulumi:"user,optional"`
	UserAgent          string   `pulumi:"userAgent,optional"`
	Urls               []string `pulumi:"urls"`
	Verbose            bool     `pulumi:"verbose,optional"`
	Version            string   `pulumi:"version,optional"`
	Wait               string   `pulumi:"wait,optional"`
}

// Cmd implements CommandArgs.
func (w WgetArgs) Cmd() *pb.Command {
	b := &cmd.B{Args: w.Urls}

	b.Opv(w.AppendOutput, "--append-output")
	b.Op(w.Background, "--background")
	b.Opv(w.Base, "--base")
	b.Opv(w.CaCertificateFile, "--ca-certificate-file")
	b.Opv(w.CaDirectory, "--ca-directory")
	b.Opv(w.Certificate, "--certificate")
	b.Opv(w.CertificateType, "--certificate-type")
	b.Opv(w.Config, "--config")
	b.Op(w.Continue, "--continue")
	b.Opv(w.CrlFile, "--crl-file")
	b.Op(w.Debug, "--debug")
	b.Opv(w.DirectoryPrefix, "--directory-prefix")
	b.Op(w.ForceDirectories, "--force-directories")
	b.Op(w.ForceHtml, "--force-html")
	b.Op(w.Help, "--help")
	b.Op(w.HttpsOnly, "--https-only")
	b.Op(w.Inet4Only, "--inet4-only")
	b.Opv(w.InputFile, "--input-file")
	b.Op(w.KeepSessionCookies, "--keep-session-cookies")
	b.Op(w.NoClobber, "--no-clobber")
	b.Op(w.NoDirectories, "--no-directories")
	b.Op(w.NoDnsCache, "--no-dns-cache")
	b.Op(w.NoVerbose, "--no-verbose")
	b.Opv(w.OutputDocument, "--output-document")
	b.Opv(w.OutputFile, "--output-file")
	b.Opv(w.Password, "--password")
	b.Opv(w.PrivateKey, "--private-key")
	b.Opv(w.PrivateKeyType, "--private-key-type")
	b.Opv(w.Progress, "--progress")
	b.Op(w.Quiet, "--quiet")
	b.Op(w.RandomWait, "--random-wait")
	b.Opv(w.ReportSpeed, "--report-speed")
	b.Opv(w.SaveCookies, "--save-cookies")
	b.Op(w.ShowProgress, "--show-progress")
	b.Opv(w.StartPos, "--start-pos")
	b.Opv(w.Timeout, "--timeout")
	b.Op(w.Timestamping, "--timestamping")
	b.Opv(w.User, "--user")
	b.Opv(w.UserAgent, "--user-agent")
	b.Op(w.Verbose, "--verbose")
	b.Opv(w.Version, "--version")
	b.Opv(w.Wait, "--wait")

	for _, e := range w.Execute {
		b.Opv(e, "--execute")
	}

	return &pb.Command{
		Bin:  pb.Bin_BIN_WGET,
		Args: b.Args,
	}
}

// ExpectCreated implements FileManipulator.
func (w WgetArgs) ExpectCreated() []string {
	files := []string{}
	if w.AppendOutput != "" {
		files = append(files, w.AppendOutput)
	}
	if w.OutputDocument != "" && w.OutputDocument != "-" {
		files = append(files, w.OutputDocument)
	}
	if w.DirectoryPrefix != "" {
		for _, u := range w.Urls {
			f := path.Base(u)
			p := path.Join(w.DirectoryPrefix, f)
			files = append(files, p)
		}
	}

	// This notably does not handle a bare command i.e. `wget https://example.com`
	// I'm not really sure how to figure out what file it creates without doint some magic in the middle.
	return files
}

type Wget struct{}

type WgetState = cmd.State[WgetArgs]

// Create implements infer.CustomCreate.
func (Wget) Create(ctx context.Context, name string, inputs cmd.CommandArgs[WgetArgs], preview bool) (id string, output WgetState, err error) {
	state := WgetState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("wget: %w", err)
	}

	return name, state, nil
}

// Diff implements infer.CustomDiff.
func (Wget) Diff(ctx context.Context, id string, olds WgetState, news cmd.CommandArgs[WgetArgs]) (provider.DiffResponse, error) {
	diff, err := olds.Diff(ctx, news)
	if err != nil {
		return provider.DiffResponse{}, fmt.Errorf("wget: %w", err)
	}

	if !slices.Equal(news.Args.Urls, olds.Args.Urls) {
		diff["args.urls"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Quiet != olds.Args.Quiet {
		diff["args.quiet"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Base != olds.Args.Base {
		diff["args.base"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.DirectoryPrefix != olds.Args.DirectoryPrefix {
		diff["args.directoryPrefix"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.ForceDirectories != olds.Args.ForceDirectories {
		diff["args.foreceDirectories"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.OutputFile != olds.Args.OutputFile {
		diff["args.outputFile"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.OutputDocument != olds.Args.OutputDocument {
		diff["args.outputDocument"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if news.Args.Timestamping != olds.Args.Timestamping {
		diff["args.timestamping"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

// Update implements infer.CustomUpdate.
func (Wget) Update(ctx context.Context, id string, olds WgetState, news cmd.CommandArgs[WgetArgs], preview bool) (WgetState, error) {
	state, err := olds.Update(ctx, news, preview)
	if err != nil {
		return olds, fmt.Errorf("wget: %w", err)
	}

	return state, nil
}

// Delete implements infer.CustomDelete.
func (Wget) Delete(ctx context.Context, id string, props WgetState) error {
	if err := props.Delete(ctx); err != nil {
		return fmt.Errorf("wget: %w", err)
	}

	return nil
}

var _ = (infer.CustomCreate[cmd.CommandArgs[WgetArgs], WgetState])((*Wget)(nil))
var _ = (infer.CustomDiff[cmd.CommandArgs[WgetArgs], WgetState])((*Wget)(nil))
var _ = (infer.CustomUpdate[cmd.CommandArgs[WgetArgs], WgetState])((*Wget)(nil))
var _ = (infer.CustomDelete[WgetState])((*Wget)(nil))
