#!/bin/sh

export GOPATH=`realpath ./_test`

go get -v github.com/stretchr/testify

go test ./...