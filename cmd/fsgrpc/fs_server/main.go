package main

import (
	"log"
	"net"

	pb "ftgoo/logic/folderstat"

	"ftgoo/logic/mongodb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9091"
)

var service mongodb.FolderStatService

type fstatServer struct{}

func (*fstatServer) GetFolderStatByDate(cxt context.Context, param *pb.GetFolderStatByDateRequest) (*pb.GetFolderStatByDateResponse, error) {
	result, err := service.GetFolderStatByDate(param)
	return result, err
}

func (*fstatServer) GetFolderStatNow(ctx context.Context, param *pb.GetFolderStatNowRequest) (*pb.GetFolderStatNowResponse, error) {
	result, err := service.GetFolderStatNow(param)
	return result, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	log.Println("server start")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFolderstatServiceServer(s, &fstatServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("server start over")
}
