#!/usr/bin/env bash

go build -o ./bin/protoc-gen-go github.com/golang/protobuf/protoc-gen-go

PATH=$PATH:$PWD/bin/

for pkg in $(find . -type f | grep "\.proto" | xargs -L1); do
  	echo "generate ==> $pkg"
	protoc -I. --go_out=paths=source_relative:. "$pkg"
done
