#!/bin/bash
go install ./
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(basename $(pwd))_linux_amd64
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o $(basename $(pwd))_darwin_arm64
