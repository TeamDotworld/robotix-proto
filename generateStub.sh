#!/bin/bash

protoc --go_out=plugins=grpc:/home/anand/Work/golang/src/robotix-agent rcc.proto
protoc --go_out=plugins=grpc:/home/anand/Work/golang/src/robotix-command-center rcc.proto
python -m grpc_tools.protoc -I. --python_out=/home/anand/Work/golang/src/robotix-agent/ros_node/protos --grpc_python_out=/home/anand/Work/golang/src/robotix-agent/ros_node/protos rcc.proto
cp rcc.proto /home/anand/Work/golang/src/delivery-robot-backend/protos/rcc.proto