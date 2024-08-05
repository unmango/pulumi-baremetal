_ := $(shell mkdir -p .make/{examples,lint,tidy} .test)
PROJECT_NAME := Pulumi baremetal Resource Provider

PACK             := baremetal
PACKDIR          := sdk
PROJECT          := github.com/unmango/pulumi-baremetal
NODE_MODULE_NAME := @unmango/baremetal
NUGET_PKG_NAME   := UnMango.Baremetal
PROVISIONER_NAME := baremetal-provisioner

PROVIDER        := pulumi-resource-${PACK}
VERSION         ?= $(shell pulumictl get version)
SUPPORTED_SDKS  := dotnet go nodejs python
PROTO_VERSION   := v1alpha1
PROVIDER_PATH   := provider
VERSION_PATH    := ${PROVIDER_PATH}.Version
VERSION_TAG     ?= $(shell cut -d'.' -f-3 <<<'${VERSION}' | sed 's/+dirty//')

GOPATH			:= $(shell go env GOPATH)

WORKING_DIR  := $(shell pwd)
EXAMPLES_DIR := ${WORKING_DIR}/examples/yaml
PROTO_PKG    := unmango/baremetal/${PROTO_VERSION}
PROTO_DIR    := proto/${PROTO_PKG}
PKG_DIR      := ${PROVIDER_PATH}/pkg

TESTPARALLELISM := 4
OS := $(shell uname)

BUF_CONFIG := buf.yaml buf.gen.yaml

MANS    := tee
MAN_SRC := $(MANS:%=$(PKG_DIR)/provider/cmd/%.man)

GO_MODULES   := gen provider sdk tests
GO_SRC       := $(subst ./,,$(shell find . -type f -name '*.go'))
PROVIDER_SRC := $(filter $(PROVIDER_PATH)/%,$(GO_SRC))
PKG_SRC      := $(filter $(PKG_DIR)/%,$(GO_SRC))
PROTO_SRC    := $(shell find $(PROTO_DIR) -type f -name '*.proto')
GO_GRPC_SRC  := $(PROTO_SRC:proto/%.proto=gen/go/%_grpc.pb.go)
GO_PB_SRC    := $(PROTO_SRC:proto/%.proto=gen/go/%.pb.go)
GEN_SRC      := $(GO_GRPC_SRC) $(GO_PB_SRC)

GINKGO ?= go run github.com/onsi/ginkgo/v2/ginkgo

default:: provider provisioner

ensure:: $(GO_MODULES:%=.make/tidy/%)

remake::
	rm -rf bin dist out .make .test hack/.work

provider:: bin/$(PROVIDER)
provisioner:: bin/provisioner

provider_debug::
	go -C ${PROVIDER_PATH} build \
		-o $(WORKING_DIR)/bin/${PROVIDER} \
		-gcflags="all=-N -l" \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER)

test_all:: test_provider test_sdks
test_provider:: .test/provider
test_sdks:: .test/sdks

docker:: .make/provisioner_docker_build
mans:: gen_mans
proto:: gen_proto

gen:: gen_proto gen_mans gen_sdks gen_examples
gen_proto:: $(GEN_SRC)
gen_mans:: $(MAN_SRC)
gen_sdks:: $(SUPPORTED_SDKS:%=%_sdk)
gen_examples: $(SUPPORTED_SDKS:%=.make/examples/%)

dotnet_sdk:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
dotnet_sdk::
	rm -rf sdk/dotnet
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language dotnet
	cd ${PACKDIR}/dotnet/&& \
		echo "${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

go_sdk::
	rm -rf sdk/go
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language go

nodejs_sdk:: VERSION := $(shell pulumictl get version --language javascript)
nodejs_sdk::
	rm -rf sdk/nodejs
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language nodejs
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock bin/ && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' bin/package.json && \
		rm ./bin/package.json.bak

python_sdk:: PYPI_VERSION := $(shell pulumictl get version --language python)
python_sdk::
	rm -rf sdk/python
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language python
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist

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

devcontainer::
	git submodule update --remote --merge .github/devcontainer
	rsync -av .github/devcontainer/.devcontainer/* .devcontainer

.PHONY: build
build:: provider provisioner dotnet_sdk go_sdk nodejs_sdk python_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

lint:: .make/lint/buf .make/lint_go

install:: install_nodejs_sdk install_dotnet_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

install_dotnet_sdk::
	rm -rf $(WORKING_DIR)/nuget/$(NUGET_PKG_NAME).*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

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

bin/$(PROVIDER):: $(GEN_SRC) $(MAN_SRC) $(PKG_SRC) provider/*go*
	go -C provider build \
		-o $(WORKING_DIR)/$@ \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER)

bin/provisioner:: $(GEN_SRC) provider/cmd/provisioner/*.go $(PKG_SRC)
	go -C provider build \
		-o ${WORKING_DIR}/$@ \
		-ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" \
		$(PROJECT)/${PROVIDER_PATH}/cmd/provisioner

gen/go/%.pb.go gen/go/%_grpc.pb.go &: $(BUF_CONFIG) proto/%.proto
	buf generate --clean --path proto/$*.proto

provider/pkg/%.man: provider/pkg/%.go
	@# man $(notdir $*) > $@
	@# This is a terrible idea when the output depends on the terminal size
	@touch $@

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
.make/lint/sdk: $(shell find sdk/go -name '*.go')
.make/lint/%:
	cd $* && golangci-lint run -c ${WORKING_DIR}/.golangci.yml --timeout 1m ./...
	@touch $@

.make/buf_build: buf.lock $(PROTO_SRC)
	buf build --path $(filter %.proto,$?)
	@touch $@

.make/lint/buf: $(PROTO_SRC)
	buf lint --path $?
	@touch $@

.make/provisioner_docker_build: provider/cmd/provisioner/Dockerfile bin/provisioner
	docker build ${WORKING_DIR} -f $< -t ${PROVISIONER_NAME}:local-${VERSION_TAG}
	@touch $@

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

TEST_FLAGS ?=

.test/provider: provisioner .make/provisioner_docker_build
	cd tests && $(GINKGO) run -v --silence-skips ${TEST_FLAGS}

.test/pkg: $(PKG_SRC)
	cd provider && $(GINKGO) run -v -r

.test/sdks: $(SUPPORTED_SDKS:%=.test/sdk_%)
.test/sdk_%: tests/sdk/%/*.go sdk/%/**
	go -C tests/sdk/$* test -v -count=1 -cover -timeout 15m ./...
	@touch $@

.test/install_script: out/install.sh $(PROVIDER_PATH)/cmd/provisioner/baremetal-provisioner.service Makefile
	DEV_MODE=true INSTALL_DIR=${WORKING_DIR}/bin $<
	@touch $@

.envrc: hack/.envrc.example
	cp $< $@
