package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/celsopires1999/grpc/pb"

	// "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
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
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
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

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "j1",
			Name:  "João 1",
			Email: "j1@j.com",
		},
		&pb.User{
			Id:    "j2",
			Name:  "João 2",
			Email: "j2@j.com",
		},
		&pb.User{
			Id:    "j3",
			Name:  "João 3",
			Email: "j3@j.com",
		},
		&pb.User{
			Id:    "j4",
			Name:  "João 4",
			Email: "j4@j.com",
		},
		&pb.User{
			Id:    "j5",
			Name:  "João 5",
			Email: "j5@j.com",
		},
		&pb.User{
			Id:    "j6",
			Name:  "João 6",
			Email: "j6@j.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep((time.Second * 3))
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "j1",
			Name:  "João 1",
			Email: "j1@j.com",
		},
		&pb.User{
			Id:    "j2",
			Name:  "João 2",
			Email: "j2@j.com",
		},
		&pb.User{
			Id:    "j3",
			Name:  "João 3",
			Email: "j3@j.com",
		},
		&pb.User{
			Id:    "j4",
			Name:  "João 4",
			Email: "j4@j.com",
		},
		&pb.User{
			Id:    "j5",
			Name:  "João 5",
			Email: "j5@j.com",
		},
		&pb.User{
			Id:    "j6",
			Name:  "João 6",
			Email: "j6@j.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 3)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}
			fmt.Printf("Receiving user %v with status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}
