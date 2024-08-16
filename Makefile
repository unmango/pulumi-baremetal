_ := $(shell mkdir -p $(addprefix .make/,examples lint tidy) .test)
PROJECT_NAME := Pulumi baremetal Resource Provider

PACK             := baremetal
PACKDIR          := sdk
PROJECT          := github.com/unmango/pulumi-baremetal
NODE_MODULE_NAME := @unmango/baremetal
NUGET_PKG_NAME   := UnMango.Baremetal
PROVISIONER_NAME := baremetal-provisioner

PROVIDER        := pulumi-resource-${PACK}
VERSION         ?= $(shell pulumictl get version --language generic)
SUPPORTED_SDKS  := dotnet go nodejs python
PROTO_VERSION   := v1alpha1
PROVIDER_PATH   := provider
VERSION_PATH    := ${PROVIDER_PATH}.Version
DOCKER_TAG      ?= $(shell cut -d'.' -f-3 <<<'${VERSION}' | sed 's/+dirty//')

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

PULUMI ?= ${WORKING_DIR}/bin/pulumi/pulumi
GINKGO ?= go run github.com/onsi/ginkgo/v2/ginkgo
DOTNET ?= ${WORKING_DIR}/bin/dotnet/dotnet
NVM    ?= ${WORKING_DIR}/bin/.nvm/nvm.sh
NODE   ?= ${WORKING_DIR}/bin/node

default:: provider provisioner

ensure:: $(GO_MODULES:%=.make/tidy/%)

remake::
	rm -rf bin dist out .make .test hack/.work

provider:: bin/$(PROVIDER)
provisioner:: bin/provisioner
sdks:: $(SUPPORTED_SDKS:%=%_sdk)

provider_debug::
	go -C ${PROVIDER_PATH} build \
		-o $(WORKING_DIR)/bin/${PROVIDER} \
		-gcflags="all=-N -l" \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER)

test_all:: test_provider test_sdks
test_provider:: .test/lifecycle
test_sdks:: .test/sdks

docker:: .make/provisioner_docker_build .make/provisioner_test_docker_build
proto:: gen_proto

gen:: gen_proto gen_sdks gen_examples
gen_proto:: $(GEN_SRC)
gen_sdks:: $(SUPPORTED_SDKS:%=sdk/%)
gen_examples: $(SUPPORTED_SDKS:%=.make/examples/%)

.PHONY: sdk/%
sdk/%: $(SCHEMA_FILE)
	rm -rf $@
	$(PULUMI) package gen-sdk --language $* $(SCHEMA_FILE) --version "${VERSION}"

sdk/python: $(SCHEMA_FILE)
	rm -rf $@
	$(PULUMI) package gen-sdk --language python $(SCHEMA_FILE) --version "${VERSION}"
	cp README.md ${PACKDIR}/python/

dotnet_sdk: sdk/dotnet
	cd ${PACKDIR}/dotnet/ && \
		echo "${VERSION}" >version.txt && \
		$(DOTNET) build

go_sdk: sdk/go

nodejs_sdk: sdk/nodejs
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock bin/

python_sdk: sdk/python
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && \
		../venv/bin/python -m build .

define pulumi_login
    export PULUMI_CONFIG_PASSPHRASE=asdfqwerty1234; \
    $(PULUMI) login --local;
endef

up::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	$(PULUMI) stack init dev && \
	$(PULUMI) stack select dev && \
	$(PULUMI) config set name dev && \
	$(PULUMI) up -y

down::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	$(PULUMI) stack select dev && \
	$(PULUMI) destroy -y && \
	$(PULUMI) stack rm dev -y

