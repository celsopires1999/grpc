package main

import (
	"log"
	"net"

	"github.com/celsopires1999/grpc/pb"
	"github.com/celsopires1999/grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v, err")
	}

	grpServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpServer,  services.NewUserService())
	reflection.Register(grpServer)

	if err := grpServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
