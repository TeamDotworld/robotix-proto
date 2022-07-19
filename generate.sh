#!/bin/bash

protoc --go_out=plugins=grpc:/home/anand/go/src/ ./protos/rcc/v1/rcc.proto
protoc --go_out=plugins=grpc:/home/anand/go/src/ ./protos/agent/v1/agent.proto


# python3 -m grpc_tools.protoc -I. --python_out=/home/anand/go/src/github.com/anand-dotworld/robotix-agent/node --grpc_python_out=/home/anand/go/src/github.com/anand-dotworld/robotix-agent/node ./go-multi-tenant/agent/v1/agent.proto

# # cp rcc.proto /home/anand/Work/golang/src/delivery-robot-backend/go-multi-tenant/rcc.proto

for f in ./protos/model/v1/*.proto 
do 
    protoc --go_out=plugins=grpc:/home/anand/go/src/ $f
    echo "Processing $f file.."
done


# PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
# GRPC_TOOLS_NODE_PROTOC_PLUGIN="./node_modules/.bin/grpc_tools_node_protoc_plugin"
# GRPC_TOOLS_NODE_PROTOC="./node_modules/.bin/grpc_tools_node_protoc"
# f="./"
# # f="./go-multi-tenant"
# # for f in ./rcc.proto; do

# #   # skip the non proto files
# #   if [ "$(basename "$f")" == "index.ts" ]; then
# #       continue
# #   fi

#   # loop over all the available proto files and compile them into respective dir
#   # JavaScript code generating
#   ${GRPC_TOOLS_NODE_PROTOC} \
#       --js_out=import_style=commonjs,binary:"${f}" \
#       --grpc_out="${f}" \
#       --plugin=protoc-gen-grpc="${GRPC_TOOLS_NODE_PROTOC_PLUGIN}" \
#       -I . \
#       "./rcc.proto"

#   ${GRPC_TOOLS_NODE_PROTOC} \
#       --plugin=protoc-gen-ts="${PROTOC_GEN_TS_PATH}" \
#       --ts_out="${f}" \
#       -I . \
#       "./rcc.proto"

# done








declare -a packages=("rcc-multi-tenant" "robotix-agent" "element3")
for p in "${packages[@]}"
do
   echo "$p"
   # or do whatever with individual element of the array
done
# for f in ./model/v1/*.proto 
# do 
#     sed -i 's+go_package="";+go_package="rcc-multi-tenant/protos/v1/model";+g' $f
#     echo "writing $f file.."
# done

# for f in ./model/v1/*.proto 
# do 
#     protoc --go_out=plugins=grpc:/home/anand/go/src/github.com/anand-dotworld/ $f
#     echo "Processing $f file.."
# done
# sed -i 's+go_package="";+go_package="rcc-multi-tenant/protos/v1/agent";+g' ./protos/rcc/v1/rcc.proto
# protoc --go_out=plugins=grpc:/home/anand/go/src/github.com/anand-dotworld/ ./protos/rcc/v1/rcc.proto
# sed -i 's+go_package="rcc-multi-tenant/protos/v1/agent";+go_package="";+g' ./protos/rcc/v1/rcc.proto


# for f in ./model/v1/*.proto 
# do 
#     sed -i 's+go_package="rcc-multi-tenant/protos/v1/model";+go_package="";+g' $f
#     echo "removing $f file.."
# done