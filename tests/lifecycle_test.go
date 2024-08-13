package tests

import (
	"context"
	"fmt"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	pr "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

const workRoot = "/tmp/lifecycle"

func containerPath(elem ...string) string {
	if len(elem) > 0 && !path.IsAbs(elem[0]) {
		elem = append([]string{workRoot}, elem...)
	}

	return path.Join(elem...)
}

func prepareIntegrationServer(ctx context.Context) integration.Server {
	By("creating an integration server")
	server := util.NewServer()

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
	return server
}

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
