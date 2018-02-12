package main

import (
	"log"

	pb "github.com/Rennbon/ftgoo/logic/folderstat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "www.rennbon.com:9091"
)

func main() {
	log.Println(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFolderstatServiceClient(conn)

	r, err := c.GetFolderStatNow(context.Background(), &pb.GetFolderStatNowRequest{FolderId: "191636da-e23f-43ed-b5f3-ae232a27d070"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Folderstat)
	log.Print("the end")
	r1, err := c.GetFolderStatByDate(
		context.Background(),
		&pb.GetFolderStatByDateRequest{
			FolderId:  "14f05406-d434-45ee-af5a-a62610f7e026",
			StartDate: 636534720000000000,
			EndDate:   636539904000000000,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r1)
}
