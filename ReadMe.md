# Robotix Proto

This repo consists of the proto used by the RCC server, Agent and The node.

## Installing Protoc

To build the proto files into respective stub files you need to install the protoc compiler. To install follow the command below

    go install github.com/golang/protobuf/protoc-gen-go

## Build Proto

For building the proto files, simple run the update.sh file. This will build the proto files, changes the version for go package, commit and update the git remote repository and creates a new version for the go package. Also if provided your local working directory of agent and server, it will update the local package to the updated version.
