package service

import (
	"context"
	"log"
	"os"
	"time"

	helloworld "go-third-party/gin/reverse_proxy_test/pb"
	"google.golang.org/grpc"
)

const (
	address = "grpc_server:50051"
	// address     = "127.0.0.1:50051"
	defaultName = "world"
)

func Api1() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
