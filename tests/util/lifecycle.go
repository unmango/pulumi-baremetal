package util

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// Based on https://github.com/pulumi/pulumi-go-provider/blob/main/integration/integration.go

func Run(server integration.Server, l integration.LifeCycleTest) {
	urn := resource.NewURN("test", "provider", "", l.Resource, "test")

	runCreate := func(op integration.Operation) (p.CreateResponse, bool) {
		ginkgo.By("sending check request")
		check, err := server.Check(p.CheckRequest{
			Urn:  urn,
			Olds: nil,
			News: op.Inputs,
		})
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		if len(op.CheckFailures) > 0 || len(check.Failures) > 0 {
			gomega.Expect(check.Failures).To(gomega.BeEquivalentTo(op.CheckFailures))
			return p.CreateResponse{}, false
		}

		ginkgo.By("sending preview create request")
		_, err = server.Create(p.CreateRequest{
			Urn:        urn,
			Properties: check.Inputs.Copy(),
			Preview:    true,
		})
		// We allow the failure from ExpectFailure to hit at either the preview or the Create.
		if op.ExpectFailure && err != nil {
			return p.CreateResponse{}, false
		}

		ginkgo.By("sending create request")
		create, err := server.Create(p.CreateRequest{
			Urn:        urn,
			Properties: check.Inputs.Copy(),
		})
		if op.ExpectFailure {
			gomega.Expect(err).To(gomega.HaveOccurred())
			return p.CreateResponse{}, false
		}

		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		if err != nil {
			return p.CreateResponse{}, false
		}
		if op.Hook != nil {
			op.Hook(check.Inputs, create.Properties.Copy())
		}
		if op.ExpectedOutput != nil {
			gomega.Expect(create.Properties).To(gomega.BeEquivalentTo(op.ExpectedOutput))
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
		ginkgo.By("performing update")

		ginkgo.By("sending check request")
		check, err := server.Check(p.CheckRequest{
			Urn:  urn,
			Olds: olds,
			News: update.Inputs,
		})

		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		if err != nil {
			return
		}
		if len(update.CheckFailures) > 0 || len(check.Failures) > 0 {
			gomega.Expect(check.Failures).To(gomega.BeEquivalentTo(update.CheckFailures))
			continue
		}

		ginkgo.By("sending diff request")
		diff, err := server.Diff(p.DiffRequest{
			ID:   id,
			Urn:  urn,
			Olds: olds,
			News: check.Inputs.Copy(),
		})

		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		if err != nil {
			return
		}
		if !diff.HasChanges {
			continue
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
				err = server.Delete(p.DeleteRequest{
					ID:         id,
					Urn:        urn,
					Properties: olds,
				})
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}
			if diff.DeleteBeforeReplace {
				runDelete()
				result, keepGoing := runCreate(update)
				if !keepGoing {
					continue
				}
				id = result.ID
				olds = result.Properties
			} else {
				result, keepGoing := runCreate(update)
				if !keepGoing {
					continue
				}

				runDelete()
				// Set the new block
				id = result.ID
				olds = result.Properties
			}
		} else {

			// Now perform the preview
			_, err = server.Update(p.UpdateRequest{
				ID:      id,
				Urn:     urn,
				Olds:    olds,
				News:    check.Inputs.Copy(),
				Preview: true,
			})

			if update.ExpectFailure && err != nil {
				continue
			}

			result, err := server.Update(p.UpdateRequest{
				ID:   id,
				Urn:  urn,
				Olds: olds,
				News: check.Inputs.Copy(),
			})
			if update.ExpectFailure {
				gomega.Expect(err).To(gomega.HaveOccurred())
				continue
			}
			if update.Hook != nil {
				update.Hook(check.Inputs, result.Properties.Copy())
			}
			if update.ExpectedOutput != nil {
				gomega.Expect(result.Properties.Copy()).To(gomega.BeEquivalentTo(update.ExpectedOutput))
			}
			olds = result.Properties
		}
	}
	err := server.Delete(p.DeleteRequest{
		ID:         id,
		Urn:        urn,
		Properties: olds,
	})
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
}
