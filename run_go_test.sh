#!/usr/bin/env bash

export GOPATH=${PWD}/go
echo "GOPATH: $GOPATH"

export GOROOT /usr/local/go
export PATH $PATH:$GOROOT/bin

if [ ! -f deliver ]; then
    mkdir -p $GOPATH/src/github.com/edmodo
    git clone git@github.com:edmodo/deliver $GOPATH/src/github.com/edmodo/deliver
    pushd $GOPATH/src/github.com/edmodo/deliver
    make build
    popd
    mv $GOPATH/src/github.com/edmodo/deliver/deliver ./
fi

mkdir -p ${PWD}/tmp

./deliver -v install

go run aws_api_tester.go
