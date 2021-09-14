#!/bin/bash

protoc --go_out=plugins=grpc:/home/anand/Work/golang/src/robotix-agent rcc.proto
protoc --go_out=plugins=grpc:/home/anand/Work/golang/src/robotix-command-center rcc.proto
python -m grpc_tools.protoc -I. --python_out=/home/anand/Work/golang/src/robotix-agent/ros_node/protos --grpc_python_out=/home/anand/Work/golang/src/robotix-agent/ros_node/protos rcc.proto
cp rcc.proto /home/anand/Work/golang/src/delivery-robot-backend/protos/rcc.proto

# Path to this plugin
PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"

# Directory to write generated code to (.js and .d.ts files)
OUT_DIR="/home/anand/Work/golang/src/hotel-robot-backend/src/protos"


PROTO_DEST="/home/anand/Work/golang/src/hotel-robot-backend/src/protos"

# JavaScript code generation
yarn run grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:${PROTO_DEST} \
    --grpc_out=${PROTO_DEST} \
    --plugin=protoc-gen-grpc="./node_modules/.bin/grpc_tools_node_protoc_plugin" \
    -I. \
    ./rcc.proto

# TypeScript code generation
yarn run grpc_tools_node_protoc \
    --plugin=protoc-gen-ts="./node_modules/.bin/protoc-gen-ts" \
    --ts_out=${PROTO_DEST} \
    -I. \
    ./rcc.proto