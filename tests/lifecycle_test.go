package tests

import (
	"bytes"
	"context"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

const workRoot = "/tmp/lifecycle"

func containerPath(elem ...string) string {
	parts := append([]string{workRoot}, elem...)
	return path.Join(parts...)
}

var _ = Describe("Command Resources", func() {
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		By("creating an integration server")
		server = util.NewServer()

		By("creating a workspace in the container")
		_, err := provisioner.Exec(ctx, "mkdir", "-p", workRoot)
		Expect(err).NotTo(HaveOccurred())

		By("fetching provisioner connection details")
		addr, port, err := provisioner.ConnectionDetails(ctx)
		Expect(err).NotTo(HaveOccurred())

		By("configuring the provider")
		err = util.ConfigureProvider(server).
			WithProvisioner(addr, port).
			WithCerts(provisioner.Ca(), clientCerts.Cert).
			Configure()

		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Chmod", Ordered, func() {
		file := containerPath("chmod.txt")

		BeforeAll(func(ctx context.Context) {
			By("creating a file to modify")
			err := provisioner.WriteFile(ctx, file, []byte("some text"))
			Expect(err).NotTo(HaveOccurred())
		})

		test := integration.LifeCycleTest{
			Resource: "baremetal:cmd:Chmod",
			Create: integration.Operation{
				Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"files":     []string{file},
						"octalMode": "0700",
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			_, err := provisioner.Exec(ctx, "touch", "blah")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Mv", Ordered, func() {
		file := containerPath("mv.txt")
		firstFile := containerPath("mv-new.txt")
		secondFile := containerPath("mv-2.txt")
		customFile := containerPath("mv-custom.txt")

		BeforeAll(func(ctx context.Context) {
			By("creating a file to be moved")
			err := provisioner.WriteFile(ctx, file, []byte("some text"))
			Expect(err).NotTo(HaveOccurred())
		})

		test := integration.LifeCycleTest{
			Resource: "baremetal:cmd:Mv",
			Create: integration.Operation{
				Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"source":      []string{file},
						"destination": firstFile,
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(Equal(resource.NewPropertyMapFromMap(map[string]interface{}{
						file: firstFile,
					})))
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(context.Background(), file))
					Expect(provisioner).To(ContainFile(context.Background(), firstFile))
				},
			},
			Updates: []integration.Operation{
				{
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"source":      []string{file},
							"destination": secondFile,
						},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(Equal(resource.NewPropertyMapFromMap(map[string]interface{}{
							file: secondFile,
						})))
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(context.Background(), file))
						Expect(provisioner).NotTo(ContainFile(context.Background(), firstFile))
						Expect(provisioner).To(ContainFile(context.Background(), secondFile))
					},
				},
				{
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"source":      []string{file},
							"destination": firstFile,
						},
						"customUpdate": []string{"mv", file, customFile},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(Equal(resource.NewPropertyMapFromMap(map[string]interface{}{
							file: firstFile, // TODO: This is wrong
						})))
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(context.Background(), file))
						Expect(provisioner).NotTo(ContainFile(context.Background(), firstFile))
						Expect(provisioner).To(ContainFile(context.Background(), customFile))
					},
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			Expect(provisioner).NotTo(ContainFile(ctx, firstFile))
			Expect(provisioner).To(ContainFile(ctx, file))
		})
	})

	Describe("Mkdir", Ordered, func() {
		expectedDir := containerPath("mkdir-test")

		test := integration.LifeCycleTest{
			Resource: "baremetal:cmd:Mkdir",
			Create: integration.Operation{
				Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"directory": []string{expectedDir},
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			_, err := provisioner.Exec(ctx, "touch", "blah")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Mktemp", Ordered, func() {
		test := integration.LifeCycleTest{
			Resource: "baremetal:cmd:Mktemp",
			Create: integration.Operation{
				Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"tmpdir": true,
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"].V).NotTo(BeEmpty())
					Expect(output["exitCode"].V).To(BeEquivalentTo(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
				},
			},
			Updates: []integration.Operation{
				{ // Add a trigger
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"tmpdir": true,
						},
						"triggers": []string{"a trigger"},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"].V).NotTo(BeEmpty())
						Expect(output["exitCode"].V).To(BeEquivalentTo(0))
						Expect(output["triggers"]).To(Equal(resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewProperty("a trigger"),
						})))
						Expect(inputs["args"]).To(Equal(output["args"]))
					},
				},
				{ // change a trigger
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"tmpdir": true,
						},
						"triggers": []string{"an updated trigger"},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"].V).NotTo(BeEmpty())
						Expect(output["exitCode"].V).To(BeEquivalentTo(0))
						Expect(output["triggers"]).To(Equal(resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewProperty("an updated trigger"),
						})))
						Expect(inputs["args"]).To(Equal(output["args"]))
					},
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			_, err := provisioner.Exec(ctx, "touch", "blah")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Rm", Ordered, func() {
		file := containerPath("rm.txt")

		BeforeAll(func(ctx context.Context) {
			By("creating a file to be removed")
			err := provisioner.WriteFile(ctx, file, []byte("some text"))
			Expect(err).NotTo(HaveOccurred())
		})

		test := integration.LifeCycleTest{
			Resource: "baremetal:cmd:Rm",
			Create: integration.Operation{
				Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"files": []string{file},
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["exitCode"].V).To(BeEquivalentTo(0))
					Expect(output["createdFiles"].V).To(BeEmpty())
					Expect(output["movedFiles"].V).To(BeEmpty())
					Expect(output["args"]).To(Equal(inputs["args"]))
					Expect(provisioner).NotTo(ContainFile(context.Background(), file))
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			Expect(provisioner).NotTo(ContainFile(ctx, file))
		})
	})

	Describe("Tar", Ordered, func() {
		work := containerPath("tar")
		fileName := "someFile.txt"
		contents := "Some text that really doesn't matter"
		archive := containerPath("tar", "test-archive.tar.gz")
		dest := containerPath("tar", "destination")
		expectedFile := containerPath("tar", "destination", fileName)

		BeforeAll(func(ctx context.Context) {
			By("ensuring container directories exist")
			_, err := provisioner.Exec(ctx, "mkdir", "-p", work, dest)
			Expect(err).NotTo(HaveOccurred())

			By("creating an archive to operate on")
			buf := &bytes.Buffer{}
			err = util.CreateTarArchive(buf, map[string]string{
				fileName: contents,
			})
			Expect(err).NotTo(HaveOccurred())

			By("writing the archive to the container")
			err = provisioner.WriteFile(ctx, archive, buf.Bytes())
			Expect(err).NotTo(HaveOccurred())
		})

		test := integration.LifeCycleTest{
			Resource: "baremetal:cmd:Tar",
			Create: integration.Operation{
				Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"extract":   true,
						"file":      archive,
						"directory": dest,
						"args":      []string{fileName},
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["createdFiles"]).To(Equal(resource.NewArrayProperty([]resource.PropertyValue{
						resource.NewProperty(expectedFile),
					})))
					Expect(output["movedFiles"].V).To(Equal(resource.PropertyMap{}))
					Expect(output["args"].V).To(Equal(inputs["args"].V))
					Expect(provisioner).To(ContainFile(context.Background(), expectedFile))
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			Expect(provisioner).To(ContainFile(ctx, archive))
			Expect(provisioner).NotTo(ContainFile(ctx, expectedFile))
		})
	})

	Describe("Tee", Ordered, func() {
		stdin := "Test lifecycle stdin"
		newStdin := "Updated stdin"
		file := containerPath("create.txt")
		newFile := containerPath("update.txt")

		test := integration.LifeCycleTest{
			Resource: "baremetal:cmd:Tee",
			Create: integration.Operation{
				Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"content": stdin,
						"files":   []string{file},
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					Expect(output["stderr"]).To(HavePropertyValue(""))
					data, err := provisioner.ReadFile(context.Background(), file)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(data)).To(Equal(stdin))
				},
				ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
					"exitCode":     0,
					"stdout":       stdin,
					"stderr":       "",
					"createdFiles": []string{file},
					"movedFiles":   map[string]string{},
					"args": map[string]interface{}{
						"append":  false,
						"content": stdin,
						"files":   []string{file},
					},
				}),
			},
			Updates: []integration.Operation{
				{
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"content": stdin,
							"files":   []string{newFile},
						},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						ctx := context.Background()
						Expect(provisioner).NotTo(ContainFile(ctx, file))

						data, err := provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(stdin))
					},
					ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
						"exitCode":     0,
						"stdout":       stdin,
						"stderr":       "",
						"createdFiles": []string{newFile},
						"movedFiles":   map[string]string{},
						"args": map[string]interface{}{
							"append":  false,
							"content": stdin,
							"files":   []string{newFile},
						},
					}),
				},
				{
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"content": newStdin,
							"files":   []string{newFile},
						},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						ctx := context.Background()
						data, err := provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))
					},
					ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
						"exitCode":     0,
						"stdout":       newStdin,
						"stderr":       "",
						"createdFiles": []string{newFile},
						"movedFiles":   map[string]string{},
						"args": map[string]interface{}{
							"append":  false,
							"content": newStdin,
							"files":   []string{newFile},
						},
					}),
				},
				{
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"content": newStdin,
							"files":   []string{file, newFile},
						},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))

						ctx := context.Background()
						data, err := provisioner.ReadFile(ctx, file)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))

						data, err = provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))
					},
					ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
						"exitCode":     0,
						"stdout":       newStdin,
						"stderr":       "",
						"createdFiles": []string{file, newFile},
						"movedFiles":   map[string]string{},
						"args": map[string]interface{}{
							"append":  false,
							"content": newStdin,
							"files":   []string{file, newFile},
						},
					}),
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			Expect(provisioner).NotTo(ContainFile(ctx, newFile))
			Expect(provisioner).NotTo(ContainFile(ctx, file))
		})
	})

	Describe("Wget", Ordered, func() {
		dir := containerPath("wget")
		url := "https://raw.githubusercontent.com/unmango/pulumi-baremetal/main/README.md"
		file := path.Join(dir, "README.md")

		BeforeAll(func(ctx context.Context) {
			By("creating a workspace for wget in the container")
			_, err := provisioner.Exec(ctx, "mkdir", "-p", dir)
			Expect(err).NotTo(HaveOccurred())
		})

		test := integration.LifeCycleTest{
			Resource: "baremetal:cmd:Wget",
			Create: integration.Operation{
				Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
					"args": map[string]interface{}{
						"directoryPrefix": dir,
						"urls":            []string{url},
						"quiet":           true,
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					Expect(output["exitCode"]).To(HavePropertyValue(0))
					Expect(output["stdout"]).To(HavePropertyValue(""))
					Expect(output["stderr"]).To(HavePropertyValue(""))
					Expect(output["createdFiles"].V).To(ContainElement(resource.NewProperty(file)))
					Expect(output["movedFiles"].V).To(BeEmpty())

					args := output["args"].ObjectValue()
					Expect(args["directoryPrefix"]).To(HavePropertyValue(dir))
					Expect(args["urls"].V).To(ContainElement(resource.NewProperty(url)))
					Expect(args["quiet"]).To(HavePropertyValue(true))

					_, err := provisioner.ReadFile(context.Background(), file)
					Expect(err).NotTo(HaveOccurred())
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			Expect(provisioner).NotTo(ContainFile(ctx, file))
		})
	})
})

