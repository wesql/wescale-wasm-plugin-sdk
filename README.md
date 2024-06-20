# wescale-wasm-plugin-sdk

This repository contains:
* the SDK for writing WeScale Wasm plugins.
* the `wescale_wasm` binary to deploy the plugin.

## How to release a new version of the SDK
You can use `./release_new_version.sh` to help you.

The SDK is released as a Go module. To release a new version of the SDK, you need to do the following:

1. Make sure you have the latest changes in the `main` branch.

2. Tag the latest commit with the new version number. For example, if the new version is `v1.0.0`, you can tag the latest commit with `v1.0.0` and push the tag to the remote repository. You can do this by running the following commands:
```bash
git tag v1.0.0
git push origin v1.0.0
```

3. Wait for the go module to be updated. You can see all the available versions of the module at the following link: https://proxy.golang.org/github.com/wesql/wescale-wasm-plugin-sdk/@v/list


## How to release a new version of the `wescale_wasm` binary
You can use `./release_new_version.sh` to help you.

If you've made changes to the `cmd/wescale_wasm` directory, you need to release a new version of the `wescale_wasm` binary. You can do this by following the steps below:

1. Make sure you have the latest changes in the `main` branch.
2. Change the version number in the `Makefile` file. You need to commit and push this change to the remote repository.
3. Use `make build` command to build the `wescale_wasm` binary.
4. Tag the latest commit with the new version number. For example, if the new version is `v1.0.0`, you can tag the latest commit with `v1.0.0` and push the tag to the remote repository. You can do this by running the following commands:
```bash
git tag v1.0.0
git push origin v1.0.0
```
5. Draft a new release on [GitHub](https://github.com/wesql/wescale-wasm-plugin-sdk/releases) with the new version number and attach the `wescale_wasm` binary to the release.
6. Change the version number in the [wescale-wasm-plugin-template](https://github.com/wesql/wescale-wasm-plugin-template/blob/main/Makefile) project. You need to commit and push this change to the remote repository.