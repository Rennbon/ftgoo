FROM golang:latest

WORKDIR $GOPATH/src/ftgoo/cmd/fsdaily
ADD . $GOPATH/src/ftgoo/cmd/fsdaily
RUN go build .

ENTRYPOINT ["./ftgoo/cmd/fsdaily"]

