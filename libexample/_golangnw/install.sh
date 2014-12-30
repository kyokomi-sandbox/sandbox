#!/bin/sh
go install
golang-nw-pkg -app=$GOPATH/bin/golangnw -name="golang-nw-example" -bin="gonw-example" -toolbar=false

