.PHONY: wazero hello

build-scripts:
	mkdir -p bin
	go build -o ./bin/insertPlugin ./tools/scripts/insertPlugin.go

build-wazero:
	mkdir -p bin
	tinygo build --no-debug -o ./bin/myguest.wasm -target=wasi -scheduler=none ./wazero/wazero_main.go

########################################################################################################

mega: uninstall clean build install

clean:
	rm -f ./bin/*

build: build-wazero build-scripts

install: build
	./bin/insertPlugin --wasm_file=./bin/myguest.wasm

uninstall:
	mysql -h127.0.0.1 -P15306 -e 'drop filter wasm'
	mysql -h127.0.0.1 -P15306 -e 'delete from mysql.wasm_binary'