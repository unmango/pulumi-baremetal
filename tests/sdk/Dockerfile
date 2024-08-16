# syntax=docker/dockerfile:1
ARG SDK=dotnet
FROM --platform=$BUILDPLATFORM pulumi/pulumi-${SDK}:3.129.0-debian
ARG SDK=dotnet
ARG TARGETOS
ARG TARGETARCH

WORKDIR /usr/local
RUN curl --fail -L -o- https://go.dev/dl/go1.23.0.${TARGETOS}-${TARGETARCH}.tar.gz | tar -zxv
ENV GOBIN=/usr/local/bin
ENV PATH=${PATH}:/usr/local/go/bin:${GOBIN}
RUN go install github.com/onsi/ginkgo/v2/ginkgo@v2

WORKDIR /work
COPY gen/go.* ./gen/
COPY tests/go.* ./tests/
COPY provider/go.* ./provider/

WORKDIR /work/tests
RUN go mod download

WORKDIR /work
COPY gen/ ./gen/
COPY examples/${SDK}/ ./examples/${SDK}/
COPY tests/ ./tests
COPY provider/ ./provider/
COPY nuget/ /root/.pulumi-dev/nuget/
RUN dotnet nuget add source /root/.pulumi-dev/nuget/ --name docker
COPY bin/pulumi-resource-baremetal ./bin/

WORKDIR /work/tests/sdk
ENTRYPOINT [ "ginkgo" ]
CMD [ "run", "-v", "--silence-skips" ]
