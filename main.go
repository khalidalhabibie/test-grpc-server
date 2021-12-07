package main

import (
	"context"
	"log"
	"net"

	pb "test-grpc-server/protos/product"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SendProductName(context context.Context, request *pb.ProductNameRequest) (*pb.ProductNameReply, error) {
	log.Println("server send product name")

	resp := &pb.ProductNameReply{
		Name: "name",
	}

	return resp, nil
}

func main() {
	log.Println("start grpc server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterProductServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
