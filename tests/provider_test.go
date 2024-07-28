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
	"fmt"
	"path"
	"time"

	"github.com/docker/go-connections/nat"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	baremetal "github.com/unmango/pulumi-baremetal/provider"
)

const protocol string = "tcp"

var _ = Describe("Provider", Ordered, func() {
	var server integration.Server
	var provisioner tc.Container
	var provisionerPort nat.Port

	BeforeAll(func(ctx context.Context) {
		By("selecting a port")
		port, err := nat.NewPort(protocol, "6969")
		Expect(err).NotTo(HaveOccurred())
		provisionerPort = port

		By("creating a generic container")
		container, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
			ContainerRequest: tc.ContainerRequest{
				FromDockerfile: tc.FromDockerfile{
					Context:    repoRoot,
					Dockerfile: path.Join("provider", "cmd", "provisioner", "Dockerfile"),
				},
				Cmd: []string{
					"--network", protocol,
					"--address", fmt.Sprintf("localhost:%d", port.Int()),
				},
				ExposedPorts: []string{provisionerPort.Port()},
				WaitingFor:   wait.ForListeningPort(provisionerPort),
				LogConsumerCfg: &tc.LogConsumerConfig{
					Consumers: []tc.LogConsumer{LogToWriter(GinkgoWriter)},
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())
		provisioner = container

		By("creating a provider server")
		server = integration.NewServer(
			baremetal.Name,
			semver.MustParse("1.0.0"),
			baremetal.Provider(),
		)
	})

	It("should create a tee", func() {
		By("creating the resource")
		response, err := server.Create(p.CreateRequest{
			Urn: urn("Tee"),
			Properties: resource.PropertyMap{
				"stdin": resource.NewStringProperty("test"),
				"files": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewStringProperty("test"),
				}),
			},
			Preview: false,
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeNil())
	})

	AfterAll(func(ctx context.Context) {
		timeout := time.Duration(10 * time.Second)

		By("stopping the container")
		err := provisioner.Stop(ctx, &timeout)
		Expect(err).NotTo(HaveOccurred())
	})
})

// urn is a helper function to build an urn for running integration tests.
func urn(typ string) resource.URN {
	return resource.NewURN("stack", "proj", "",
		tokens.Type("test:index:"+typ), "name")
}
