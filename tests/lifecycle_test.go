package tests

import (
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

const work = "/tmp/lifecycle"

func containerPath(name string) string {
	return path.Join(work, name)
}

var _ = Describe("Command Resources", func() {
	var server integration.Server

	BeforeEach(func(ctx context.Context) {
		By("creating an integration server")
		server = util.NewServer()

		By("creating a workspace in the container")
		_, err := provisioner.Exec(ctx, "mkdir", "-p", work)
		Expect(err).NotTo(HaveOccurred())

		By("fetching provisioner connection details")
		addr, port, err := provisioner.ConnectionDetails(ctx)
		Expect(err).NotTo(HaveOccurred())

		By("configuring the provider")
		err = util.ConfigureProvider(server).
			WithProvisioner(addr, port).
			WithCerts(pcerts).
			Configure()

		Expect(err).NotTo(HaveOccurred())
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
					"content": stdin,
					"files":   []string{file},
				}),
				ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
					"exitCode":     0,
					"stdout":       stdin,
					"stderr":       "",
					"createdFiles": []string{file},
					"args": map[string]interface{}{
						"append":  false,
						"content": stdin,
						"files":   []string{file},
					},
				}),
				Hook: func(inputs, output resource.PropertyMap) {
					data, err := provisioner.ReadFile(context.Background(), file)
					Expect(err).NotTo(HaveOccurred())
					Expect(string(data)).To(Equal(stdin))
				},
			},
			Updates: []integration.Operation{
				{
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"content": stdin,
						"files":   []string{newFile},
					}),
					ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
						"exitCode":     0,
						"stdout":       stdin,
						"stderr":       "",
						"createdFiles": []string{newFile},
						"args": map[string]interface{}{
							"append":  false,
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
				},
				{
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"content": newStdin,
						"files":   []string{newFile},
					}),
					ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
						"exitCode":     0,
						"stdout":       newStdin,
						"stderr":       "",
						"createdFiles": []string{newFile},
						"args": map[string]interface{}{
							"append":  false,
							"content": newStdin,
							"files":   []string{newFile},
						},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						ctx := context.Background()
						data, err := provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))
					},
				},
				{
					Inputs: resource.NewPropertyMapFromMap(map[string]interface{}{
						"content": newStdin,
						"files":   []string{file, newFile},
					}),
					ExpectedOutput: resource.NewPropertyMapFromMap(map[string]interface{}{
						"exitCode":     0,
						"stdout":       newStdin,
						"stderr":       "",
						"createdFiles": []string{file, newFile},
						"args": map[string]interface{}{
							"append":  false,
							"content": newStdin,
							"files":   []string{file, newFile},
						},
					}),
					Hook: func(inputs, output resource.PropertyMap) {
						ctx := context.Background()

						data, err := provisioner.ReadFile(ctx, file)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))

						data, err = provisioner.ReadFile(ctx, newFile)
						Expect(err).NotTo(HaveOccurred())
						Expect(string(data)).To(Equal(newStdin))
					},
				},
			},
		}

		It("should complete a full lifecycle", func(ctx context.Context) {
			run(server, test)

			Expect(provisioner).NotTo(ContainFile(ctx, file))
			Expect(provisioner).NotTo(ContainFile(ctx, newFile))
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
			Expect(create.Properties).To(BeEquivalentTo(op.ExpectedOutput))
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
			Expect(check.Failures).To(BeEquivalentTo(update.CheckFailures))
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
				Expect(result.Properties.Copy()).To(BeEquivalentTo(update.ExpectedOutput))
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
