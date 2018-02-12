package main

import (
	"log"
	"net"

	//"github.com/Rennbon/ftgoo/config"

	pb "github.com/Rennbon/ftgoo/logic/folderstat"

	"github.com/Rennbon/ftgoo/logic/mongodb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9091"
)

var service mongodb.FolderStatService

//var ca config.Certificate

type fstatServer struct{}

func (*fstatServer) GetFolderStatByDate(cxt context.Context, param *pb.GetFolderStatByDateRequest) (*pb.GetFolderStatByDateResponse, error) {
	result, err := service.GetFolderStatByDate(param)
	return result, err
}

func (*fstatServer) GetFolderStatNow(ctx context.Context, param *pb.GetFolderStatNowRequest) (*pb.GetFolderStatNowResponse, error) {
	result, err := service.GetFolderStatNow(param)
	return result, err
}

/*
func init() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	var keys []string
	keys = append(keys, "Certificate")
	err = config.CheckConfig(conf, keys)
	if err != nil {
		panic(err)
	}
	ca = conf.Certificate
} */
func main() {
	lis, err := net.Listen("tcp", port)
	log.Println("server start")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	/* 	creds, err := credentials.NewServerTLSFromFile(ca.CertFile, ca.KeyFile)
	   	if err != nil {
	   		log.Fatalf("could not load keys: %s", err)
	   	}
	   	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)} */
	s := grpc.NewServer()
	pb.RegisterFolderstatServiceServer(s, &fstatServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("server start over")
}
