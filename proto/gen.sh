#!/bin/bash

# Install protoc: brew install protobuf
# Install go gencode: https://grpc.io/docs/languages/go/quickstart/#prerequisites
# Install js gencode: https://github.com/improbable-eng/ts-protoc-gen

protoc service.proto \
    --plugin="protoc-gen-ts=`which protoc-gen-ts`" \
    --js_out="import_style=commonjs,binary:./web/" \
    --ts_out="service=grpc-web:./web/" \
    --go-grpc_out="./go/" \
    --go_out="./go/"