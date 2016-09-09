#!/usr/bin/env sh
protoc add.proto \
    --go_out=plugins=grpc:. \
    --python_out=plugins=grpc:. \
    --ruby_out=plugins=grpc:. \
    --js_out=. \
