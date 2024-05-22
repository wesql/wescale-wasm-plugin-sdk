.PHONY: wazero hello

build-scripts:
	mkdir -p bin
	go build -o ./bin/insertPlugin ./tools/scripts/insertPlugin.go

build-wazero:
	mkdir -p bin
	tinygo build --no-debug -o ./bin/myguest.wasm -target=wasi -scheduler=none ./wazero/wazero_main.go

########################################################################################################

clean:
	rm -f ./bin/*

build: build-wazero build-scripts

install: build
	./bin/insertPlugin --wasm_file=./bin/myguest.wasm

