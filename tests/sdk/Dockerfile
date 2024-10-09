# syntax=docker/dockerfile:1
ARG SDK=dotnet
FROM --platform=$BUILDPLATFORM pulumi/pulumi-${SDK}:3.136.1-debian AS base

ARG TARGETOS
ARG TARGETARCH

ARG SDK
ARG GOVERSION=1.23.1

WORKDIR /usr/local
RUN curl --fail -L -o- https://go.dev/dl/go${GOVERSION}.${TARGETOS}-${TARGETARCH}.tar.gz | tar -zxv

ENV GOBIN=/usr/local/bin
ENV GOLANG_VERSION 1.23.1
ENV GOTOOLCHAIN=local
ENV PATH=/usr/local/go/bin:${GOBIN}:${PATH}

RUN go install github.com/onsi/ginkgo/v2/ginkgo@v2

WORKDIR /work
COPY gen/go.* ./gen/
COPY tests/go.* ./tests/
COPY provider/go.* ./provider/

WORKDIR /work/tests
RUN go mod download

WORKDIR /work
COPY gen/ ./gen/
COPY tests/ ./tests
COPY provider/ ./provider/
COPY bin/pulumi-resource-baremetal ./bin/

ENV SDK=${SDK}
ENTRYPOINT [ "ginkgo" ]
CMD [ "run", "-v", "--silence-skips", "/work/tests/sdk" ]

FROM base AS dotnet
COPY sdk/dotnet/ sdk/dotnet/

ENV NUGET_PATH=/root/.pulumi-dev/nuget/
RUN dotnet build ./sdk/dotnet/UnMango.Baremetal.csproj && \
	dotnet pack ./sdk/dotnet/UnMango.Baremetal.csproj --output ${NUGET_PATH} && \
	dotnet nuget add source ${NUGET_PATH} --name docker

COPY examples/dotnet/ ./examples/dotnet/

FROM base AS nodejs
COPY sdk/nodejs/package.json ./sdk/nodejs/
RUN cd sdk/nodejs && yarn install

COPY sdk/nodejs/ ./sdk/nodejs/
RUN cd sdk/nodejs && \
	yarn run tsc && \
	cp package.json yarn.lock bin/ && \
	yarn link --cwd $(pwd)/bin

COPY examples/nodejs/ ./examples/nodejs/
