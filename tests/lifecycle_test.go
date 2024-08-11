package tests

import (
	"bytes"
	"context"
	"fmt"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/unmango/pulumi-baremetal/tests/expect"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/asset"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
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

	Describe("Chmod", func() {
		var resource tokens.Type = "baremetal:coreutils:Chmod"

		It("should complete a full lifecycle", func(ctx context.Context) {
			file := containerPath("chmod.txt")

			By("creating a file to modify")
			err := provisioner.WriteFile(ctx, file, []byte("some text"))
			Expect(err).NotTo(HaveOccurred())

			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"files":     []string{file},
							"octalMode": "0700",
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))
					},
				},
			})

			_, err = provisioner.Exec(ctx, "touch", "blah")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Mv", func() {
		var resource tokens.Type = "baremetal:coreutils:Mv"

		It("should complete a full lifecycle", func(ctx context.Context) {
			file := containerPath("mv.txt")
			firstFile := containerPath("mv-new.txt")
			secondFile := containerPath("mv-2.txt")

			By("creating a file to be moved")
			err := provisioner.WriteFile(ctx, file, []byte("some text"))
			Expect(err).NotTo(HaveOccurred())

			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"source":      []string{file},
							"destination": firstFile,
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
							file: firstFile,
						})))
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(context.Background(), file))
						Expect(provisioner).To(ContainFile(context.Background(), firstFile))
					},
				},
				Updates: []integration.Operation{{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"source":      []string{file},
							"destination": secondFile,
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
							file: secondFile,
						})))
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(ctx, file))
						Expect(provisioner).NotTo(ContainFile(ctx, firstFile))
						Expect(provisioner).To(ContainFile(ctx, secondFile))
					},
				}},
			})

			Expect(provisioner).NotTo(ContainFile(ctx, secondFile))
			Expect(provisioner).NotTo(ContainFile(ctx, firstFile))
			Expect(provisioner).To(ContainFile(ctx, file))
		})

		It("should support custom updates", func(ctx context.Context) {
			source := containerPath("mv-custom1.txt")
			dest := containerPath("mv-custom2.txt")
			final := containerPath("mv-custom-final.txt")

			By("creating a file to be moved")
			err := provisioner.WriteFile(ctx, source, []byte("some text"))
			Expect(err).NotTo(HaveOccurred())

			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"source":      []string{source},
							"destination": dest,
						},
						"customUpdate": []string{"mv", dest, final},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
							source: dest,
						})))
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(ctx, source))
						Expect(provisioner).To(ContainFile(ctx, dest))
					},
				},
				Updates: []integration.Operation{{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"source":      []string{"this is kinda nonsensical so I might change how this works in the future"},
							"destination": dest,
						},
						"customUpdate": []string{"mv", dest, final},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(ctx, source))
						Expect(provisioner).NotTo(ContainFile(ctx, dest))
						Expect(provisioner).To(ContainFile(ctx, final))
					},
				}},
			})

			Expect(provisioner).NotTo(ContainFile(ctx, source))
			Expect(provisioner).NotTo(ContainFile(ctx, dest))
			Expect(provisioner).To(ContainFile(ctx, final))
		})

		It("should support custom deletes", func(ctx context.Context) {
			source := containerPath("mv-custom1.txt")
			dest := containerPath("mv-custom2.txt")
			final := containerPath("mv-custom-final.txt")

			By("creating a file to be moved")
			err := provisioner.WriteFile(ctx, source, []byte("some text"))
			Expect(err).NotTo(HaveOccurred())

			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"source":      []string{source},
							"destination": dest,
						},
						"customDelete": []string{"mv", dest, final},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(Equal(pr.NewPropertyMapFromMap(map[string]interface{}{
							source: dest,
						})))
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(ctx, source))
						Expect(provisioner).To(ContainFile(ctx, dest))
					},
				},
			})

			Expect(provisioner).NotTo(ContainFile(ctx, source))
			Expect(provisioner).NotTo(ContainFile(ctx, dest))
			Expect(provisioner).To(ContainFile(ctx, final))
		})
	})

	Describe("Mkdir", func() {
		var resource tokens.Type = "baremetal:coreutils:Mkdir"

		It("should complete a full lifecycle", func(ctx context.Context) {
			expectedDir := containerPath("mkdir-test")

			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"directory": []string{expectedDir},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))
					},
				},
			})

			_, err := provisioner.Exec(ctx, "touch", "blah")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Mktemp", func() {
		var resource tokens.Type = "baremetal:coreutils:Mktemp"

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"tmpdir": true,
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"].V).NotTo(BeEmpty())
						Expect(output["exitCode"].V).To(BeEquivalentTo(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(BeEmpty())
					},
				},
				Updates: []integration.Operation{
					{ // Add a trigger
						Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
							"args": map[string]interface{}{
								"tmpdir": true,
							},
							"triggers": []string{"a trigger"},
						}),
						Hook: func(inputs, output pr.PropertyMap) {
							Expect(output["stderr"]).To(HavePropertyValue(""))
							Expect(output["stdout"].V).NotTo(BeEmpty())
							Expect(output["exitCode"].V).To(BeEquivalentTo(0))
							Expect(output["triggers"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
								pr.NewProperty("a trigger"),
							})))
							Expect(inputs["args"]).To(Equal(output["args"]))
						},
					},
					{ // change a trigger
						Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
							"args": map[string]interface{}{
								"tmpdir": true,
							},
							"triggers": []string{"an updated trigger"},
						}),
						Hook: func(inputs, output pr.PropertyMap) {
							Expect(output["stderr"]).To(HavePropertyValue(""))
							Expect(output["stdout"].V).NotTo(BeEmpty())
							Expect(output["exitCode"].V).To(BeEquivalentTo(0))
							Expect(output["triggers"]).To(Equal(pr.NewArrayProperty([]pr.PropertyValue{
								pr.NewProperty("an updated trigger"),
							})))
							Expect(inputs["args"]).To(Equal(output["args"]))
						},
					},
				},
			})

			_, err := provisioner.Exec(ctx, "touch", "blah")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Rm", func() {
		var resource tokens.Type = "baremetal:coreutils:Rm"

		It("should complete a full lifecycle", func(ctx context.Context) {
			file := containerPath("rm.txt")

			By("creating a file to be removed")
			err := provisioner.WriteFile(ctx, file, []byte("some text"))
			Expect(err).NotTo(HaveOccurred())

			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"files": []string{file},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["exitCode"].V).To(BeEquivalentTo(0))
						Expect(output["createdFiles"].V).To(BeEmpty())
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))
						Expect(provisioner).NotTo(ContainFile(context.Background(), file))
					},
				},
			})

			Expect(provisioner).NotTo(ContainFile(ctx, file))
		})
	})

	Describe("Tar", func() {
		var resource tokens.Type = "baremetal:coreutils:Tar"
		work := containerPath("tar")

		It("should complete a full lifecycle", func(ctx context.Context) {
			fileName := "someFile.txt"
			contents := "Some text that really doesn't matter"
			archive := containerPath("tar", "test-archive.tar.gz")
			dest := containerPath("tar", "destination")
			expectedFile := containerPath("tar", "destination", fileName)

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
						Expect(provisioner).To(ContainFile(context.Background(), expectedFile))
					},
				},
			})

			Expect(provisioner).To(ContainFile(ctx, archive))
			Expect(provisioner).NotTo(ContainFile(ctx, expectedFile))
		})
	})

	Describe("Tee", func() {
		var resource tokens.Type = "baremetal:coreutils:Tee"

		It("should complete a full lifecycle", func(ctx context.Context) {
			file := containerPath("create.txt")
			newFile := containerPath("update.txt")

			By("creating the stdin asset")
			stdin, err := asset.FromText("Test lifecycle stdin")
			Expect(err).NotTo(HaveOccurred())

			By("creating the updated asset")
			newStdin, err := asset.FromText("Updated stdin")
			Expect(err).NotTo(HaveOccurred())

			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"content": stdin,
							"files":   []string{file},
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["stdout"]).To(HavePropertyValue(stdin.Text))
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
						Expect(output["movedFiles"].V).To(BeEmpty())
						Expect(output["args"]).To(Equal(inputs["args"]))

						data, err := provisioner.ReadFile(ctx, file)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(stdin.Text))
					},
				},
				Updates: []integration.Operation{
					{
						Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
							"args": map[string]interface{}{
								"content": stdin,
								"files":   []string{newFile},
							},
						}),
						Hook: func(inputs, output pr.PropertyMap) {
							Expect(output["stderr"]).To(HavePropertyValue(""))
							Expect(output["stdout"]).To(HavePropertyValue(stdin.Text))
							Expect(output["exitCode"]).To(HavePropertyValue(0))
							Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
							Expect(output["movedFiles"].V).To(BeEmpty())
							Expect(output["args"]).To(Equal(inputs["args"]))
							Expect(provisioner).NotTo(ContainFile(ctx, file))

							data, err := provisioner.ReadFile(ctx, newFile)
							Expect(err).NotTo(HaveOccurred())
							Expect(string(data)).To(Equal(stdin.Text))
						},
					},
					{
						Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
							"args": map[string]interface{}{
								"content": newStdin,
								"files":   []string{newFile},
							},
						}),
						Hook: func(inputs, output pr.PropertyMap) {
							Expect(output["stderr"]).To(HavePropertyValue(""))
							Expect(output["stdout"]).To(HavePropertyValue(newStdin.Text))
							Expect(output["exitCode"]).To(HavePropertyValue(0))
							Expect(output["createdFiles"].V).NotTo(BeEmpty()) // TODO: Make this better
							Expect(output["movedFiles"].V).To(BeEmpty())
							Expect(output["args"]).To(Equal(inputs["args"]))

							data, err := provisioner.ReadFile(ctx, newFile)
							Expect(err).NotTo(HaveOccurred())
							Expect(string(data)).To(Equal(newStdin.Text))
						},
					},
					{
						Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
							"args": map[string]interface{}{
								"content": newStdin,
								"files":   []string{file, newFile},
							},
						}),
						Hook: func(inputs, output pr.PropertyMap) {
							Expect(output["stderr"]).To(HavePropertyValue(""))
							Expect(output["stdout"]).To(HavePropertyValue(newStdin.Text))
							Expect(output["exitCode"]).To(HavePropertyValue(0))
							Expect(output["createdFiles"].V).To(HaveLen(2)) // TODO: Make this better
							Expect(output["movedFiles"].V).To(BeEmpty())
							Expect(output["args"]).To(Equal(inputs["args"]))

							data, err := provisioner.ReadFile(ctx, file)
							Expect(err).NotTo(HaveOccurred())
							Expect(string(data)).To(Equal(newStdin.Text))

							data, err = provisioner.ReadFile(ctx, newFile)
							Expect(err).NotTo(HaveOccurred())
							Expect(string(data)).To(Equal(newStdin.Text))
						},
					},
				},
			})

			Expect(provisioner).NotTo(ContainFile(ctx, newFile))
			Expect(provisioner).NotTo(ContainFile(ctx, file))
		})
	})

	Describe("Wget", Ordered, func() {
		var resource tokens.Type = "baremetal:coreutils:Wget"
		dir := containerPath("wget")

		It("should complete a full lifecycle", func(ctx context.Context) {
			url := "https://raw.githubusercontent.com/unmango/pulumi-baremetal/main/README.md"
			file := path.Join(dir, "README.md")

			By("creating a workspace for wget in the container")
			_, err := provisioner.Exec(ctx, "mkdir", "-p", dir)
			Expect(err).NotTo(HaveOccurred())

			run(server, integration.LifeCycleTest{
				Resource: resource,
				Create: integration.Operation{
					Inputs: pr.NewPropertyMapFromMap(map[string]interface{}{
						"args": map[string]interface{}{
							"directoryPrefix": dir,
							"urls":            []string{url},
							"quiet":           true,
						},
					}),
					Hook: func(inputs, output pr.PropertyMap) {
						Expect(output["exitCode"]).To(HavePropertyValue(0))
						Expect(output["stdout"]).To(HavePropertyValue(""))
						Expect(output["stderr"]).To(HavePropertyValue(""))
						Expect(output["createdFiles"].V).To(ContainElement(pr.NewProperty(file)))
						Expect(output["movedFiles"].V).To(BeEmpty())

						args := output["args"].ObjectValue()
						Expect(args["directoryPrefix"]).To(HavePropertyValue(dir))
						Expect(args["urls"].V).To(ContainElement(pr.NewProperty(url)))
						Expect(args["quiet"]).To(HavePropertyValue(true))

						_, err := provisioner.ReadFile(context.Background(), file)
						Expect(err).NotTo(HaveOccurred())
					},
				},
			})

			Expect(provisioner).NotTo(ContainFile(ctx, file))
		})
	})
})

