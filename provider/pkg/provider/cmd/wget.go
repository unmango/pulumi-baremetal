package cmd

import (
	"context"
	"fmt"
	"path"

	"github.com/pulumi/pulumi-go-provider/infer"
	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
)

type WgetArgs struct {
	CommandArgs

	AppendOutput      string   `pulumi:"appendOutput,optional"`
	Background        bool     `pulumi:"background,optional"`
	Base              string   `pulumi:"base,optional"`
	CaCertificateFile string   `pulumi:"caCertificateFile,optional"`
	CaDirectory       string   `pulumi:"caDirectory,optional"`
	Certificate       string   `pulumi:"certificate,optional"`
	CertificateType   string   `pulumi:"certificateType,optional"`
	Config            string   `pulumi:"config,optional"`
	Continue          bool     `pulumi:"continue,optional"`
	CrlFile           string   `pulumi:"crlFile,optional"`
	CutDirs           int      `pulumi:"cutDirs,optional"`
	Debug             bool     `pulumi:"debug,optional"`
	DirectoryPrefix   string   `pulumi:"directoryPrefix,optional"`
	Execute           []string `pulumi:"execute,optional"`
	// ExpectedFiles      []string `pulumi:"expectedFiles,optional"`
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
	b := &builder{w.Urls}
	b.opv(w.AppendOutput, "--append-output")
	b.op(w.Background, "--background")
	b.opv(w.Base, "--base")
	b.opv(w.CaCertificateFile, "--ca-certificate-file")
	b.opv(w.CaDirectory, "--ca-directory")
	b.opv(w.Certificate, "--certificate")
	b.opv(w.CertificateType, "--certificate-type")
	b.opv(w.Config, "--config")
	b.op(w.Continue, "--continue")
	b.opv(w.CrlFile, "--crl-file")
	b.op(w.Debug, "--debug")
	b.opv(w.DirectoryPrefix, "--directory-prefix")
	b.op(w.ForceDirectories, "--force-directories")
	b.op(w.ForceHtml, "--force-html")
	b.op(w.Help, "--help")
	b.op(w.HttpsOnly, "--https-only")
	b.op(w.Inet4Only, "--inet4-only")
	b.opv(w.InputFile, "--input-file")
	b.op(w.KeepSessionCookies, "--keep-session-cookies")
	b.op(w.NoClobber, "--no-clobber")
	b.op(w.NoDirectories, "--no-directories")
	b.op(w.NoDnsCache, "--no-dns-cache")
	b.op(w.NoVerbose, "--no-verbose")
	b.opv(w.OutputDocument, "--output-document")
	b.opv(w.OutputFile, "--output-file")
	b.opv(w.Password, "--password")
	b.opv(w.PrivateKey, "--private-key")
	b.opv(w.PrivateKeyType, "--private-key-type")
	b.opv(w.Progress, "--progress")
	b.op(w.Quiet, "--quiet")
	b.op(w.RandomWait, "--random-wait")
	b.opv(w.ReportSpeed, "--report-speed")
	b.opv(w.SaveCookies, "--save-cookies")
	b.op(w.ShowProgress, "--show-progress")
	b.opv(w.StartPos, "--start-pos")
	b.opv(w.Timeout, "--timeout")
	b.op(w.Timestamping, "--timestamping")
	b.opv(w.User, "--user")
	b.opv(w.UserAgent, "--user-agent")
	b.op(w.Verbose, "--verbose")
	b.opv(w.Version, "--version")
	b.opv(w.Wait, "--wait")

	for _, e := range w.Execute {
		b.opv(e, "--execute")
	}

	return &pb.Command{
		Bin:  pb.Bin_BIN_WGET,
		Args: b.args,
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

type WgetState = CommandState[WgetArgs]

// Create implements infer.CustomCreate.
func (Wget) Create(ctx context.Context, name string, inputs WgetArgs, preview bool) (id string, output WgetState, err error) {
	state := WgetState{}
	if err := state.Create(ctx, inputs, preview); err != nil {
		return name, state, fmt.Errorf("wget: %w", err)
	}

	return name, state, nil
}

// Update implements infer.CustomUpdate.
func (Wget) Update(ctx context.Context, id string, olds WgetState, news WgetArgs, preview bool) (WgetState, error) {
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

var _ = (infer.CustomCreate[WgetArgs, WgetState])((*Wget)(nil))
var _ = (infer.CustomUpdate[WgetArgs, WgetState])((*Wget)(nil))
var _ = (infer.CustomDelete[WgetState])((*Wget)(nil))
