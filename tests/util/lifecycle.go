package util

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func LifecycleTest(server integration.Server, l integration.LifeCycleTest) bool {
	urn := resource.NewURN("test", "provider", "", l.Resource, "test")

	runCreate := func(op integration.Operation) p.CreateResponse {
		ginkgo.By("sending a check request")
		check, err := server.Check(p.CheckRequest{
			Urn:  urn,
			Olds: nil,
			News: op.Inputs,
		})

		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(check.Failures).To(gomega.BeEquivalentTo(op.CheckFailures))

		ginkgo.By("sending a preview create request")
		_, err = server.Create(p.CreateRequest{
			Urn:        urn,
			Properties: check.Inputs.Copy(),
			Preview:    true,
		})

		if op.ExpectFailure {
			gomega.Expect(err).To(gomega.HaveOccurred())
		} else {
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		}

		ginkgo.By("sending a create request")
		create, err := server.Create(p.CreateRequest{
			Urn:        urn,
			Properties: check.Inputs.Copy(),
		})

		if op.ExpectFailure {
			gomega.Expect(err).To(gomega.HaveOccurred())
		} else {
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		}

		return create
	}

	runDiff := func(op integration.Operation, id string, olds resource.PropertyMap) (p.DiffResponse, bool) {
		ginkgo.By("sending a check request")
		check, err := server.Check(p.CheckRequest{
			Urn:  urn,
			Olds: olds,
			News: op.Inputs,
		})

		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(check.Failures).To(gomega.BeEquivalentTo(op.CheckFailures))

		ginkgo.By("sending a diff request")
		diff, err := server.Diff(p.DiffRequest{
			ID:   id,
			Urn:  urn,
			Olds: olds,
			News: check.Inputs.Copy(),
		})

		gomega.Expect(err).NotTo(gomega.HaveOccurred())

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

		return diff, isDelete
	}

	runDelete := func(id string, olds resource.PropertyMap, diff p.DiffResponse) {
		runDelete := func() {
			err := server.Delete(p.DeleteRequest{
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
	}

	return ginkgo.Describe("Resource Lifecycle", ginkgo.Ordered, func() {

		createResponse := runCreate(l.Create)

		id := createResponse.ID
		olds := createResponse.Properties
		for i, update := range l.Updates {
			diff, isDelete := runDiff(update, id, olds)
		}
	})
}
