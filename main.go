package main

import (
	"context"
	"gRPC_Prac/Greeting"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myGreeterServer struct {
	Greeting.UnimplementedGreeterServer
}

func (s myGreeterServer) Create(ctx context.Context, server *Greeting.HelloRequest) (*Greeting.HelloReply, error) {
	return &Greeting.HelloReply{
		Message: "Hello my name is Shashwat",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")

	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myGreeterServer{}

	Greeting.RegisterGreeterServer(serverRegistrar, service)
	serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to server: %s", err)
	}
}
