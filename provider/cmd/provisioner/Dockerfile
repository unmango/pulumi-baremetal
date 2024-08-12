# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.22.6-bookworm AS base
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

FROM --platform=$BUILDPLATFORM base AS download
WORKDIR /build/gen
COPY gen/go.mod gen/go.sum ./
RUN go mod download
COPY gen/go ./go

WORKDIR /build/provider
COPY provider/go.mod provider/go.sum ./
RUN go mod download

FROM --platform=$BUILDPLATFORM download AS build
WORKDIR /build/gen
COPY gen/go ./go

WORKDIR /build/provider
COPY provider/cmd/provisioner/main.go ./cmd/
COPY provider/*.go ./
COPY provider/pkg/ ./pkg/

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /out/provisioner cmd/main.go

FROM --platform=$BUILDPLATFORM ubuntu:noble-20240605 AS test

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update \
    && apt install -y \
        wget \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /out/provisioner /
ENTRYPOINT ["/provisioner"]

FROM --platform=$BUILDPLATFORM scratch
COPY --from=build /out/provisioner /
ENTRYPOINT ["/provisioner"]
