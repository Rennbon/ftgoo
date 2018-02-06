FROM golang:latest

WORKDIR $GOPATH/src/ftgoo/cmd/fsgrpc/fs_server
ADD . $GOPATH/src/ftgoo/cmd/fsgrpc/fs_server
RUN go build .

EXPOSE 9091
ENTRYPOINT ["./ftgoo/cmd/fsgrpc/fs_server"]

