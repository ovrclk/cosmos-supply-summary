GO                    := GO111MODULE=on go
GOBIN                 := $(shell go env GOPATH)/bin

UNAME_OS              := $(shell uname -s)

CACHE_BASE            ?= $(abspath .cache)
CACHE                 := $(CACHE_BASE)
CACHE_BIN             := $(CACHE)/bin
CACHE_INCLUDE         := $(CACHE)/include
CACHE_VERSIONS        := $(CACHE)/versions

BUF_VERSION           ?= 0.25.0
PROTOC_VERSION        ?= 3.13.0
GRPC_GATEWAY_VERSION  ?= 1.14.7

# <TOOL>_VERSION_FILE points to the marker file for the installed version.
# If <TOOL>_VERSION_FILE is changed, the binary will be re-downloaded.
PROTOC_VERSION_FILE        = $(CACHE_VERSIONS)/protoc/$(PROTOC_VERSION)
GRPC_GATEWAY_VERSION_FILE  = $(CACHE_VERSIONS)/protoc-gen-grpc-gateway/$(GRPC_GATEWAY_VERSION)
MODVENDOR                  = $(CACHE_BIN)/modvendor
PROTOC                    := $(CACHE_BIN)/protoc
GRPC_GATEWAY              := $(CACHE_BIN)/protoc-gen-grpc-gateway

DOCKER_RUN            := docker run --rm -v $(shell pwd):/workspace -w /workspace
DOCKER_CLANG          := $(DOCKER_RUN) tendermintdev/docker-build-proto

ifeq ($(UNAME_OS),Linux)
	PROTOC_ZIP       ?= protoc-${PROTOC_VERSION}-linux-x86_64.zip
	GRPC_GATEWAY_BIN ?= protoc-gen-grpc-gateway-v${GRPC_GATEWAY_VERSION}-linux-x86_64
endif
ifeq ($(UNAME_OS),Darwin)
	PROTOC_ZIP       ?= protoc-${PROTOC_VERSION}-osx-x86_64.zip
	GRPC_GATEWAY_BIN ?= protoc-gen-grpc-gateway-v${GRPC_GATEWAY_VERSION}-darwin-x86_64
endif

.PHONY: deps-tidy
deps-tidy:
	$(GO) mod tidy

.PHONY: deps-vendor
deps-vendor:
	go mod vendor

.PHONY: modsensure
modsensure: deps-tidy deps-vendor

.PHONY: modvendor
modvendor: modsensure $(MODVENDOR)
	@echo "vendoring non-go files..."
	$(MODVENDOR) -copy="**/*.proto" -include=\
github.com/cosmos/cosmos-sdk/proto,\
github.com/cosmos/cosmos-sdk/third_party/proto

###############################################################################
###                          Setup Cache                                   ###
###############################################################################

$(CACHE):
	@echo "creating .cache dir structure..."
	mkdir -p $@
	mkdir -p $(CACHE_BIN)
	mkdir -p $(CACHE_INCLUDE)
	mkdir -p $(CACHE_VERSIONS)

$(PROTOC_VERSION_FILE): $(CACHE)
	@echo "installing protoc compiler..."
	rm -f $(PROTOC)
	(cd /tmp; \
	curl -sOL "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"; \
	unzip -oq ${PROTOC_ZIP} -d $(CACHE) bin/protoc; \
	unzip -oq ${PROTOC_ZIP} -d $(CACHE) 'include/*'; \
	rm -f ${PROTOC_ZIP})
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC): $(PROTOC_VERSION_FILE)

$(PROTOC_GEN_COSMOS_VERSION_FILE): $(CACHE)
	@echo "installing protoc-gen-cosmos..."
	rm -f $(PROTOC_GEN_COSMOS)
	GOBIN=$(CACHE_BIN) $(GO) install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(PROTOC_GEN_COSMOS): $(PROTOC_GEN_COSMOS_VERSION_FILE)

$(GRPC_GATEWAY_VERSION_FILE): $(CACHE)
	@echo "Installing protoc-gen-grpc-gateway..."
	rm -f $(GRPC_GATEWAY)
	curl -o "${CACHE_BIN}/protoc-gen-grpc-gateway" -L \
	"https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${GRPC_GATEWAY_VERSION}/${GRPC_GATEWAY_BIN}"
	chmod +x "$(CACHE_BIN)/protoc-gen-grpc-gateway"
	rm -rf "$(dir $@)"
	mkdir -p "$(dir $@)"
	touch $@
$(GRPC_GATEWAY): $(GRPC_GATEWAY_VERSION_FILE)

$(MODVENDOR): $(CACHE)
	@echo "installing modvendor..."
	GOBIN=$(CACHE_BIN) GO111MODULE=off go get github.com/goware/modvendor

cache-clean:
	rm -rf $(CACHE)

###############################################################################
###                           Protobuf                                    ###
###############################################################################

.PHONY: proto-gen
proto-gen: $(PROTOC) $(GRPC_GATEWAY) $(PROTOC_GEN_COSMOS) modvendor proto-format
	./script/protocgen.sh

.PHONY: proto-format
proto-format:
	$(DOCKER_CLANG) find ./ ! -path "./vendor/*" -name *.proto -exec clang-format -i {} \;
