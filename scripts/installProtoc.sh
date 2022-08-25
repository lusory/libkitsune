#!/bin/bash

# set current working directory to script location
cd "${0%/*}" || return

PROTOC_VERSION=21.4
curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip"
# shellcheck disable=SC2086
unzip protoc-$PROTOC_VERSION-linux-x86_64.zip -d $HOME/.local
rm protoc-$PROTOC_VERSION-linux-x86_64.zip
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest