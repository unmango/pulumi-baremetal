_ := $(shell mkdir -p $(addprefix .make/,examples lint tidy gen build test install docker))
PROJECT_NAME := Pulumi baremetal Resource Provider

PACK             := baremetal
PACKDIR          := sdk
PROJECT          := github.com/unmango/pulumi-baremetal
NODE_MODULE_NAME := @unmango/baremetal
NUGET_PKG_NAME   := UnMango.Baremetal
PROVISIONER_NAME := baremetal-provisioner

PROVIDER        := pulumi-resource-${PACK}
SUPPORTED_SDKS  := dotnet go nodejs python
PROTO_VERSION   := v1alpha1
PROVIDER_PATH   := provider

PROVIDER_VERSION ?= 1.0.0-alpha.0+dev
VERSION_GENERIC  = $(shell pulumictl convert-version --language generic --version "$(PROVIDER_VERSION)")
VERSION_PATH     := ${PROVIDER_PATH}.Version
DOCKER_TAG       ?= $(shell echo '${VERSION_GENERIC}' | cut -d'.' -f-3 | sed 's/+dirty//')

GOPATH			:= $(shell go env GOPATH)

WORKING_DIR  := $(shell pwd)
EXAMPLES_DIR := ${WORKING_DIR}/examples/yaml
PROTO_PKG    := unmango/baremetal/${PROTO_VERSION}
PROTO_DIR    := proto/${PROTO_PKG}
PKG_DIR      := ${PROVIDER_PATH}/pkg

# The schema file is currently embedded in the provider binary
SCHEMA_FILE := bin/${PROVIDER}

TESTPARALLELISM := 4
OS := $(shell uname)

BUF_CONFIG := buf.yaml buf.gen.yaml

GO_MODULES   := gen provider sdk tests
GO_SRC       := $(subst ./,,$(shell find . -type f -name '*.go'))
PROVIDER_SRC := $(filter $(PROVIDER_PATH)/%,$(GO_SRC))
PKG_SRC      := $(filter $(PKG_DIR)/%,$(GO_SRC))
PROTO_SRC    := $(shell find $(PROTO_DIR) -type f -name '*.proto')
GO_GRPC_SRC  := $(PROTO_SRC:proto/%.proto=gen/go/%_grpc.pb.go)
GO_PB_SRC    := $(PROTO_SRC:proto/%.proto=gen/go/%.pb.go)
GEN_SRC      := $(GO_GRPC_SRC) $(GO_PB_SRC)

GINKGO ?= go run github.com/onsi/ginkgo/v2/ginkgo

export PULUMI_LOCAL_NUGET := ${WORKING_DIR}/nuget

default:: provider provisioner
provider:: bin/$(PROVIDER)
provisioner:: bin/provisioner

tidy: $(GO_MODULES:%=.make/tidy/%)
lint:: .make/lint/buf $(GO_MODULES:%=.make/lint/%)

remake::
	rm -rf .make bin dist out hack/.work

test_all:: test_provider test_sdks test_pkg .make/test/install_script
test_provider:: .make/test/lifecycle
test_sdks:: $(SUPPORTED_SDKS:%=.make/test/%_sdk)
test_pkg:: .make/test/pkg

docker:: \
	.make/docker/provisioner \
	.make/docker/provisioner_test \
	.make/docker/provider

gen:: gen_proto gen_sdks gen_examples
gen_proto:: $(GEN_SRC)
gen_sdks:: $(SUPPORTED_SDKS:%=.make/gen/%)
gen_examples:: $(SUPPORTED_SDKS:%=.make/examples/%)

build:: provider provisioner $(SUPPORTED_SDKS:%=.make/build/%)
dotnet_sdk:: .make/build/dotnet
nodejs_sdk:: .make/build/nodejs
go_sdk:: .make/build/go
python_sdk:: .make/build/python

install:: install_nodejs_sdk install_dotnet_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

define pulumi_login
    export PULUMI_CONFIG_PASSPHRASE=asdfqwerty1234; \
    pulumi login --local;
endef

up::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	pulumi stack init dev && \
	pulumi stack select dev && \
	pulumi config set name dev && \
	pulumi up -y

down::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	pulumi stack select dev && \
	pulumi destroy -y && \
	pulumi stack rm dev -y

provider_debug::
	go -C ${PROVIDER_PATH} build \
		-o $(WORKING_DIR)/bin/${PROVIDER} \
		-gcflags="all=-N -l" \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION_GENERIC}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER)

# ------- Real Targets -------
out/install.sh: $(PROVIDER_PATH)/cmd/provisioner/install.sh
	mkdir -p '${@D}' && cp '$<' '$@'

out/baremetal-provisioner.service: $(PROVIDER_PATH)/cmd/provisioner/baremetal-provisioner.service
	mkdir -p '${@D}' && cp '$<' '$@'

