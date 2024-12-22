#!/usr/bin/env bash

# Install protoc
if [ ! -f ~/bin/protoc ]; then
    cd ~/
    wget https://github.com/protocolbuffers/protobuf/releases/download/v29.2/protoc-29.2-linux-x86_64.zip
    unzip protoc-29.2-linux-x86_64.zip bin/protoc
    rm protoc-29.2-linux-x86_64.zip
fi

# Install protoc-gen-go and protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
