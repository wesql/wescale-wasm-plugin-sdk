.PHONY: wazero hello

build-scripts:
	mkdir -p bin
	go build -o ./bin/wescale_wasm ./tools/scripts/wescale_wasm.go

build-wazero:
	mkdir -p bin
	tinygo build --no-debug -o ./bin/myguest.wasm -target=wasi -scheduler=none ./wazero/wazero_main.go

########################################################################################################

reborn: uninstall clean build install

clean:
	rm -f ./bin/*

build: build-wazero build-scripts

install: build
	./bin/wescale_wasm --command=install --wasm_file=./bin/myguest.wasm

uninstall:
	./bin/wescale_wasm --command=uninstall
