#!/bin/bash

# Prompt user to input version number
read -p "Please enter the version number (format: vx.x.x, where x represents a digit): " version

# Validate the version number format
if [[ ! $version =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Invalid version format, the correct format is vx.x.x, where x represents a digit"
    exit 1
fi

# Switch to the main branch and pull the latest main branch from remote
git checkout main
if [ $? -ne 0 ]; then
    echo "Failed to switch to the main branch"
    exit 1
fi

git pull origin main
if [ $? -ne 0 ]; then
    echo "Failed to pull the latest main branch from remote"
    exit 1
fi

# Create and switch to a new branch
new_branch="branch_$version"
git checkout -b $new_branch
if [ $? -ne 0 ]; then
    echo "Failed to create and switch to the new branch"
    exit 1
fi

# Create git tag and push to remote
git tag $version
if [ $? -ne 0 ]; then
    echo "Failed to create the tag"
    exit 1
fi

git push origin $version
if [ $? -ne 0 ]; then
    echo "Failed to push the tag to remote"
    exit 1
fi

# Set environment variable and run make build
export WASM_PLUGIN_SDK_VERSION=$version
make build
if [ $? -ne 0 ]; then
    echo "Failed to execute make build"
    exit 1
fi

echo "tag $version has been uploaded to origin successfully, please release a new version in https://github.com/wesql/wescale-wasm-plugin-sdk/releases and update sdk version in https://github.com/wesql/wescale-wasm-plugin-template"
