#!/usr/bin/env bash
docker run --rm -v "$PWD":/go/src/github.com/Rennbon/ftgoo/cmd/fsgrpc/fs_server -w /go/src/github.com/Rennbon/ftgoo/cmd/fsgrpc/fs_server golang:latest go build -o myapp
docker build -t rennbonzhu/fs_server:latest .
rm -rf myapp

