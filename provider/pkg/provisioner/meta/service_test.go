package meta_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	pb "github.com/unmango/pulumi-baremetal/gen/go/unmango/baremetal/v1alpha1"
	"github.com/unmango/pulumi-baremetal/provider"
	"github.com/unmango/pulumi-baremetal/provider/pkg/provisioner/meta"
)

var _ = Describe("Grpc Server", func() {
	var service pb.MetaServiceServer

	BeforeEach(func() {
		service = meta.NewServer()
	})

	It("should construct", func() {
		Expect(service).NotTo(BeNil())
	})

	It("should ping", func(ctx context.Context) {
		res, err := service.Ping(ctx, &pb.PingRequest{})

		Expect(err).NotTo(HaveOccurred())
		Expect(res).NotTo(BeNil())
		Expect(res.Message).To(Equal("pong"))
	})

	It("should return the current version", func(ctx context.Context) {
		res, err := service.Version(ctx, &pb.VersionRequest{})

		Expect(err).NotTo(HaveOccurred())
		Expect(res).NotTo(BeNil())
		Expect(res.Version).To(Equal(provider.Version))
	})
})
