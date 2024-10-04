# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.23.2-bookworm AS base
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

FROM --platform=$BUILDPLATFORM base AS download
WORKDIR /build
COPY gen/go.mod gen/go.sum ./gen/
RUN go -C gen mod download

COPY provider/go.mod provider/go.sum ./provider/
RUN go -C provider mod download

FROM --platform=$BUILDPLATFORM download AS build
ARG VERSION=1.0.0-alpha.0+dev
COPY gen/ ./gen/
WORKDIR /build/provider
COPY provider/cmd/pulumi-resource-baremetal/main.go ./cmd/
COPY provider/ ./

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
	go build -o /out/pulumi-resource-baremetal \
	-ldflags "-X github.com/unmango/pulumi-baremetal/provider.Version=${VERSION}" \
	cmd/main.go

FROM --platform=$BUILDPLATFORM scratch AS bin
COPY --from=build /out/pulumi-resource-baremetal /
ENTRYPOINT [ "/pulumi-resource-baremetal" ]

FROM --platform=$BUILDPLATFORM pulumi/pulumi-base:3.135.1-debian AS pulumi
COPY --from=build /out/pulumi-resource-baremetal /plugins/
ENTRYPOINT [ "/bin/bash" ]
