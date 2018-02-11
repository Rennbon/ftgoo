#!/usr/bin/env bash
docker run --rm -v "$PWD":/go/src/ftgoo/cmd/fsgrpc/fs_server -w /go/src/ftgoo/cmd/fsgrpc/fs_server golang:1.9.1 go build -o myapp
docker build -t rennbonzhu/fs_server:latest .
rm -rf myapp