// Based on https://github.com/pulumi/pulumi-go-provider/blob/main/integration/integration.go

func run(server integration.Server, l integration.LifeCycleTest) {
	urn := resource.NewURN("test", "provider", "", l.Resource, "test")

	runCreate := func(op integration.Operation) (p.CreateResponse, bool) {
		By("sending check request")
		check, err := server.Check(p.CheckRequest{
			Urn:  urn,
			Olds: nil,
			News: op.Inputs,
		})
		Expect(err).NotTo(HaveOccurred())

		if len(op.CheckFailures) > 0 || len(check.Failures) > 0 {
			Expect(check.Failures).To(BeEquivalentTo(op.CheckFailures))
			return p.CreateResponse{}, false
		}

		By("sending preview create request")
		_, err = server.Create(p.CreateRequest{
			Urn:        urn,
			Properties: check.Inputs.Copy(),
			Preview:    true,
		})
		// We allow the failure from ExpectFailure to hit at either the preview or the Create.
		if op.ExpectFailure && err != nil {
			return p.CreateResponse{}, false
		}

		By("sending create request")
		create, err := server.Create(p.CreateRequest{
			Urn:        urn,
			Properties: check.Inputs.Copy(),
		})
		if op.ExpectFailure {
			Expect(err).To(HaveOccurred())
			return p.CreateResponse{}, false
		}

		Expect(err).NotTo(HaveOccurred())
		if err != nil {
			return p.CreateResponse{}, false
		}
		if op.Hook != nil {
			op.Hook(check.Inputs, create.Properties.Copy())
		}
		if op.ExpectedOutput != nil {
			Expect(create.Properties).To(Equal(op.ExpectedOutput))
		}

		return create, true
	}

	createResponse, keepGoing := runCreate(l.Create)

	if !keepGoing {
		return
	}

	id := createResponse.ID
	olds := createResponse.Properties
	for _, update := range l.Updates {
		By("sending check request")
		check, err := server.Check(p.CheckRequest{
			Urn:  urn,
			Olds: olds,
			News: update.Inputs,
		})

		Expect(err).NotTo(HaveOccurred())
		if err != nil {
			return
		}
		if len(update.CheckFailures) > 0 || len(check.Failures) > 0 {
			Expect(check.Failures).To(Equal(update.CheckFailures))
			return
		}

		By("sending diff request")
		diff, err := server.Diff(p.DiffRequest{
			ID:   id,
			Urn:  urn,
			Olds: olds,
			News: check.Inputs.Copy(),
		})

		Expect(err).NotTo(HaveOccurred())
		if err != nil {
			return
		}
		if !diff.HasChanges {
			return
		}

		isDelete := false
		for _, v := range diff.DetailedDiff {
			switch v.Kind {
			case p.AddReplace:
				fallthrough
			case p.DeleteReplace:
				fallthrough
			case p.UpdateReplace:
				isDelete = true
			}
		}
		if isDelete {
			runDelete := func() {
				By("sending a delete request")
				err = server.Delete(p.DeleteRequest{
					ID:         id,
					Urn:        urn,
					Properties: olds,
				})
				Expect(err).NotTo(HaveOccurred())
			}

			if diff.DeleteBeforeReplace {
				runDelete()
				result, keepGoing := runCreate(update)
				if !keepGoing {
					return
				}
				id = result.ID
				olds = result.Properties
			} else {
				result, keepGoing := runCreate(update)
				if !keepGoing {
					return
				}

				runDelete()
				// Set the new block
				id = result.ID
				olds = result.Properties
			}
		} else {
			// Now perform the preview
			By("sending a preview update request")
			_, err = server.Update(p.UpdateRequest{
				ID:      id,
				Urn:     urn,
				Olds:    olds,
				News:    check.Inputs.Copy(),
				Preview: true,
			})

			if update.ExpectFailure && err != nil {
				return
			}

			By("sending an update request")
			result, err := server.Update(p.UpdateRequest{
				ID:   id,
				Urn:  urn,
				Olds: olds,
				News: check.Inputs.Copy(),
			})
			if update.ExpectFailure {
				Expect(err).To(HaveOccurred())
				return
			}
			if update.Hook != nil {
				update.Hook(check.Inputs, result.Properties.Copy())
			}
			if update.ExpectedOutput != nil {
				Expect(result.Properties.Copy()).To(Equal(update.ExpectedOutput))
			}
			olds = result.Properties
		}
	}

	err := server.Delete(p.DeleteRequest{
		ID:         id,
		Urn:        urn,
		Properties: olds,
	})
	Expect(err).NotTo(HaveOccurred())
}