bin/$(PROVIDER): $(GEN_SRC) $(PKG_SRC) provider/*go*
	go -C provider build \
		-o $(WORKING_DIR)/$@ \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION_GENERIC}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER)

bin/provisioner:: $(GEN_SRC) provider/cmd/provisioner/*.go $(PKG_SRC)
	go -C provider build \
		-o ${WORKING_DIR}/$@ \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION_GENERIC}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/provisioner

gen/go/%_grpc.pb.go gen/go/%.pb.go &:: proto/%.proto $(BUF_CONFIG)
	buf generate $(patsubst %,--path %,$(filter %.proto,$?))

buf.lock: $(BUF_CONFIG)
	buf dep update

.envrc: hack/example.envrc
	cp $< $@

# ------ Sentinal Targets ------
.make/tidy/gen: $(filter gen/%,$(GO_SRC))
.make/tidy/provider: $(filter provider/%,$(GO_SRC))
.make/tidy/sdk: $(filter sdk/%,$(GO_SRC))
.make/tidy/tests: $(filter tests/%,$(GO_SRC))
$(GO_MODULES:%=.make/tidy/%): .make/tidy/%: $(addprefix %/,go.mod go.sum)
	go -C $* mod tidy
	@touch $@

.make/lint/provider: $(PROVIDER_SRC)
.make/lint/tests: $(shell find tests -name '*.go')
# .make/lint/sdk: $(shell find sdk/go -name '*.go')
$(GO_MODULES:%=.make/lint/%): .make/lint/%:
	cd $* && golangci-lint run -c ${WORKING_DIR}/.golangci.yml --timeout 1m ./...
	@touch $@

.make/lint/buf: $(PROTO_SRC)
	buf lint $(?:%=--path %)
	@touch $@

# -------- SDKs --------
.make/gen/%: $(SCHEMA_FILE)
	rm -rf $@
	pulumi package gen-sdk --language $* $(SCHEMA_FILE) --version "${VERSION_GENERIC}"
	@touch $@
.make/gen/python: $(SCHEMA_FILE)
	rm -rf $@
	pulumi package gen-sdk --language python $(SCHEMA_FILE) --version "${VERSION_GENERIC}"
	cp README.md ${PACKDIR}/python/
	@touch $@

.make/build/dotnet: .make/gen/dotnet
	cd ${PACKDIR}/dotnet/ && \
		echo "${VERSION_GENERIC}" >version.txt && \
		dotnet build
	@touch $@
.make/build/go: .make/gen/go
.make/build/nodejs: .make/gen/nodejs
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc
	cp README.md LICENSE ${PACKDIR}/nodejs/package.json ${PACKDIR}/nodejs/yarn.lock ${PACKDIR}/nodejs/bin/
	@touch $@
.make/build/python: .make/gen/python
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && \
		../venv/bin/python -m build .
	@touch $@

.make/install/dotnet:
	rm -rf $(WORKING_DIR)/nuget/$(NUGET_PKG_NAME).*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;
	if ! dotnet nuget list source | grep ${WORKING_DIR}; then \
		dotnet nuget add source ${WORKING_DIR}/nuget --name ${WORKING_DIR} \
	; fi
	@touch $@
.make/install/nodejs:
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin
	@touch $@

# ------- Protobuf -------
.make/buf_build: buf.lock $(PROTO_SRC)
	buf build --path $(filter %.proto,$?)
	@touch $@

# -------- Docker --------
.make/docker/provisioner: provider/cmd/provisioner/Dockerfile .dockerignore $(PROVIDER_SRC)
	docker build ${WORKING_DIR} -f $< -t ${PROVISIONER_NAME}:${DOCKER_TAG} --build-arg VERSION=${VERSION_GENERIC}
	@touch $@
.make/docker/provisioner_test: provider/cmd/provisioner/Dockerfile .dockerignore $(PROVIDER_SRC)
	docker build ${WORKING_DIR} -f $< --target test -t ${PROVISIONER_NAME}:test --build-arg VERSION=${VERSION_GENERIC}
	@touch $@
.make/docker/provider: provider/cmd/$(PROVIDER)/Dockerfile .dockerignore $(PROVIDER_SRC)
	docker build ${WORKING_DIR} -f $< --target bin -t ${PROVIDER}:${DOCKER_TAG} --build-arg VERSION=${VERSION_GENERIC}
	@touch $@
.make/docker/%_build: compose.yml tests/sdk/Dockerfile provider/cmd/provisioner/Dockerfile .dockerignore $(GO_SRC) bin/$(PROVIDER)
	VERSION=${VERSION_GENERIC} docker compose build $*-test
	@touch $@

# ------- Examples -------
.make/examples/%: examples/yaml/** bin/$(PROVIDER)
	rm -rf ${WORKING_DIR}/examples/$*
	pulumi convert \
		--cwd $(<D) \
		--logtostderr \
		--generate-only \
		--non-interactive \
		--language $* \
		--out ${WORKING_DIR}/examples/$*
	@touch $@

# ------- Tests -------
export GRPC_GO_LOG_SEVERITY_LEVEL ?=
TEST_FLAGS ?=

.make/test/%_sdk: .make/docker/%_build
	VERSION=${VERSION_GENERIC} docker compose up provisioner-test $*-test --exit-code-from $*-test
	VERSION=${VERSION_GENERIC} docker compose down

.make/test/lifecycle: .make/docker/provisioner_test
	cd tests/lifecycle && $(GINKGO) run -v --silence-skips ${TEST_FLAGS}

.make/test/pkg: $(PKG_SRC)
	cd provider && $(GINKGO) run -v -r

# .make/test/dotnet_sdk: .make/install/dotnet
# $(SUPPORTED_SDKS:%=.make/test/%_sdk): .make/test/%_sdk:
# 	cd tests/sdk && $(GINKGO) run -v --silence-skips ${TEST_FLAGS}
# 	@touch $@

.make/test/install_script: out/install.sh $(PROVIDER_PATH)/cmd/provisioner/baremetal-provisioner.service Makefile
	DEV_MODE=true INSTALL_DIR=${WORKING_DIR}/bin $<
	@touch $@
