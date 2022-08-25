#!/bin/bash

# set current working directory to script location
cd "${0%/*}" || return

# try to delete directory, ignore error if directory doesn't exist
rm -rf ../proto ||:
mkdir ../proto
# shellcheck disable=SC2046
protoc --go_out=../proto \
    --go_opt=paths=source_relative \
    --go-grpc_out=../proto \
    --go-grpc_opt=paths=source_relative \
    --proto_path=../kitsune/core/src/main/proto \
    $(find ../kitsune/core/src/main/proto -iname "*.proto")