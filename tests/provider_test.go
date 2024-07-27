// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"context"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	tc "github.com/testcontainers/testcontainers-go"

	baremetal "github.com/unmango/pulumi-baremetal/provider"
)

var _ = Describe("Provider", Ordered, func() {
	var prov integration.Server
	var container tc.Container

	BeforeAll(func(ctx context.Context) {
		ct, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
			ContainerRequest: tc.ContainerRequest{
				FromDockerfile: tc.FromDockerfile{
					Context:    repoRoot,
					Dockerfile: path.Join("tests", "Dockerfile"),
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())
		container = ct

		prov = integration.NewServer(
			baremetal.Name,
			semver.MustParse("1.0.0"),
			baremetal.Provider(),
		)
	})

	It("should create a random", func() {
		response, err := prov.Create(p.CreateRequest{
			Urn: urn("Random"),
			Properties: resource.PropertyMap{
				"length": resource.NewNumberProperty(12),
			},
			Preview: false,
		})

		Expect(err).NotTo(HaveOccurred())
		result := response.Properties["result"].StringValue()
		Expect(result).To(HaveLen(12))
	})

	AfterAll(func(ctx context.Context) {
		err := container.Terminate(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
})

// urn is a helper function to build an urn for running integration tests.
func urn(typ string) resource.URN {
	return resource.NewURN("stack", "proj", "",
		tokens.Type("test:index:"+typ), "name")
}
