#!/bin/bash

# The protoc compiles the proto code to go code. The files will be generated in the following directory
# /home/anand/go/src/github.com/TeamDotworld/robotix-proto/v1
# The folder structure for this repository should be like above. change the username
protoc --go_out=plugins=grpc:/home/amizhthan/go/src/ ./protos/rcc/v1/rcc.proto
protoc --go_out=plugins=grpc:/home/amizhthan/go/src/ ./protos/node_sdk/v1/node_sdk.proto
protoc --go_out=plugins=grpc:/home/amizhthan/go/src/ ./protos/agent/v1/agent.proto

# This line is for compiling python files for the agent.proto file
python3 -m grpc_tools.protoc -I. --python_out=/home/amizhthan/go/src/github.com/robotix-agent/node --grpc_python_out=/home/amizhthan/go/src/github.com/robotix-agent/node ./protos/agent/v1/agent.proto

# # cp rcc.proto /home/anand/Work/golang/src/delivery-robot-backend/go-multi-tenant/rcc.proto

# This will loop over all the model files and create both go and python files.
for f in ./protos/model/v1/*.proto 
do 
    protoc --go_out=plugins=grpc:/home/amizhthan/go/src/ $f
    python3 -m grpc_tools.protoc -I. --python_out=/home/amizhthan/go/src/github.com/robotix-agent/node --grpc_python_out=/home/amizhthan/go/src/github.com/robotix-agent/node $f

    echo "Processing $f file.."
done