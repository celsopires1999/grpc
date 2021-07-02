package main

import (
	"context"
	"fmt"
	"log"
	"io"
	"github.com/celsopires1999/grpc/pb"
	// "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1"
	"google.golang.org/grpc"
)

func main()  {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id: "0",
		Name: "João",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id: "0",
		Name: "João",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}