package main

import (
	"context"
	"fmt"
	"log"

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

	AddUser(client)
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