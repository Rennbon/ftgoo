#!/usr/bin/env bash
docker run --rm -v "$PWD":/go/src/ftgoo/cmd/fsdaily -w /go/src/ftgoo/cmd/fsdaily golang:1.9.1 go build -o myapp
docker build -t rennbonzhu/fsdaily:latest .
rm -rf myapp

