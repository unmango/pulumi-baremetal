package cmd_test

import (
	"context"
	"fmt"
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

	Describe("Mv", func() {
		var work string

		BeforeEach(func() {
			By("creating a temp workspace")
			tmp, err := os.MkdirTemp("", "")
			Expect(err).NotTo(HaveOccurred())
			Expect(tmp).To(BeADirectory())
			work = tmp
		})

		It("should move a file", func(ctx context.Context) {
			var file, newFile string
			var command *pb.Command
			var create *pb.CreateResponse

			file = path.Join(work, "mv.txt")
			newFile = path.Join(work, "mv-new.txt")
			command = &pb.Command{
				Bin:  pb.Bin_BIN_MV,
				Args: []string{file, newFile},
			}

			By(fmt.Sprintf("writing a file to %s", file))
			err := os.WriteFile(file, []byte("NA"), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())

			By("creating the command")
			create, err = service.Create(ctx, &pb.CreateRequest{
				Command: command,
				ExpectMoved: map[string]string{
					file: newFile,
				},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(create).NotTo(BeNil())
			Expect(create.Result).NotTo(BeNil())
			Expect(create.Result.ExitCode).To(BeEquivalentTo(0))
			Expect(create.Result.Stdout).To(BeEmpty())
			Expect(create.CreatedFiles).To(BeEmpty())
			Expect(create.MovedFiles).To(HaveKeyWithValue(file, newFile))
			Expect(file).NotTo(BeARegularFile())
			Expect(newFile).To(BeARegularFile())

			By("deleting the command")
			res, err := service.Delete(ctx, &pb.DeleteRequest{
				Previous: &pb.Operation{
					Command:      command,
					Result:       create.Result,
					CreatedFiles: create.CreatedFiles,
					MovedFiles:   create.MovedFiles,
				},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.Commands).To(HaveLen(1))
			operation := res.Commands[0]
			Expect(operation.Result).NotTo(BeNil())
			Expect(operation.Result.ExitCode).To(BeEquivalentTo(0))
			Expect(operation.Result.Stdout).To(BeEmpty())
			Expect(operation.CreatedFiles).To(BeEmpty())
			Expect(operation.MovedFiles).To(HaveKeyWithValue(newFile, file))
			Expect(file).To(BeARegularFile())
			Expect(newFile).NotTo(BeARegularFile())
		})

		AfterEach(func() {
			By("cleaning up the workspace")
			err := os.RemoveAll(work)
			Expect(err).NotTo(HaveOccurred())
		})
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
				ExpectCreated: []string{expectedPath},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
			Expect(res.Result).NotTo(BeNil())
			Expect(res.Result.ExitCode).To(BeEquivalentTo(0))
			Expect(res.Result.Stdout).To(BeEmpty())
			Expect(res.Result.Stderr).To(ContainSubstring(url))
			Expect(res.CreatedFiles).To(ContainElement(expectedPath))
			Expect(expectedPath).To(BeARegularFile())
		})

		AfterEach(func() {
			err := os.RemoveAll(work)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