// Based on https://github.com/pulumi/pulumi-go-provider/blob/main/integration/integration.go

func run(server integration.Server, l integration.LifeCycleTest) {
	urn := pr.NewURN("test", "provider", "", l.Resource, "test")

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
			By("expecting failure")
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

		// TODO: This throws, so the next condition will never get hit
		// Double check if this is ok and remove the condition
		Expect(err).NotTo(HaveOccurred())
		if err != nil {
			return p.CreateResponse{}, false
		}
		if op.Hook != nil {
			By("executing the create hook")
			op.Hook(check.Inputs, create.Properties.Copy())
		}
		if op.ExpectedOutput != nil {
			Expect(create.Properties).To(Equal(op.ExpectedOutput))
		}

		return create, true
	}

	createResponse, keepGoing := runCreate(l.Create)

	if !keepGoing {
		By("finishing the test")
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

		// TODO: This throws, so the next condition will never get hit
		// Double check if this is ok and remove the condition
		Expect(err).NotTo(HaveOccurred())
		if err != nil {
			return
		}
		if len(update.CheckFailures) > 0 || len(check.Failures) > 0 {
			Expect(check.Failures).To(Equal(update.CheckFailures))
			By("finishing the test")
			return
		}

		By("sending diff request")
		diff, err := server.Diff(p.DiffRequest{
			ID:   id,
			Urn:  urn,
			Olds: olds,
			News: check.Inputs.Copy(),
		})

		// TODO: This throws, so the next condition will never get hit
		// Double check if this is ok and remove the condition
		Expect(err).NotTo(HaveOccurred())
		if err != nil {
			return
		}
		if !diff.HasChanges {
			By("continuing because no changes")
			continue
		}

		isDelete := false
		for d, v := range diff.DetailedDiff {
			switch v.Kind {
			case p.AddReplace:
				fallthrough
			case p.DeleteReplace:
				fallthrough
			case p.UpdateReplace:
				By(fmt.Sprintf("changing `%s` to trigger %s", d, v.Kind))
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
				By("deleting before replacing")
				runDelete()
				result, keepGoing := runCreate(update)
				if !keepGoing {
					By("finishing the test")
					return
				}
				id = result.ID
				olds = result.Properties
			} else {
				result, keepGoing := runCreate(update)
				if !keepGoing {
					By("finishing the test")
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
				By("expecting failure")
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
				By("executing update hook")
				update.Hook(check.Inputs, result.Properties.Copy())
			}
			if update.ExpectedOutput != nil {
				Expect(result.Properties.Copy()).To(Equal(update.ExpectedOutput))
			}
			olds = result.Properties
		}
	}

	By("sending the final delete request")
	err := server.Delete(p.DeleteRequest{
		ID:         id,
		Urn:        urn,
		Properties: olds,
	})
	Expect(err).NotTo(HaveOccurred())
}
