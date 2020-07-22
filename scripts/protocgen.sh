#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  protoc \
  -I "proto" \
  -I "third_party/proto" \
  --gogofaster_out=\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
plugins=grpc,paths=source_relative:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

  if [[ $dir == *"abci"* ]]
  then
    protoc \
    -I "proto" \
    -I "third_party/proto" \
    --grpc-gateway_out=\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
paths=source_relative:. \
    $(find "${dir}" -maxdepth 1 -name '*.proto')
  fi
done

cp -r ./tendermint/* ./proto/*
rm -rf tendermint

mv ./proto/tendermint/abci/types.pb.go ./abci/types
mv ./proto/tendermint/abci/types.pb.gw.go ./abci/types

mv ./proto/tendermint/rpc/grpc/types.pb.go ./rpc/grpc
