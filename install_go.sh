#!/usr/bin/env bash

GO_VER="1.7.4"

curl https://storage.googleapis.com/golang/go${GO_VER}.linux-amd64.tar.gz > go.linux-amd64.tar.gz
tar -C /usr/local -xzf go.linux-amd64.tar.gz
