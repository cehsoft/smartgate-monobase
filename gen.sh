#!/bin/bash

# Install protoc: brew install protobuf
# Install go gencode: https://grpc.io/docs/languages/go/quickstart/#prerequisites
# Install js gencode: https://github.com/improbable-eng/ts-protoc-gen
protoc proto/service.proto \
    --plugin="protoc-gen-ts=`which protoc-gen-ts`" \
    --js_out="import_style=commonjs,binary:./js/" \
    --ts_out="service=grpc-web:./js/" \
    --go-grpc_out="./services/manager/mygrpc/" \
    --go_out="./services/manager/mygrpc/"


arch -x86_64 python -m grpc_tools.protoc -I. --python_out="./python/service" --grpc_python_out="./python/service" proto/service.proto # MAC M1 chip
arch -x86_64 python -m grpc_tools.protoc -I. --python_out="./python/ocr" --grpc_python_out="./python/ocr" proto/ocr.proto # MAC M1 chip
# python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. proto/service.proto # Intel chip