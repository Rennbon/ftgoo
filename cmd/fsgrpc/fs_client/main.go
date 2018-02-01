package main

import (
	"log"

	pb "ftgoo/logic/folderstat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:12345"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFolderstatServiceClient(conn)

	r, err := c.GetFolderStatNow(context.Background(), &pb.GetFolderStatNowRequest{FolderId: "123"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Folderstat)
}
