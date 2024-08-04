package cmd_test

import (
	"context"
	"log/slog"
	"os"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider/pkg/internal"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provisioner/cmd"
)

var _ = Describe("Grpc Server", func() {
	var service pb.CommandServiceServer

	BeforeEach(func() {
		service = cmd.NewServer(internal.State{
			Log: slog.New(slog.NewJSONHandler(GinkgoWriter, nil)),
		})
	})

	It("should construct", func() {
		Expect(service).NotTo(BeNil())
	})

	Describe("Wget", func() {
		var work string

		BeforeEach(func() {
			tmp, err := os.MkdirTemp("", "")
			Expect(err).NotTo(HaveOccurred())
			Expect(tmp).To(BeADirectory())
			work = tmp
		})

		It("should download file to directory prefix", func(ctx context.Context) {
			url := "https://raw.githubusercontent.com/unmango/pulumi-baremetal/main/README.md"
			expectedPath := path.Join(work, "README.md")

			res, err := service.Create(ctx, &pb.CreateRequest{
				Command: &pb.Command{
					Bin: pb.Bin_BIN_WGET,
					Args: []string{
						"--directory-prefix", work,
						"--no-verbose",
						"--no-netrc",
						url,
					},
				},
				ExpectFiles: []string{expectedPath},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.Result.ExitCode).To(BeEquivalentTo(0))
			Expect(res.Result.Stdout).To(BeEmpty())
			Expect(res.Result.Stderr).To(ContainSubstring(url))
			Expect(res.Files).To(ContainElement(expectedPath))
			Expect(expectedPath).To(BeARegularFile())
		})

		AfterEach(func() {
			err := os.RemoveAll(work)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})