package main

import (
	"context"
	"fmt"
	pb "github.com/tranngoclam/hello-grpc/hellopb"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	arg := os.Args[1]
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client := pb.NewHelloServiceClient(cc)
	request := &pb.HelloRequest{Name: arg}
	response, err := client.Hello(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Greeting)
}
