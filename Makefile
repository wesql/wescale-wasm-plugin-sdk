ifndef WESCALEROOT
export WESCALEROOT=${PWD}
endif

export WESCALEROOTBIN=${WESCALEROOT}/bin
export WASM_PLUGIN_SDK_VERSION=v0.1.12

########################################################################################################

PLATFORMS := darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64 windows/386 windows/amd64
SOURCE_DIR := ./cmd/wescale_wasm
BINARY_NAME := wescale_wasm
TARGET_DIR := ./bin

build:
	@if [ -z "$$WASM_PLUGIN_SDK_VERSION" ]; then \
		echo "Error: WASM_PLUGIN_SDK_VERSION environment variable is not defined"; \
		exit 1; \
	fi
	@for platform in $(PLATFORMS); do \
		platform_split=($${platform//\// }); \
		GOOS=$${platform_split[0]}; \
		GOARCH=$${platform_split[1]}; \
		output_name=$(BINARY_NAME)_$${WASM_PLUGIN_SDK_VERSION}_$${GOOS}_$${GOARCH}; \
		if [ $$GOOS = "windows" ]; then \
			output_name+=".exe"; \
		fi; \
		output_path=$(TARGET_DIR)/$${output_name}; \
		echo "Building $${WASM_PLUGIN_SDK_VERSION} for $$GOOS/$$GOARCH..."; \
		env GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$output_path $(SOURCE_DIR); \
		if [ $$? -ne 0 ]; then \
			echo "An error has occurred! Aborting the script execution..."; \
			exit 1; \
		fi; \
	done


clean:
	rm -f ./bin/*


