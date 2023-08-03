package main

import (
	"fmt"
	"github.com/shubhamdixit863/grpctutorial/greet/proto"
	"github.com/shubhamdixit863/grpctutorial/greet/server/handlers"
	"github.com/shubhamdixit863/grpctutorial/greet/server/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

const address = "0.0.0.0:50051"

func main() {

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Grpc Server is Listening on %s\n", address)

	s := grpc.NewServer()
	repo := repository.MysqlRepository{}

	proto.RegisterGreetServiceServer(s, &handlers.Server{
		Repo: &repo,
	})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalln("Error Starting grpc server", err)

	}

}
