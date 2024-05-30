ifndef WESCALEROOT
export WESCALEROOT=${PWD}
endif

export WESCALEROOTBIN=${WESCALEROOT}/bin

minimaltools:
	echo $$(date): Installing minimal dependencies
	BUILD_CHROME=0 BUILD_JAVA=0 BUILD_CONSUL=0 ./tools/bootstrap.sh

install_protoc-gen-go:
	GOBIN=$(WESCALEROOT)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@$(shell go list -m -f '{{ .Version }}' google.golang.org/protobuf)
	GOBIN=$(WESCALEROOT)/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 # the GRPC compiler its own pinned version
	GOBIN=$(WESCALEROOT)/bin go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@$(shell go list -m -f '{{ .Version }}' github.com/planetscale/vtprotobuf)

PROTO_SRCS = $(wildcard proto/*.proto)
PROTO_SRC_NAMES = $(basename $(notdir $(PROTO_SRCS)))
PROTO_GO_OUTS = $(foreach name, $(PROTO_SRC_NAMES), pkg/proto/$(name)/$(name).pb.go)
# This rule rebuilds all the go files from the proto definitions for gRPC.
proto: $(PROTO_GO_OUTS)

$(PROTO_GO_OUTS): minimaltools install_protoc-gen-go proto/*.proto
	$(WESCALEROOT)/bin/protoc \
		--go_out=. --plugin protoc-gen-go="${WESCALEROOTBIN}/protoc-gen-go" \
		--go-grpc_out=. --plugin protoc-gen-go-grpc="${WESCALEROOTBIN}/protoc-gen-go-grpc" \
		--go-vtproto_out=. --plugin protoc-gen-go-vtproto="${WESCALEROOTBIN}/protoc-gen-go-vtproto" \
		--go-vtproto_opt=features=marshal+unmarshal+size+pool \
		--go-vtproto_opt=pool=pkg/proto/query/query.Row \
		-I${PWD}/dist/vt-protoc-21.3/include:proto $(PROTO_SRCS)


########################################################################################################

build-tools:
	mkdir -p bin
	go build -o ./bin/wescale_wasm ./cmd/wescale_wasm/main.go

VERSION := 'v0.1.2-beta3'
PLATFORMS := darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64 windows/386 windows/amd64
SOURCE_DIR := ./cmd/wescale_wasm
BINARY_NAME := wescale_wasm
TARGET_DIR := ./bin

cross-build-tools:
	@for platform in $(PLATFORMS); do \
		platform_split=($${platform//\// }); \
		GOOS=$${platform_split[0]}; \
		GOARCH=$${platform_split[1]}; \
		output_name=$(BINARY_NAME)_$(VERSION)_$${GOOS}_$${GOARCH}; \
		if [ $$GOOS = "windows" ]; then \
			output_name+='.exe'; \
		fi; \
		output_path=$(TARGET_DIR)/$${output_name}; \
		echo "Building $(VERSION) for $$GOOS/$$GOARCH..."; \
		env GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$output_path $(SOURCE_DIR); \
		if [ $$? -ne 0 ]; then \
			echo "An error has occurred! Aborting the script execution..."; \
			exit 1; \
		fi; \
	done

########################################################################################################

build-examples:
	# Iterate over all the examples and build them
	for example in $(shell ls ./examples); do \
		echo "Building example: $$example"; \
		tinygo build --no-debug -o ./bin/$$example.wasm -target=wasi -scheduler=none ./examples/$$example/main.go; \
	done

########################################################################################################

reborn: clean build uninstall install

clean:
	rm -f ./bin/*
	rm -rf ./dist/*

build: clean build-tools build-examples

install:
	./bin/wescale_wasm --command=install --wasm_file=./bin/datamasking.wasm

uninstall:
	./bin/wescale_wasm --command=uninstall




