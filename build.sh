#!/bin/sh

export PROJECT=github.com/DavidSantia/sftp_util

# Build for Linux, statically linked

export NAME=sftp_cmd

echo "## Building $NAME utility"
docker run --rm --name golang -v $GOPATH/src:/go/src golang:alpine /bin/sh -l -c \
    "cd /go/src/$PROJECT/$NAME; CGO_ENABLED=0 /usr/local/go/bin/go build -v -i"
