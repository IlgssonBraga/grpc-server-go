package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/IlgssonBraga/grpc-go/pb"
	"google.golang.org/grpc"
)

func main()  {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to GRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)

	AddUsers(client)
}

func AddUser(client pb.UserServiceClient)  {
	req := &pb.User{
		Id:"0",
		Name:"Ilgsson",
		Email:"ilgsson@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req ) 

	if err != nil {
		log.Fatalf("Could not make GRPC Request: %v", err)
	}

	fmt.Println(res)

}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:"0",
		Name:"Ilgsson",
		Email:"ilgsson@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req ) 

	if err != nil {
		log.Fatalf("Could not make GRPC Request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}

}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id: "i1",
			Name: "Ilgsson",
			Email: "ilgsson@gmail.com",
		},
		&pb.User{
			Id: "i1",
			Name: "Ilgsson",
			Email: "ilgsson@gmail.com",
		},
		&pb.User{
			Id: "i2",
			Name: "Ilgsson2",
			Email: "ilgsson2@gmail.com",
		},
		&pb.User{
			Id: "i3",
			Name: "Ilgsson3",
			Email: "ilgsson3@gmail.com",
		},
		&pb.User{
			Id: "i4",
			Name: "Ilgsson4",
			Email: "ilgsson4@gmail.com",
		},
		&pb.User{
			Id: "i5",
			Name: "Ilgsson5",
			Email: "ilgsson5@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Erro creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Erro receiving response: %v", err)
	}

	fmt.Println(res)
}