devcontainer::
	git submodule update --remote --merge .github/devcontainer
	rsync -av .github/devcontainer/.devcontainer/* .devcontainer

.PHONY: build
build:: provider provisioner dotnet_sdk go_sdk nodejs_sdk python_sdk

lint:: .make/lint/buf .make/lint_go

install:: install_nodejs_sdk install_dotnet_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

install_dotnet_sdk::
	rm -rf $(WORKING_DIR)/nuget/$(NUGET_PKG_NAME).*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;
	if ! dotnet nuget list source | grep ${WORKING_DIR}; then \
		dotnet nuget add source ${WORKING_DIR}/nuget --name ${WORKING_DIR} \
	; fi

install_python_sdk::
	#target intentionally blank

install_go_sdk::
	#target intentionally blank

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

# ------- Real Targets -------
out/install.sh: $(PROVIDER_PATH)/cmd/provisioner/install.sh
	mkdir -p '${@D}' && cp '$<' '$@'

out/baremetal-provisioner.service: $(PROVIDER_PATH)/cmd/provisioner/baremetal-provisioner.service
	mkdir -p '${@D}' && cp '$<' '$@'

# ----------- Tools -----------
bin/install-pulumi.sh: .versions/pulumi
	curl -L https://get.pulumi.com -o $@ && chmod +x $@
bin/pulumi: .versions/pulumi bin/install-pulumi.sh
	bin/install-pulumi.sh \
		--version $(shell cat $<) \
		--install-root ${WORKING_DIR}/$@ \
		--no-edit-path
	cd $@ && mv bin/* . && rm -r bin

bin/dotnet-install.sh: .versions/dotnet
	curl -L https://dotnet.microsoft.com/download/dotnet/scripts/v1/dotnet-install.sh -o $@ && chmod +x $@
bin/dotnet: .versions/dotnet bin/dotnet-install.sh
	bin/dotnet-install.sh \
		--channel $(shell cat $<) \
		--install-dir ${WORKING_DIR}/$@ \
		--verbose

# What a headache this was https://github.com/nvm-sh/nvm/issues/1985
export NVM_DIR := ${WORKING_DIR}/bin/.nvm
bin/install-nvm.sh: .versions/nvm
	curl https://raw.githubusercontent.com/nvm-sh/nvm/v$(shell cat $<)/install.sh -o $@ && chmod +x $@
bin/.nvm: .versions/nvm bin/install-nvm.sh
	mkdir -p $@ && PROFILE=/dev/null bin/install-nvm.sh --no-use
bin/.nvm/versions/node/v$(shell cat .nvmrc): bin/.nvm
	. ${NVM_DIR}/nvm.sh --no-use && nvm install && nvm use
.make/bin_node: bin/.nvm/versions/node/v$(shell cat .nvmrc)
	ln -sf ${WORKING_DIR}/$</bin/node ${WORKING_DIR}/bin
	@touch $@
.PHONY: bin/node
bin/node: .make/bin_node

bin/$(PROVIDER):: $(GEN_SRC) $(PKG_SRC) provider/*go*
	go -C provider build \
		-o $(WORKING_DIR)/$@ \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER)

bin/provisioner:: $(GEN_SRC) provider/cmd/provisioner/*.go $(PKG_SRC)
	go -C provider build \
		-o ${WORKING_DIR}/$@ \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/provisioner

$(GEN_SRC) &: $(PROTO_SRC) $(BUF_CONFIG)
	buf generate $(patsubst %,--path %,$(filter %.proto,$?))

buf.lock: $(BUF_CONFIG)
	buf dep update

.make/tidy/gen: $(filter gen/%,$(GO_SRC))
.make/tidy/provider: $(filter provider/%,$(GO_SRC))
.make/tidy/sdk: $(filter sdk/%,$(GO_SRC))
.make/tidy/tests: $(filter tests/%,$(GO_SRC))
$(GO_MODULES:%=.make/tidy/%): .make/tidy/%: $(addprefix %/,go.mod go.sum)
	go -C $* mod tidy
	@touch $@

.make/lint_go: $(patsubst %,.make/lint/%,provider sdk tests)
.make/lint/provider: $(PROVIDER_SRC)
.make/lint/tests: $(shell find tests -name '*.go')
# .make/lint/sdk: $(shell find sdk/go -name '*.go')
.make/lint/%:
	cd $* && golangci-lint run -c ${WORKING_DIR}/.golangci.yml --timeout 1m ./...
	@touch $@

.make/buf_build: buf.lock $(PROTO_SRC)
	buf build --path $(filter %.proto,$?)
	@touch $@

.make/lint/buf: $(PROTO_SRC)
	buf lint $(?:%=--path %)
	@touch $@

.make/provisioner_docker: provider/cmd/provisioner/Dockerfile .dockerignore $(PROVIDER_SRC)
	docker build ${WORKING_DIR} -f $< -t ${PROVISIONER_NAME}:${DOCKER_TAG}
	@touch $@

.make/provisioner_docker_test: provider/cmd/provisioner/Dockerfile .dockerignore $(PROVIDER_SRC)
	docker build ${WORKING_DIR} -f $< --target test -t ${PROVISIONER_NAME}:test
	@touch $@

.make/sdk_docker: tests/sdk/Dockerfile .dockerignore $(PROVIDER_SRC) bin/$(PROVIDER)
	docker build ${WORKING_DIR} -f $< -t sdk-test:dotnet
	@touch $@
.test/sdk_docker: .make/sdk_docker
	docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock sdk-test:dotnet

.make/examples/%: examples/yaml/** bin/$(PROVIDER)
	rm -rf ${WORKING_DIR}/examples/$*
	$(PULUMI) convert \
		--cwd $(<D) \
		--logtostderr \
		--generate-only \
		--non-interactive \
		--language $* \
		--out ${WORKING_DIR}/examples/$*
	@touch $@

export GRPC_GO_LOG_SEVERITY_LEVEL ?=
TEST_FLAGS ?=

.test/lifecycle: .make/provisioner_docker_test
	cd tests/lifecycle && $(GINKGO) run -v --silence-skips ${TEST_FLAGS}

.test/pkg: $(PKG_SRC)
	cd provider && $(GINKGO) run -v -r

export PULUMI_LOCAL_NUGET := ${WORKING_DIR}/nuget

.test/sdks: $(SUPPORTED_SDKS:%=.test/sdk_%)
.test/sdk_dotnet: install_dotnet_sdk bin/dotnet
$(SUPPORTED_SDKS:%=.test/sdk_%): .test/sdk_%:
	cd tests/sdk && $(GINKGO) run -v --silence-skips ${TEST_FLAGS}
	@touch $@

.test/install_script: out/install.sh $(PROVIDER_PATH)/cmd/provisioner/baremetal-provisioner.service Makefile
	DEV_MODE=true INSTALL_DIR=${WORKING_DIR}/bin $<
	@touch $@

.envrc: hack/.envrc.example
	cp $< $@
