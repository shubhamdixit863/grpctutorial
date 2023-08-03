package handlers

import (
	"context"
	"fmt"
	"github.com/shubhamdixit863/grpctutorial/greet/proto"
	"github.com/shubhamdixit863/grpctutorial/greet/server/repository"
)

type Server struct {
	proto.GreetServiceServer
	Repo repository.Repository
}

func (s Server) GreetOnce(context.Context, *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Println("Client invoked the server", s.Repo.Save())
	return &proto.GreetResponse{
		Message: s.Repo.Save(),
		Address: s.Repo.Save(),
	}, nil
}

func (s Server) CreateGreet(context.Context, *proto.CreateGreetRequest) (*proto.GreetResponse, error) {
	fmt.Println("Client invoked the server", s.Repo.Save())
	return &proto.GreetResponse{
		Message: s.Repo.Save(),
		Address: s.Repo.Save(),
	}, nil
}
