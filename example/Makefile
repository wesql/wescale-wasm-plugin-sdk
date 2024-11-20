build:
	tinygo build --no-debug -o example.wasm -target=wasi -scheduler=none ./main.go

clean:
	rm -f example.wasm

# download wescale_wasm binary at: https://github.com/wesql/wescale-wasm-plugin-sdk/releases/
deploy:
	../bin/wescale_wasm --command=install --wasm_file=./example.wasm --mysql_host=127.0.0.1 --mysql_port=15306 --mysql_user=root --mysql_password=root --create_filter

undeploy:
	../bin/wescale_wasm --command=uninstall --filter_name=example_wasm_filter