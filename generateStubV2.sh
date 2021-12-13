#!/bin/bash

protoc --go_out=plugins=grpc:/home/anand/Work/golang/src/robotix-agent rcc.proto
protoc --go_out=plugins=grpc:/home/anand/Work/golang/src/robotix-command-center rcc.proto
python -m grpc_tools.protoc -I. --python_out=/home/anand/Work/golang/src/robotix-agent/ros_node/protos --grpc_python_out=/home/anand/Work/golang/src/robotix-agent/ros_node/protos rcc.proto
# cp rcc.proto /home/anand/Work/golang/src/delivery-robot-backend/protos/rcc.proto

PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
GRPC_TOOLS_NODE_PROTOC_PLUGIN="./node_modules/.bin/grpc_tools_node_protoc_plugin"
GRPC_TOOLS_NODE_PROTOC="./node_modules/.bin/grpc_tools_node_protoc"
f="./"
# f="./protos"
# for f in ./rcc.proto; do

#   # skip the non proto files
#   if [ "$(basename "$f")" == "index.ts" ]; then
#       continue
#   fi

  # loop over all the available proto files and compile them into respective dir
  # JavaScript code generating
  ${GRPC_TOOLS_NODE_PROTOC} \
      --js_out=import_style=commonjs,binary:"${f}" \
      --grpc_out="${f}" \
      --plugin=protoc-gen-grpc="${GRPC_TOOLS_NODE_PROTOC_PLUGIN}" \
      -I . \
      "./rcc.proto"

  ${GRPC_TOOLS_NODE_PROTOC} \
      --plugin=protoc-gen-ts="${PROTOC_GEN_TS_PATH}" \
      --ts_out="${f}" \
      -I . \
      "./rcc.proto"

# done