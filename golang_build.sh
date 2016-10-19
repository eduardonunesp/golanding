#!/bin/bash
go_version=1.7.1

docker run --rm -it -v $GOPATH:/gopath -v $(pwd):/app -e GOPATH=/gopath -w /app golang:$go_version sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags=-s -o main'
