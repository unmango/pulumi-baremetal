package tests

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

var _ = Describe("Tar", func() {
	var resource tokens.Type = "baremetal:coreutils:Tar"
	var server integration.Server
	work := containerPath("tar")

	BeforeEach(func(ctx context.Context) {
		server = prepareIntegrationServer(ctx)
	})

	When("archive doesn't exist", func() {
		It("should fail", func() {
			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"file": "/does/not/exist",
						},
					}),
					ExpectFailure: true,
				},
			})
		})
	})

	When("archive contains a single file", func() {
		fileName := "someFile.txt"
		text := "Some text that really doesn't matter"

		archive := containerPath(work, "archive-contains-a-single-file.tar")
		dest := containerPath(work, "destinationA")
		expectedFile := containerPath(dest, fileName)

		BeforeEach(func(ctx context.Context) {
			By("ensuring container directories exist")
			_, err := provisioner.Exec(ctx, "mkdir", "-p", work, dest)
			Expect(err).NotTo(HaveOccurred())

			By("creating an archive to operate on")
			buf, err := createTar(map[string]string{
				fileName: text,
			})
			Expect(err).NotTo(HaveOccurred())

			By("writing the archive to the container")
			err = provisioner.WriteFile(ctx, archive, buf.Bytes())
			Expect(err).NotTo(HaveOccurred())
		})

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"extract":   true,
							"file":      archive,
							"directory": dest,
							"args":      []string{fileName},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["createdFiles"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
							pr.NewProperty(expectedFile),
						})))
						Expect(output["movedFiles"].V).To(Equal(pr.PropertyMap{}))
						Expect(output["args"].V).To(Equal(inputs["args"].V))
						Expect(provisioner).To(ContainFile(ctx, expectedFile))
					},
				},
			})

			Expect(provisioner).To(ContainFile(ctx, archive))
			Expect(provisioner).NotTo(ContainFile(ctx, expectedFile))
		})
	})

	When("archive contents are gzipped", func() {
		fileA := "someFile"
		contents := "Some text that really doesn't matter"
		fileB := "a-different-file"

		archive := containerPath(work, "archive-contents-are-gzipped.tar.gz")
		dest := containerPath(work, "destinationB")
		fileAPath := containerPath(dest, fileA)
		fileBPath := containerPath(dest, fileB)

		BeforeEach(func(ctx context.Context) {
			By("ensuring container directories exist")
			_, err := provisioner.Exec(ctx, "mkdir", "-p", work, dest)
			Expect(err).NotTo(HaveOccurred())

			By("creating an archive to operate on")
			buf, err := createTarGz(map[string]string{
				fileA: contents,
				fileB: "I have a sense of humor",
			})
			Expect(err).NotTo(HaveOccurred())

			By("writing the archive to the container")
			err = provisioner.WriteFile(ctx, archive, buf.Bytes())
			Expect(err).NotTo(HaveOccurred())
		})

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"extract":   true,
							"gzip":      true,
							"file":      archive,
							"directory": dest,
							"args":      []string{fileB},
							"verbose":   true,
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(fileB + "\n"))
						Expect(output["createdFiles"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
							pr.NewProperty(fileBPath),
						})))
						Expect(output["movedFiles"].V).To(Equal(pr.PropertyMap{}))
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(ctx, fileAPath))
						Expect(provisioner).To(ContainFile(ctx, fileBPath))
					},
				},
			})

			Expect(provisioner).To(ContainFile(ctx, archive))
			Expect(provisioner).NotTo(ContainFile(ctx, fileAPath))
			Expect(provisioner).NotTo(ContainFile(ctx, fileBPath))
		})
	})

	When("cni plugins are downloaded", func() {
		dest := containerPath(work, "cni-plugins")
		archive := path.Join("/testdata", "cni-plugins-linux-amd64-v1.5.1.tgz")

		BeforeEach(func(ctx context.Context) {
			By("ensuring container directories exist")
			_, err := provisioner.Exec(ctx, "mkdir", "-p", work, dest)
			Expect(err).NotTo(HaveOccurred())

			By("ensuring the cni-plugins exist")
			Expect(provisioner).To(ContainFile(ctx, archive))
		})

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"extract":    true,
							"gzip":       true,
							"file":       archive,
							"directory":  dest,
							"args":       []string{"macvlan", "vlan"},
							"verbose":    true,
							"noAnchored": true,
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue("./vlan\n./macvlan\n"))
						Expect(output["createdFiles"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
							pr.NewProperty(containerPath(dest, "macvlan")),
							pr.NewProperty(containerPath(dest, "vlan")),
						})))
						Expect(output["movedFiles"].V).To(Equal(pr.PropertyMap{}))
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).To(ContainFile(ctx, containerPath(dest, "macvlan")))
						Expect(provisioner).To(ContainFile(ctx, containerPath(dest, "vlan")))
					},
				},
			})

			Expect(provisioner).To(ContainFile(ctx, archive))
		})
	})
})

func createTar(c map[string]string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	if err := writeContents(buf, c); err != nil {
		return nil, err
	}

	return buf, nil
}

func createTarGz(c map[string]string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	if err := writeArchive(buf, c); err != nil {
		return nil, err
	}

	return buf, nil
}

func writeArchive(w io.Writer, c map[string]string) error {
	writer := gzip.NewWriter(w)
	defer writer.Close()

	return writeContents(writer, c)
}

func writeContents(w io.Writer, c map[string]string) error {
	writer := tar.NewWriter(w)
	defer writer.Close()

	return util.WriteTarContents(writer, c)
}
