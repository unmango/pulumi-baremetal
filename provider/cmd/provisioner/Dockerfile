# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.23.5-bookworm AS base
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

FROM --platform=$BUILDPLATFORM base AS download
WORKDIR /build/gen
COPY gen/go.mod gen/go.sum ./
RUN go mod download

WORKDIR /build/provider
COPY provider/go.mod provider/go.sum ./
RUN go mod download

FROM --platform=$BUILDPLATFORM download AS build
ARG VERSION=1.0.0-alpha.0+dev
WORKDIR /build/gen
COPY gen/go ./go

WORKDIR /build/provider
COPY provider/cmd/provisioner/main.go ./cmd/
COPY provider/*.go ./
COPY provider/pkg/ ./pkg/

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
	go build -o /out/provisioner \
	-ldflags "-X github.com/unmango/pulumi-baremetal/provider.Version=${VERSION}" \
	cmd/main.go

FROM --platform=$BUILDPLATFORM ubuntu:noble-20241015 AS test-base

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update \
	&& apt install -y \
	curl \
	wget \
	&& rm -rf /var/lib/apt/lists/* \
	&& mkdir /testdata \
	&& curl -fLo /testdata/cni-plugins-linux-amd64-v1.5.1.tgz https://github.com/containernetworking/plugins/releases/download/v1.5.1/cni-plugins-linux-amd64-v1.5.1.tgz

FROM test-base AS test
COPY --from=build /out/provisioner /
ENTRYPOINT ["/provisioner"]

FROM --platform=$BUILDPLATFORM scratch
COPY --from=build /out/provisioner /
ENTRYPOINT ["/provisioner"]
