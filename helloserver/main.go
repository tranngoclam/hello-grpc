package main

import (
	"context"
	"fmt"
	pb "github.com/tranngoclam/hello-grpc/hellopb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
}

func (*server) Hello(_ context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := request.Name
	response := &pb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("server is listening on %v ...", address)
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
