package main

import (
	"context"
	"fmt"
	"github.com/shubhamdixit863/grpctutorial/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var address string = "localhost:50051" // Server's address on whihc grpc server is running

func main() {
	// it will call the grpc server running on the address
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return
		log.Fatalln("Failed to connect")
	}
	defer conn.Close()

	c := proto.NewGreetServiceClient(conn)
	d := proto.GreetRequest{
		Message: "",
		Address: "",
	}
	once, err := c.GreetOnce(context.Background(), &d)

	if err != nil {
		fmt.Println("Error in calling the grpc server GreetOnce Method", err)
		return
	}
	fmt.Println("Response received from the greet server", once)
}
