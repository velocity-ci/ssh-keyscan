#!/bin/sh -e

go get -u github.com/mitchellh/gox

gox -output="dist/ssh-keyscan_{{.OS}}_{{.Arch}}" \
    -osarch="darwin/amd64 linux/amd64" \
    ./cmd
